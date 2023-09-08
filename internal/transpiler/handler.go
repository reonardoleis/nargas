package transpiler

import (
	"github.com/reonardoleis/nargas/internal/ast"
	"github.com/reonardoleis/nargas/token"
)

func ParamList2JS(paramList ast.Attrib) string {
	paramListStr := ""
	current := paramList.(*ast.FunctionParamList)
	if current == nil {
		return ""
	}

	for current.Next != nil {
		paramListStr += string(current.Param.(*token.Token).Lit) + ", "
		current = current.Next.(*ast.FunctionParamList)
	}

	paramListStr += string(current.Param.(*token.Token).Lit)

	return paramListStr
}

func Expr2JS(expr ast.Attrib) string {
	if expr == nil {
		return ""
	}

	switch expr := expr.(type) {
	case *ast.Expr:
		if expr.Op == nil {
			return Expr2JS(expr.Left)
		}
		return Expr2JS(expr.Left) + string(expr.Op.(*token.Token).Lit) + Expr2JS(expr.Right)
	case *ast.Function:
		return "function (" + ParamList2JS(expr.ParamList) + ") {\n" + Body2JS(expr.Body) + "}"
	case *ast.Literal:
		return string(expr.Value.(*token.Token).Lit)
	case *token.Token:
		return string(expr.Lit)
	}

	return ""
}

func Body2JS(body ast.Attrib) string {
	if body == nil {
		return ""
	}

	commandList := body.(*ast.Body).CommandList

	current := commandList.(*ast.CommandList)

	if current == nil {
		return ""
	}

	bodyStr := ""

	for current.Next != nil {
		bodyStr += Command2JS(current.Command)
		current = current.Next.(*ast.CommandList)
	}

	bodyStr += Command2JS(current.Command)

	return bodyStr
}

func Command2JS(node ast.Attrib) string {
	if node == nil {
		return ""
	}

	command := node.(*ast.Command)

	switch specificCommand := command.Value.(type) {
	case *ast.CommandVarAssign:
		return string(specificCommand.ID.(*token.Token).Lit) + " = " + Expr2JS(specificCommand.Value) + ";\n"
	case *ast.CommandReturn:
		return "return " + Expr2JS(specificCommand.Value) + ";\n"
	}

	return ""
}

func Node2JS(node ast.Attrib) string {
	if node == nil {
		return ""
	}

	switch node := node.(type) {
	case *ast.Program:
		return ""
	case *ast.DeclarationList:
		return ""
	case *ast.Declaration:
		return "let " + string(node.ID.(*token.Token).Lit) + " = " + Expr2JS(node.Value) + ";\n"
	}

	return ""
}
