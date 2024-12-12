package test

import (
	"fmt"
	. "go_jsonparser/api/lpjsonparser"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `{"name":"John","age":30}`
	l := NewLexer(input)

	fmt.Println("Input runes:")
	for i, r := range []rune(input) {
		fmt.Printf("[%d] %q (hex: %X)\n", i, r, r)
	}

	fmt.Println("\nTokens:")
	for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
		fmt.Printf("%s\n", tok)
	}
}
