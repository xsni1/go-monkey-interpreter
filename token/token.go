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
	EOF        = "EOF"
	ILLEGAL    = "ILLEGAL"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if toktype, ok := keywords[ident]; ok {
		return toktype
	}
	return IDENTIFIER
}
