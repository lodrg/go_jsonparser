package lpjsonparser

import (
	"strings"
	"unicode"
)

type lexer struct {
	input    []rune // 改用 rune 切片存储字符
	position int
	ch       rune // 改用 rune 存储当前字符
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

func (l *lexer) NextToken() Token {
	l.skipWhitespace()

	var tok Token
	switch l.ch {
	case '{':
		tok = Token{LEFT_BRACE, string(l.ch)}
	case '}':
		tok = Token{RIGHT_BRACE, string(l.ch)}
	case '[':
		tok = Token{LEFT_BRACKET, string(l.ch)}
	case ']':
		tok = Token{RIGHT_BRACKET, string(l.ch)}
	case ',':
		tok = Token{COMMA, string(l.ch)}
	case ':':
		tok = Token{COLON, string(l.ch)}
	case '"':
		tok = Token{STRING, l.readString()}
		l.readChar()
		return tok
	case 0: // EOF
		return Token{EOF, ""}
	default:
		if isDigit(l.ch) {
			return Token{NUMBER, l.readNumber()}
		}
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
	return sb.String()
}
