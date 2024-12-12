package lpjsonparser

import "fmt"

// alias type of int
type TokenType int

// otia as enum
const (
	LEFT_BRACE TokenType = iota
	RIGHT_BRACE
	LEFT_BRACKET
	RIGHT_BRACKET
	COMMA
	COLON
	STRING
	NUMBER
	EOF
	ILLEGAL
)

// struct
type Token struct {
	Type  TokenType
	Value string
}

// String implement Stringer interface then you can use fmt.PrintLn() func to print
func (t Token) String() string {
	return fmt.Sprintf("Token{  Type:%v, Value:%v  }", t.Type, t.Value)
}

func newToken(tokenType TokenType, value string) Token {
	return Token{
		Type:  tokenType,
		Value: value,
	}
}
