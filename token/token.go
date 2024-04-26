package token

type Token struct {
	Type    TokenType
	Literal string
}

// idk what's this for
type TokenType string

const (
	LET        = "LET"
	INT        = "INT"
	PLUS       = "PLUS"
	COMMA      = "COMMA"
	LPAREN     = "LAPREN"
	RPAREN     = "RPAREN"
	LBRACE     = "LBRACE"
	RBRACE     = "RBRACE"
	ASSIGN     = "ASSIGN"
	INTEGER    = "INTEGER"
	FUNCTION   = "FUNCTION"
	SEMICOLON  = "SEMICOLON"
	IDENTIFIER = "IDENTIFIER"
)
