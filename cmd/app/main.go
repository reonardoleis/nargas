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

	if len(os.Args) < 3 {
		panic("Output language & input file's name must be specified")
	}

	outputFileName := "main"
	if len(os.Args) > 3 {
		outputFileName = os.Args[3]
	}

	file, _ := os.ReadFile(os.Args[1])
	outputLanguage := os.Args[2]

	ast := ast.AST{}
	err := json.Unmarshal(file, &ast)
	if err != nil {
		panic(err)
	}

	transpiler := transpilerPkg.NewTranspiler(&ast, transpilerPkg.OutputType(outputLanguage))
	out := transpiler.Transpile()

	f, err := os.Create("./generated/main." + outputLanguage)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(out)
	if err != nil {
		panic(err)
	}

	if outputLanguage == "go" {
		build.Generate(out, outputFileName, transpilerPkg.Go)
	}

	fmt.Println("Transpiling done!")
}
