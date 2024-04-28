package lexer

import "github.com/xsni1/go-monkey-interpreter/token"

type Lexer struct {
	input        string
	readPosition int
	position     int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespaces()

	switch l.ch {
	case '=':
		if l.PeekChar() == '=' {
			tok = token.Token{
				Type:    token.EQ,
				Literal: "==",
			}
			l.readChar()
		} else {
			tok = token.Token{
				Type:    token.ASSIGN,
				Literal: "=",
			}
		}

	case '+':
		tok = token.Token{
			Type:    token.PLUS,
			Literal: "+",
		}
	case '(':
		tok = token.Token{
			Type:    token.LPAREN,
			Literal: "(",
		}
	case ')':
		tok = token.Token{
			Type:    token.RPAREN,
			Literal: ")",
		}
	case '{':
		tok = token.Token{
			Type:    token.LBRACE,
			Literal: "{",
		}
	case '}':
		tok = token.Token{
			Type:    token.RBRACE,
			Literal: "}",
		}
	case ',':
		tok = token.Token{
			Type:    token.COMMA,
			Literal: ",",
		}
	case ';':
		tok = token.Token{
			Type:    token.SEMICOLON,
			Literal: ";",
		}
	case '!':
		if l.PeekChar() == '=' {
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: "!=",
			}
			l.readChar()
		} else {
			tok = token.Token{
				Type:    token.BANG,
				Literal: "!",
			}
		}
	case '/':
		tok = token.Token{
			Type:    token.SLASH,
			Literal: "/",
		}
	case '-':
		tok = token.Token{
			Type:    token.MINUS,
			Literal: "-",
		}
	case '*':
		tok = token.Token{
			Type:    token.ASTERISK,
			Literal: "*",
		}
	case '<':
		tok = token.Token{
			Type:    token.LT,
			Literal: "<",
		}
	case '>':
		tok = token.Token{
			Type:    token.GT,
			Literal: ">",
		}
    case 0:
		tok = token.Token{
			Type:    token.EOF,
			Literal: "",
		}
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			tok = token.Token{
				Type:    token.LookupIdent(ident),
				Literal: ident,
			}
			return tok
		}

		if isDigit(l.ch) {
			return token.Token{
				Type:    token.INT,
				Literal: l.readNumber(),
			}
		}

		tok = token.Token{
			Type:    token.ILLEGAL,
			Literal: string(l.ch),
		}
	}
	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readIdentifier() string {
	startpos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startpos:l.position]
}

func (l *Lexer) readNumber() string {
	startpos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[startpos:l.position]
}

func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.readChar()
	}
}

func (l *Lexer) PeekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
