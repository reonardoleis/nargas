package ast

import (
	"encoding/json"
	"fmt"
)

type Name struct {
	Text string `json:"text"`
}

type IValue interface {
	Value()
}

type Value struct {
	IntVal  *int32
	StrVal  *string
	BoolVal *bool
	TermVal *Term
}

func (v *Value) Value() (interface{}, ValueType) {
	if v.IntVal != nil {
		return *v.IntVal, IntVal
	}

	if v.StrVal != nil {
		return *v.StrVal, StrVal
	}

	if v.BoolVal != nil {
		return *v.BoolVal, BoolVal
	}

	if v.TermVal != nil {
		return v.TermVal, TermVal
	}

	return nil, ""
}

func (v Value) String() string {
	val, _ := v.Value()
	return fmt.Sprintf("%v", val)
}

func (v *Value) UnmarshalJSON(b []byte) error {
	var i int32
	var s string
	var b2 bool
	var t Term

	if err := json.Unmarshal(b, &i); err == nil {
		v.IntVal = &i
		return nil
	}

	if err := json.Unmarshal(b, &s); err == nil {
		v.StrVal = &s
		return nil
	}

	if err := json.Unmarshal(b, &b2); err == nil {
		v.BoolVal = &b2
		return nil
	}

	if err := json.Unmarshal(b, &t); err == nil {
		v.TermVal = &t
		return nil
	}

	return fmt.Errorf("invalid value")
}

type BinaryOp struct{}

type Term struct {
	Kind string `json:"kind"`
	Name *Name  `json:"name"`

	OnReturn bool `json:"-"`

	// If
	Condition *Value `json:"condition"`
	Then      *Value `json:"then"`
	Otherwise *Value `json:"otherwise"`

	// Let
	Value *Value `json:"value"`
	Next  *Value `json:"next"`

	// Binary
	Lhs *Value `json:"lhs"`
	Op  string `json:"op"`
	Rhs *Value `json:"rhs"`

	// Call
	Callee    *Value   `json:"callee"`
	Arguments []*Value `json:"arguments"`

	// Function
	Parameters []Name `json:"parameters"`

	// Tuple
	First  *Value `json:"first"`
	Second *Value `json:"second"`

	// Var
	Text string `json:"text"`
}

type AST struct {
	Expression *Value `json:"expression"`
}
