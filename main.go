package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/byron1st/go-dependencies-reader/lib"
)

func main() {
	mainPkgName, dir, err := getMainPkgName()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := lib.ReadDependencies(mainPkgName, dir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getMainPkgName() (string, string, error) {
	mainPkgName := ""
	dir := ""

	flag.StringVar(&mainPkgName, "main", "", "Main package name")
	flag.StringVar(&dir, "dir", ".", "Directory to search for go files")
	flag.Parse()

	if mainPkgName == "" {
		return "", "", fmt.Errorf("Please specify main package name")
	}

	return mainPkgName, dir, nil
}
