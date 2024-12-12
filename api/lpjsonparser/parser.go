package lpjsonparser

import (
	"errors"
	"fmt"
	"strconv"
)

type parser struct {
	tokens   []Token
	position int
}

func newParser(tokens []Token) *parser {
	return &parser{tokens, 0}
}

func Parse(json string) (interface{}, error) {
	lexer := NewLexer(json)
	tokens := make([]Token, 0)

	for {
		token := lexer.NextToken()
		tokens = append(tokens, token)

		if token.Type == EOF {
			break
		}
	}

	parser := newParser(tokens)
	return parser.parse()
}

func (p *parser) parse() (interface{}, error) {
	value, _ := p.parseValue()
	if p.peek().Type != EOF {
		return nil, errors.New("expected EOF")
	}
	return value, nil
}

func (p *parser) parseValue() (interface{}, error) {
	if len(p.tokens) == 0 {
		return nil, fmt.Errorf("empty tokens")
	}

	token := p.peek()
	switch token.Type {
	case LEFT_BRACE:
		return p.parseObject()
	case LEFT_BRACKET:
		return p.parseArray()
	case STRING:
		return p.parseString()
	case NUMBER:
		return p.parseNumber()
	default:
		return nil, errors.New("invalid token type")
	}
}

func (p *parser) parseObject() (interface{}, error) {
	m := map[string]interface{}{}
	p.match(TokenType(LEFT_BRACE))
	for !p.match(TokenType(RIGHT_BRACE)) {
		if !(p.peek().Type == STRING) {
			return nil, errors.New("expected string")
		}
		parseValue, _ := p.parseValue()
		key := parseValue.(string)
		if !(p.match(COLON)) {
			return nil, errors.New("expected :")
		}
		value, _ := p.parseValue()
		m[key] = value
		if !(p.match(COMMA)) && !(p.peek().Type == RIGHT_BRACE) {
			return nil, errors.New("expected , } or object")
		}
	}
	return m, nil
}

func (p *parser) parseArray() (interface{}, error) {
	list := make([]interface{}, 0)
	p.match(LEFT_BRACKET)

	for !p.match(RIGHT_BRACKET) {
		value, err := p.parseValue()
		if err != nil {
			// 处理错误
			return nil, nil
		}
		list = append(list, value)

		if !p.match(COMMA) && p.peek().Type != RIGHT_BRACKET {
			panic("Expected ',' or ']' in array")
			// 或者返回错误：
			// return nil, fmt.Errorf("Expected ',' or ']' in array")
		}
	}

	return list, nil
}

// 其他必要的解析方法...
func (p *parser) parseString() (interface{}, error) {
	token := p.next()
	return token.Value, nil
}

func (p *parser) parseNumber() (interface{}, error) {
	token := p.next()
	// 可以根据需要返回 int 或 float64
	num, err := strconv.ParseFloat(token.Value, 64)
	if err != nil {
		return nil, err
	}
	return num, nil
}

func (p *parser) next() Token {
	p.position++
	return p.tokens[p.position-1]

}

func (p *parser) peek() Token {
	return p.tokens[p.position]
}

func (p *parser) match(tokenType TokenType) bool {
	if p.peek().Type == tokenType {
		p.next()
		return true
	}
	return false
}
