package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/reonardoleis/nargas/internal/ast"
	"github.com/reonardoleis/nargas/internal/build"
	transpilerPkg "github.com/reonardoleis/nargas/internal/transpiler"
)

func main() {
	if len(os.Args) < 2 {
		panic("Input file's name must be specified")
	}

	outputFileName := "main"
	if len(os.Args) > 2 {
		outputFileName = os.Args[2]
	}

	file, _ := os.ReadFile(os.Args[1])

	ast := ast.AST{}
	err := json.Unmarshal(file, &ast)
	if err != nil {
		panic(err)
	}

	transpiler := transpilerPkg.NewTranspiler(&ast, transpilerPkg.Go)
	out := transpiler.Transpile()

	f, err := os.Create("./generated/main.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(out)
	if err != nil {
		panic(err)
	}

	build.Generate(out, outputFileName, transpilerPkg.Go)

	fmt.Println("Transpiling done!")
}
