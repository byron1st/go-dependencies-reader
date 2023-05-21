package lib

import (
	"fmt"
	"strings"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/static"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa/ssautil"
)

func ReadDependencies(mainPkgName string, dir string) error {
	pkgs, err := loadPkgs(mainPkgName, dir)
	if err != nil {
		return err
	}

	prog, _ := ssautil.AllPackages(pkgs, 0)
	prog.Build()

	cg := static.CallGraph(prog)
	cg.DeleteSyntheticNodes()

	if err := callgraph.GraphVisitEdges(cg, func(edge *callgraph.Edge) error {
		callerPkg := edge.Caller.Func.Pkg.Pkg.Path()
		callerFunc := edge.Caller.Func.Name()
		calleePkg := edge.Callee.Func.Pkg.Pkg.Path()
		calleeFunc := edge.Callee.Func.Name()

		if !strings.HasPrefix(callerPkg, mainPkgName) ||
			callerFunc == "init" ||
			calleeFunc == "init" {
			return nil
		}

		dr := DependencyRelation{
			Caller: fmt.Sprintf("%s.%s", callerPkg, callerFunc),
			Callee: fmt.Sprintf("%s.%s", calleePkg, calleeFunc),
		}
		if err := dr.Print(); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func loadPkgs(mainPkgName string, dir string) ([]*packages.Package, error) {
	cfg := &packages.Config{
		Mode: packages.NeedDeps |
			packages.NeedSyntax |
			packages.NeedTypesInfo |
			packages.NeedTypes |
			packages.NeedTypesSizes |
			packages.NeedImports |
			packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles,
		Tests: false,
		Dir:   dir,
	}

	pkgs, err := packages.Load(cfg, fmt.Sprintf("%s/...", mainPkgName))
	if err != nil {
		return nil, err
	} else if packages.PrintErrors(pkgs) > 0 {
		return nil, fmt.Errorf("packages contain errors")
	}

	return pkgs, nil
}
