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

	switch l.ch {
	case '=':
		tok = token.Token{
			Type:    token.ASSIGN,
			Literal: "=",
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
	case 0:
	}
	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = '0'
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}
