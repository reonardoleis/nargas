package transpiler

import (
	"github.com/reonardoleis/nargas/internal/ast"
	"github.com/reonardoleis/nargas/token"
)

func ParamList2Go(paramList ast.Attrib) string {
	paramListStr := ""
	current := paramList.(*ast.FunctionParamList)
	if current == nil {
		return ""
	}

	for current.Next != nil {
		paramListStr += string(current.Param.(*token.Token).Lit) + " interface{}, "
		current = current.Next.(*ast.FunctionParamList)
	}

	paramListStr += string(current.Param.(*token.Token).Lit) + " interface{}"

	return paramListStr
}

func Expr2Go(expr ast.Attrib) string {
	if expr == nil {
		return ""
	}

	switch expr := expr.(type) {
	case *ast.Expr:
		if expr.Op == nil {
			return Expr2Go(expr.Left)
		}
		return Expr2Go(expr.Left) + string(expr.Op.(*token.Token).Lit) + Expr2Go(expr.Right)
	case *ast.Function:
		return "func (" + ParamList2Go(expr.ParamList) + ") interface{} {\n" + Body2Go(expr.Body) + "}"
	case *ast.Literal:
		return string(expr.Value.(*token.Token).Lit)
	case *token.Token:
		return string(expr.Lit)
	}

	return ""
}

func Body2Go(body ast.Attrib) string {
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
		bodyStr += Command2Go(current.Command)
		current = current.Next.(*ast.CommandList)
	}

	bodyStr += Command2Go(current.Command)

	return bodyStr
}

func Command2Go(node ast.Attrib) string {
	if node == nil {
		return ""
	}

	command := node.(*ast.Command)

	switch specificCommand := command.Value.(type) {
	case *ast.CommandVarAssign:
		return string(specificCommand.ID.(*token.Token).Lit) + " = " + Expr2Go(specificCommand.Value) + "\n"
	case *ast.CommandReturn:
		return "return " + Expr2Go(specificCommand.Value) + "\n"
	}

	return ""
}

func Node2Go(node ast.Attrib) string {
	if node == nil {
		return ""
	}

	switch node := node.(type) {
	case *ast.Program:
		return ""
	case *ast.DeclarationList:
		return ""
	case *ast.Declaration:
		return "" + string(node.ID.(*token.Token).Lit) + " := " + Expr2Go(node.Value) + "\n"
	}

	return ""
}
