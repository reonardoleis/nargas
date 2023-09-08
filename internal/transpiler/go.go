package transpiler

import (
	"fmt"

	"github.com/reonardoleis/nargas/internal/ast"
)

var functionsGo = `
	func add(a, b any) any {
		switch a.(type) {
		case int:
			switch b.(type) {
			case int:
				return a.(int) + b.(int)
			default:
				return fmt.Sprintf("%v%v", a, b)
			}
		default:
			return fmt.Sprintf("%v%v", a, b)
		}
	}

	func sub(a, b any) any {
		switch a.(type) {
		case int:
			switch b.(type) {
			case int:
				return a.(int) - b.(int)
			default:
				return fmt.Sprintf("%v%v", a, b)
			}
		default:
			return ""
		}
	}

	func div(a, b any) any {
		switch a.(type) {
		case int:
			switch b.(type) {
			case int:
				return a.(int) / b.(int)
			default:
				return fmt.Sprintf("%v%v", a, b)
			}
		default:
			return fmt.Sprintf("%v%v", a, b)
		}
	}

	func mul(a, b any) any {
		switch a.(type) {
		case int:
			switch b.(type) {
			case int:
				return a.(int) * b.(int)
			default:
				return fmt.Sprintf("%v%v", a, b)
			}
		default:
			return fmt.Sprintf("%v%v", a, b)
		}
	}

	func rem(a, b any) any {
		switch a.(type) {
		case int:
			switch b.(type) {
			case int:
				return a.(int) % b.(int)
			default:
				return fmt.Sprintf("%v%v", a, b)
			}
		default:
			return fmt.Sprintf("%v%v", a, b)
		}	
	}
	
	func eq(a, b any) bool {
		return a == b
	}

	func neq(a, b any) bool {
		return a != b
	}

	func lt(a, b any) bool {
		switch a.(type) {
		case int:
			switch b.(type) {
			case int:
				return a.(int) < b.(int)
			default:
				return false
			}
		default:
			return false
		}
	}

	func gt(a, b any) bool {
		switch a.(type) {
		case int:
			switch b.(type) {
			case int:
				return a.(int) > b.(int)
			default:
				return false
			}
		default:
			return false
		}
	}

	func lte(a, b any) bool {
		switch a.(type) {
		case int:
			switch b.(type) {
			case int:
				return a.(int) <= b.(int)
			default:
				return false
			}
		default:
			return false
		}
	}
	
	func gte(a, b any) bool {
		switch a.(type) {
		case int:
			switch b.(type) {
			case int:
				return a.(int) >= b.(int)
			default:
				return false
			}
		default:
			return false
		}
	}

	func or(a, b any) bool {
		switch a.(type) {
		case bool:
			switch b.(type) {
			case bool:
				return a.(bool) || b.(bool)
			default:
				return false
			}
		default:
			return false
		}
	}

	func and(a, b any) bool {
		switch a.(type) {
		case bool:
			switch b.(type) {
			case bool:
				return a.(bool) && b.(bool)
			default:
				return false
			}
		default:
			return false
		}
	}

	
`

type GoStrategy struct{}

func (g GoStrategy) Literal(val interface{}) string {
	return fmt.Sprintf("%v", val)
}

func (g GoStrategy) Let(let ast.Term) string {
	return fmt.Sprintf("\n%s = ", let.Name.Text)
}

func (g GoStrategy) Function(function ast.Term, parameterList []ast.Name) string {
	out := "func("
	for i, param := range parameterList {
		out += param.Text + " any"
		if i < len(parameterList)-1 {
			out += ", "
		}
	}

	return out + ") any"
}

func (g GoStrategy) NeedsPrototype() bool {
	return true
}

func (g GoStrategy) FunctionPrototype(let ast.Term, function ast.Term, parameterList []ast.Name) string {
	out := "var " + let.Name.Text + " func("
	for i, param := range parameterList {
		out += param.Text + " any"
		if i < len(parameterList)-1 {
			out += ", "
		}
	}

	return out + ") any\n"
}

func (g GoStrategy) VariablePrototype(let ast.Term) string {
	return "\nvar " + let.Name.Text + " any"
}

func (g GoStrategy) Main(program string) string {
	return fmt.Sprintf("package main\n\nimport \"fmt\"\n\n%s\n\nfunc main() {\n%s\n}", functionsGo, program)
}

func (g GoStrategy) Return() string {
	return "return "
}

func (g GoStrategy) PrintStart() string {
	return "fmt.Println("
}

func (g GoStrategy) PrintEnd() string {
	return ")"
}

func (g GoStrategy) ConditionStart() string {
	return ""
}

func (g GoStrategy) ConditionEnd() string {
	return ""
}
