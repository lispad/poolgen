// Copyright 2019 Pavel Kirikov
package main

import (
	"flag"
	"fmt"
	"github.com/lispad/poolgen/generator"
	"github.com/lispad/poolgen/parser"
	"os"
)

func main() {
	var typeName, poolName string

	flag.CommandLine.Usage = printUsage // customizes the output of `poolgen -h`
	flag.StringVar(&typeName, "type", "", "Name of type to generate pool for.")
	flag.StringVar(&poolName, "pool", "", "Name of pool to be generated, use only with type name.")
	flag.Parse()

	files := flag.Args()
	if len(files) < 1 || (poolName != "" && typeName == "") {
		printUsage()
		os.Exit(1)
	}

	for _, name := range files {
		if err := generate(name, typeName, poolName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "PoolGen is a tool for generating pool of objects for Go projects")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "\t\"poolgen [options] /path/to/the/go/file.go\" or")
	fmt.Fprintln(os.Stderr, "\tplace \"//go:generate poolgen $GOFILE\" in your go file and run \"go generate ...\"")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Options:")
	fmt.Fprintf(os.Stderr, "\t-typet=<name of typet>. By default generates pool for all types in file")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintf(os.Stderr, "\t-pool=<name of pool to be generated>. Use only with \"type\" option")
}

func generate(source string, typeName string, poolName string) error {
	if _, err := os.Stat(source); err != nil {
		return err
	}

	parsing, err := parser.Parse(source, typeName)
	if err != nil {
		return err
	}

	if len(parsing.Types) == 0 {
		fmt.Fprintf(os.Stderr, "No types found in file %s\n", source)
		os.Exit(1)
	}

	outFile, err := generator.NewGoFile(source, parsing.Package)
	if err != nil {
		return err
	}

	for _, typeName := range parsing.Types {
		if err := outFile.WriteReset(typeName); err != nil {
			return err
		}
		if err := outFile.WritePool(typeName, poolName); err != nil {
			return err
		}
	}

	return nil
}
