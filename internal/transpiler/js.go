package transpiler

import (
	"fmt"

	"github.com/reonardoleis/nargas/internal/ast"
)

var functionsJs = `
	function add(a, b) {
		return a + b;
	}

	function sub(a, b) {
		return a - b;
	}

	function div(a, b) {
		return Math.round(a / b);
	}

	function mul(a, b) {
		return a * b;
	}

	function rem(a, b) {
		return a % b;
	}
	
	function eq(a, b) {
		return a == b;
	}

	function neq(a, b) {
		return a != b;
	}

	function lt(a, b) {
		return a < b;
	}

	function gt(a, b) {
		return a > b;
	}

	function lte(a, b) {
		return a <= b;
	}
	
	function gte(a, b) {
		return a >= b;
	}

	function or(a, b) {
		return a || b;
	}

	function and(a, b) {
		return a && b;
	}

	
`

type JsStrategy struct{}

func (g JsStrategy) Literal(val interface{}) string {
	return fmt.Sprintf("%v", val)
}

func (g JsStrategy) Let(let ast.Term) string {
	return fmt.Sprintf("\n%s = ", let.Name.Text)
}

func (g JsStrategy) Function(function ast.Term, parameterList []ast.Name) string {
	out := "function("
	for i, param := range parameterList {
		out += param.Text
		if i < len(parameterList)-1 {
			out += ", "
		}
	}

	return out + ")"
}

func (g JsStrategy) NeedsPrototype() bool {
	return true
}

func (g JsStrategy) FunctionPrototype(let ast.Term, function ast.Term, parameterList []ast.Name) string {
	return ""
}

func (g JsStrategy) VariablePrototype(let ast.Term) string {
	return ""
}

func (g JsStrategy) Main(program string) string {
	return fmt.Sprintf("%s\n\n%s", functionsJs, program)
}

func (g JsStrategy) Return() string {
	return "return "
}

func (g JsStrategy) PrintStart() string {
	return "console.log("
}

func (g JsStrategy) PrintEnd() string {
	return ")"
}

func (g JsStrategy) ConditionStart() string {
	return "("
}

func (g JsStrategy) ConditionEnd() string {
	return ")"
}
