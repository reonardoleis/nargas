package ast

var _ = (Node)(&Program{})
var _ = (Node)(&Declaration{})
var _ = (Node)(&DeclarationList{})
var _ = (Node)(&Literal{})

type Attrib interface {
}

type Node interface {
	Children() []Attrib
}

type Program struct {
	DeclarationList Attrib
}

func (p *Program) Children() []Attrib {
	return []Attrib{p.DeclarationList}
}

func NewProgram(declarationList Attrib) (*Program, error) {
	return &Program{declarationList}, nil
}

type Declaration struct {
	ID    Attrib
	Value Attrib
}

func (vd *Declaration) Children() []Attrib {
	return []Attrib{vd.ID, vd.Value}
}

func NewDeclaration(id, value Attrib) (*Declaration, error) {
	return &Declaration{id, value}, nil
}

type DeclarationList struct {
	Declaration Attrib
	Next        Attrib
}

func (dl *DeclarationList) Children() []Attrib {
	return []Attrib{dl.Declaration, dl.Next}
}

func NewDeclarationList(declaration, next Attrib) (*DeclarationList, error) {
	return &DeclarationList{declaration, next}, nil
}

type Literal struct {
	Value interface{}
}

func (l *Literal) Children() []Attrib {
	return []Attrib{}
}

func NewLiteral(value Attrib) (*Literal, error) {
	return &Literal{value}, nil
}

type Expr struct {
	Left  Attrib
	Op    Attrib
	Right Attrib
}

func (e *Expr) Children() []Attrib {
	return []Attrib{e.Left, e.Op, e.Right}
}

func NewExpr(left, op, right Attrib) (*Expr, error) {
	return &Expr{left, op, right}, nil
}

type Function struct {
	ParamList Attrib
	Body      Attrib
}

func (f *Function) Children() []Attrib {
	return []Attrib{f.ParamList, f.Body}
}

func NewFunction(paramList, body Attrib) (*Function, error) {
	return &Function{paramList, body}, nil
}

type FunctionParamList struct {
	Param Attrib
	Next  Attrib
}

func (fpl *FunctionParamList) Children() []Attrib {
	return []Attrib{fpl.Param, fpl.Next}
}

func NewFunctionParamList(param, next Attrib) (*FunctionParamList, error) {
	return &FunctionParamList{param, next}, nil
}

type Body struct {
	CommandList Attrib
}

func (b *Body) Children() []Attrib {
	return []Attrib{b.CommandList}
}

func NewBody(commandList Attrib) (*Body, error) {
	return &Body{commandList}, nil
}

type CommandList struct {
	Command Attrib
	Next    Attrib
}

func (cl *CommandList) Children() []Attrib {
	return []Attrib{cl.Command, cl.Next}
}

func NewCommandList(command, next Attrib) (*CommandList, error) {
	return &CommandList{command, next}, nil
}

type Command struct {
	Value Attrib
}

func (c *Command) Children() []Attrib {
	return []Attrib{c.Value}
}

func NewCommand(value Attrib) (*Command, error) {
	return &Command{value}, nil
}

type CommandVarAssign struct {
	ID    Attrib
	Value Attrib
}

func (cva *CommandVarAssign) Children() []Attrib {
	return []Attrib{cva.ID, cva.Value}
}

func NewCommandVarAssign(id, value Attrib) *CommandVarAssign {
	return &CommandVarAssign{id, value}
}

type CommandReturn struct {
	Value Attrib
}

func (cr *CommandReturn) Children() []Attrib {
	return []Attrib{cr.Value}
}

func NewCommandReturn(value Attrib) (*CommandReturn, error) {
	return &CommandReturn{value}, nil
}
