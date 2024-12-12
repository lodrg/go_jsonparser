package lpjsonparser

import (
	"strings"
	"unicode"
)

// rune
type lexer struct {
	input    []rune // 改用 rune 切片存储字符
	position int
	ch       rune // 改用 rune 存储当前字符
}

var SingleCharTokens = map[rune]TokenType{
	'{': LEFT_BRACKET,
	'}': RIGHT_BRACKET,
	'[': LEFT_BRACKET,
	']': RIGHT_BRACKET,
	',': COMMA,
	';': COLON,
}

func NewLexer(input string) *lexer {
	l := &lexer{input: []rune(input)}
	l.readChar()
	return l
}

func (l *lexer) readChar() {
	if l.position >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.position]
	}
	l.position++
}

// NextToken is the core func
func (l *lexer) NextToken() Token {
	l.skipWhitespace()
	var tok Token

	if l.ch == 0 {
		return Token{EOF, ""}
	}

	if tokenType, ok := SingleCharTokens[l.ch]; ok {
		tok = Token{Type: tokenType}
		l.readChar()
		return tok
	}

	switch {
	case l.ch == '"':
		tok = Token{STRING, l.readString()}
	case isDigit(l.ch):
		tok = Token{NUMBER, l.readNumber()}
	default:
		tok = Token{ILLEGAL, string(l.ch)}
	}

	l.readChar()
	return tok
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func (l *lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func (l *lexer) readNumber() string {
	var sb strings.Builder
	for isDigit(l.ch) {
		sb.WriteRune(l.ch)
		l.readChar()
	}
	return sb.String()
}

func (l *lexer) readString() string {
	var sb strings.Builder
	l.readChar() // 跳过开始的引号

	for l.ch != 0 && l.ch != '"' {
		if l.ch == '\\' {
			l.readChar()
			switch l.ch {
			case 'n':
				sb.WriteRune('\n')
			case 'r':
				sb.WriteRune('\r')
			case 't':
				sb.WriteRune('\t')
			case '\\':
				sb.WriteRune('\\')
			case '"':
				sb.WriteRune('"')
			case 'u':
				// TODO: 处理 Unicode 转义序列 \uXXXX
				// 这部分可以根据需要添加
			default:
				sb.WriteRune(l.ch)
			}
		} else {
			sb.WriteRune(l.ch)
		}
		l.readChar()
	}
	// 跳过结束的引号
	l.readChar()
	return sb.String()
}
