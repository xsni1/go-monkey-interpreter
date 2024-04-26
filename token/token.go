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
	FUNCTION   = "FUNCTION"
	SEMICOLON  = "SEMICOLON"
	IDENTIFIER = "IDENTIFIER"
	BANG       = "BANG"
	LT         = "<"
	GT         = ">"
	MINUS      = "MINUS"
	SLASH      = "SLASH"
	ASTERISK   = "ASTERISK"
	EQ         = "EQ"
	NOT_EQ     = "NOT_EQ"

	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"

	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if toktype, ok := keywords[ident]; ok {
		return toktype
	}
	return IDENTIFIER
}
