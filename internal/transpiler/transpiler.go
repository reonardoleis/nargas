package transpiler

import (
	"fmt"
	"strings"

	"github.com/reonardoleis/nargas/internal/ast"
)

type Transpiler struct {
	TranspilerStrategy
	out string
	in  *ast.AST
}

func NewTranspiler(in *ast.AST, outputType ...OutputType) *Transpiler {
	_outputType := Go
	if len(outputType) > 0 {
		_outputType = outputType[0]
	}

	transpiler := &Transpiler{
		in:  in,
		out: "",
	}

	transpiler.setStrategy(_outputType)

	return transpiler
}

func (t *Transpiler) setStrategy(outputType OutputType) {
	switch outputType {
	case JS:
		fmt.Println("Using JS")
		t.TranspilerStrategy = &JsStrategy{}
	default:
		fmt.Println("Using Go")
		t.TranspilerStrategy = &GoStrategy{}
	}
}

func (t Transpiler) Transpile() string {
	t.transpile(t.in.Expression)
	return t.Main(t.out)
}

func (t *Transpiler) transpile(node *ast.Value) {
	if node == nil {
		return
	}

	val, ty := node.Value()

	switch ty {
	case ast.IntVal, ast.BoolVal, ast.StrVal:
		t.out += t.Literal(val)
	case ast.TermVal:
		term := val.(*ast.Term)
		switch term.Kind {
		case Let:
			if term.Value != nil && t.NeedsPrototype() {
				if term.Value.TermVal.Kind == Function {
					t.out += t.FunctionPrototype(*term, *term.Value.TermVal, term.Value.TermVal.Parameters)
				} else {
					t.out += t.VariablePrototype(*term)
				}
			}

			t.out += t.Let(*term)

			t.transpile(term.Value)
			t.transpile(term.Next)

		case Function:
			t.out += t.Function(*term, term.Parameters)
			t.out += "{"
			t.transpile(term.Value)
			t.out += "\n}\n"
			t.transpile(term.Next)
		case Binary:
			t.out += strings.ToLower(term.Op) + "("
			t.transpile(term.Lhs)
			t.out += ", "
			t.transpile(term.Rhs)
			t.out += ")"
		case Var:
			t.out += term.Text
		case Int, Str, Bool:

			t.out += term.Value.String()

		case If:
			t.out += "\nif "
			t.out += t.ConditionStart()
			t.transpile(term.Condition)
			t.out += t.ConditionEnd()
			t.out += " {\n"
			if term.Then != nil && term.Then.TermVal.Kind != Let {
				t.out += t.Return()
			}
			t.transpile(term.Then)
			t.out += "\n}"
			if term.Otherwise != nil {
				t.out += " else {\n"
				if term.Otherwise != nil && term.Otherwise.TermVal.Kind != Let {
					t.out += t.Return()
				}
				t.transpile(term.Otherwise)
				t.out += "\n}"
			}

		case Call:
			t.transpile(term.Callee)
			t.out += "("
			for i, arg := range term.Arguments {
				t.transpile(arg)
				if i < len(term.Arguments)-1 {
					t.out += ", "
				}
			}
			t.out += ")"

		case Print:
			t.out += t.PrintStart()
			t.transpile(term.Value)
			t.out += t.PrintEnd()
		}

	}

}
