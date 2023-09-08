package main

import (
	"fmt"
	"os"

	"github.com/reonardoleis/nargas/internal/cli"
	"github.com/reonardoleis/nargas/internal/transpiler"
	"github.com/reonardoleis/nargas/lexer"
	"github.com/reonardoleis/nargas/parser"
)

func main() {
	argHandlers := []cli.ArgHandler{
		{Arg: "-h", Handler: cli.HandleHelp},
	}

	cli.HandleArgs(os.Args, argHandlers)

	input, _ := os.ReadFile("in.rinha")
	lex := lexer.NewLexer(input)
	par := parser.NewParser()
	st, err := par.Parse(lex)
	if err != nil {
		panic(err)
	}

	transpiler.Transpile2JS(st)
	fmt.Println(transpiler.Out)
}
