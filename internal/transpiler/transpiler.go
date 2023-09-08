package transpiler

import (
	"github.com/reonardoleis/nargas/internal/ast"
	"github.com/reonardoleis/nargas/token"
)

var Out = ""

func add(str string) {
	Out += str
}

func Transpile2JS(node ast.Attrib) {

	if node == nil {
		return
	}

	switch node := node.(type) {
	case *token.Token:
		println()
		return
	default:
		add(Node2JS(node))
	}

	children := node.(ast.Node).Children()
	for _, child := range children {
		Transpile2JS(child)
	}
}
