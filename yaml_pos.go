package main

import (
	"fmt"
	"io/ioutil"

	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/lexer"
	"github.com/goccy/go-yaml/parser"
)

type Visitor struct {
	searchLine int
	searchCol  int
	found      string
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	tk := node.GetToken()
	if tk.Position.Line == v.searchLine {
		if v.searchCol >= tk.Position.Column && v.searchCol <= tk.Position.Column+len(tk.Value) {
			v.found = node.GetPath()
			return nil
		}
	}

	return v
}

func yamlPos(inputFilePath string, line int, col int) error {
	contents, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		return err
	}

	tokens := lexer.Tokenize(string(contents))
	f, err := parser.Parse(tokens, 0)
	if err != nil {
		return fmt.Errorf("%+v", err)
	}

	v := Visitor{searchLine: line, searchCol: col}
	for _, doc := range f.Docs {
		ast.Walk(&v, doc.Body)
	}

	if v.found != "" {
		fmt.Printf("%+v\n", v.found)
	} else {
		return fmt.Errorf("unknown position")
	}

	return nil
}
