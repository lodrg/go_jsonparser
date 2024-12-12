package lpjsonparser

import "fmt"

type TokenType int

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

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("Token{  Type:%v, Value:%v  }", t.Type, t.Value)
}

func newToken(tokenType TokenType, value string) Token {
	return Token{
		Type:  tokenType,
		Value: value,
	}
}
