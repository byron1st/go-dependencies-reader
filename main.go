package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/byron1st/go-dependencies-reader/lib"
)

func main() {
	mainPkgName, err := getMainPkgName()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := lib.ReadDependencies(mainPkgName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getMainPkgName() (string, error) {
	mainPkgName := ""

	flag.StringVar(&mainPkgName, "main", "", "Main package name")
	flag.Parse()

	if mainPkgName == "" {
		return "", fmt.Errorf("Please specify main package name")
	}

	return mainPkgName, nil
}
