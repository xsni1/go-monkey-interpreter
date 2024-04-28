package ast

import "github.com/xsni1/go-monkey-interpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// let a = 3;
type LetStatement struct {
	Token token.Token // LET token
	Name  *Identifier
	Value Expression
}

func (s *LetStatement) TokenLiteral() string {
	return s.Token.Literal
}
func (s *LetStatement) statementNode() {}

type Identifier struct {
	Token token.Token // IDENT token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) expressionNode() {}

type ReturnStatement struct {
	Token token.Token // RETURN
	Value Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// binary equality - == / !=
// binary comparison - > / >=
// binary factor - + / -
// binary term - * / /
// unary - ! / -
// literal

// 1 * 3 + 2
// 1 + 3 * 2

// func parseDodawanie() {
//   left := parseMnozenie()
//
//   for token == "+" {
//     right := parseMnozenie()
//     ast_node := ASTNode{left: left, right: right, sign: "+"}
//   }
//
//   return ast_node
// }

// func parseMnozenie() {
//   left := parseMnozenie()
//
//   for token == "+" {
//     right := parseMnozenie()
//     ast_node := ASTNode{left: left, right: right, sign: "+"}
//   }
//
//   return ast_node
// }
