// Code generated by gocc; DO NOT EDIT.

package errors

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/reonardoleis/nargas/token"
)

type ErrorSymbol interface {
}

type Error struct {
	Err            error
	ErrorToken     *token.Token
	ErrorSymbols   []ErrorSymbol
	ExpectedTokens []string
	StackTop       int
}

func (e *Error) String() string {
	w := new(strings.Builder)
	if e.Err != nil {
		fmt.Fprintln(w, "Error ", e.Err)
	} else {
		fmt.Fprintln(w, "Error")
	}
	fmt.Fprintf(w, "Token: type=%d, lit=%s\n", e.ErrorToken.Type, e.ErrorToken.Lit)
	fmt.Fprintf(w, "Pos: offset=%d, line=%d, column=%d\n", e.ErrorToken.Pos.Offset, e.ErrorToken.Pos.Line, e.ErrorToken.Pos.Column)
	fmt.Fprint(w, "Expected one of: ")
	for _, sym := range e.ExpectedTokens {
		fmt.Fprint(w, string(sym), " ")
	}
	fmt.Fprintln(w, "ErrorSymbol:")
	for _, sym := range e.ErrorSymbols {
		fmt.Fprintf(w, "%v\n", sym)
	}

	return w.String()
}

func DescribeExpected(tokens []string) string {
	switch len(tokens) {
	case 0:
		return "unexpected additional tokens"

	case 1:
		return "expected " + tokens[0]

	case 2:
		return "expected either " + tokens[0] + " or " + tokens[1]

	case 3:
		// Oxford-comma rules require more than 3 items in a list for the
		// comma to appear before the 'or'
		return fmt.Sprintf("expected one of %s, %s or %s", tokens[0], tokens[1], tokens[2])

	default:
		// Oxford-comma separated alternatives list.
		tokens = append(tokens[:len(tokens)-1], "or "+tokens[len(tokens)-1])
		return "expected one of " + strings.Join(tokens, ", ")
	}
}

func DescribeToken(tok *token.Token) string {
	switch tok.Type {
	case token.INVALID:
		return fmt.Sprintf("unknown/invalid token %q", tok.Lit)
	case token.EOF:
		return "end-of-file"
	default:
		return fmt.Sprintf("%q", tok.Lit)
	}
}

func (e *Error) Error() string {
	// identify the line and column of the error in 'gnu' style so it can be understood
	// by editors and IDEs; user will need to prefix it with a filename.
	text := fmt.Sprintf("%d:%d: error: ", e.ErrorToken.Pos.Line, e.ErrorToken.Pos.Column)

	// See if the error token can provide us with the filename.
	switch src := e.ErrorToken.Pos.Context.(type) {
	case token.Sourcer:
		text = src.Source() + ":" + text
	}

	if e.Err != nil {
		// Custom error specified, e.g. by << nil, errors.New("missing newline") >>
		text += e.Err.Error()
	} else {
		tokens := make([]string, len(e.ExpectedTokens))
		for idx, token := range e.ExpectedTokens {
			if !unicode.IsLetter(rune(token[0])) {
				token = strconv.Quote(token)
			}
			tokens[idx] = token
		}
		text += DescribeExpected(tokens)
		actual := DescribeToken(e.ErrorToken)
		text += fmt.Sprintf("; got: %s", actual)
	}

	return text
}
