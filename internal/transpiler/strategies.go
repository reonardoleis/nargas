package transpiler

import (
	"github.com/reonardoleis/nargas/internal/ast"
)

var _ = (TranspilerStrategy)(GoStrategy{})

type TranspilerStrategy interface {
	Literal(val interface{}) string
	Let(let ast.Term) string
	Function(function ast.Term, parameterList []ast.Name) string
	Main(program string) string
	Return() string
	PrintStart() string
	PrintEnd() string
	FunctionPrototype(let ast.Term, function ast.Term, parameterList []ast.Name) string
	VariablePrototype(let ast.Term) string
	NeedsPrototype() bool
}
