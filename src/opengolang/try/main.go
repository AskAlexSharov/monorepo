package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	goroot := os.Getenv("GOROOT")
	p := path.Join(goroot, "src", "io")

	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if filepath.Ext(f.Name()) != ".go" {
			continue
		}
		if strings.HasSuffix(f.Name(), "_test.go") {
			continue
		}


		v := visitor {
			file: path.Join(p, f.Name()),
		}
		v.Parse()
	}
}

type visitor struct{
	file string
}

func (v visitor) Parse() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, v.file, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	ast.Walk(v, f)
}

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch d := n.(type) {
	case *ast.FuncDecl:
		//for _, f := range d.Type.Params.List {
		//	for _, name := range f.Names {
		//		v.local(name)
		//	}
		//}

		var doc string
		if d.Doc != nil {
			doc = d.Doc.Text()
		}
		_ = doc

		fmt.Printf("%#v %#v %#v %#v %#v \n", d.Pos(), d.End(), d.Type.End(), d.Type.Params.List, d.Type.Results.List)

		return nil
	}

	return v
}
