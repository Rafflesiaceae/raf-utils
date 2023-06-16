package main

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
)

func yamlLoc(inputFilePath string, pathQuery string) error {
	f, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	path, err := yaml.PathString(pathQuery)
	if err != nil {
		return err
	}

	var n ast.Node
	if n, err = path.ReadNode(f); err != nil {
		return err
	}

	t := n.GetToken()
	fmt.Printf("%d;%d;%d;%d\n",
		t.Position.IndentLevel,
		t.Position.Line,
		t.Position.Column,
		t.Position.Offset,
	)

	return nil
}
