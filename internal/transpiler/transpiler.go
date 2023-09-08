package transpiler

import (
	"github.com/reonardoleis/nargas/internal/ast"
	"github.com/reonardoleis/nargas/token"
)

var Out = ""

func add(str string) {
	Out += str
}

func Transpile2Go(node ast.Attrib) {

	if node == nil {
		return
	}

	switch node := node.(type) {
	case *token.Token:
		println()
		return
	default:
		add(Node2Go(node))
	}

	children := node.(ast.Node).Children()
	for _, child := range children {
		Transpile2Go(child)
	}
}
