package transpiler

type Kind = string

const (
	If        Kind = "If"
	Let       Kind = "Let"
	Str       Kind = "Str"
	Bool      Kind = "Bool"
	Int       Kind = "Int"
	Binary    Kind = "Binary"
	Call      Kind = "Call"
	Function  Kind = "Function"
	Print     Kind = "Print"
	First     Kind = "First"
	Second    Kind = "Second"
	Tuple     Kind = "Tuple"
	Parameter Kind = "Parameter"
	Var       Kind = "Var"
	term      Kind = "term"
)
