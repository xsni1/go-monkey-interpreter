package parser

import (
	"fmt"
	"strconv"

	"github.com/xsni1/go-monkey-interpreter/ast"
	"github.com/xsni1/go-monkey-interpreter/lexer"
	"github.com/xsni1/go-monkey-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	tok := p.l.NextToken()
	p.curToken = p.peekToken
	p.peekToken = tok
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseStatement() ast.Statement {
	if p.curToken.Type == token.LET {
		return p.parseLetStmt()
	} else if p.curToken.Type == token.RETURN {
		return p.parseReturnStmt()
	}

	return nil
}

func (p *Parser) parseLetStmt() *ast.LetStatement {
	stmt := &ast.LetStatement{
		Token: p.curToken,
	}

	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStmt() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{
		Token: p.curToken,
	}

	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) ParseExpression() ast.Expression {
	return p.parseExp()
}

func (p *Parser) parseExp() ast.Expression {
	return p.parseTerm()
}

func (p *Parser) parseTerm() ast.Expression {
	expr := p.parseFactor()

	for p.curToken.Type == token.PLUS {
		p.nextToken()
		expr = &ast.ExpressionNode{
			Left:  expr,
            Operator: "+",
			Right: p.parseFactor(),
		}
	}

	return expr
}

func (p *Parser) parseFactor() ast.Expression {
	expr := p.parseLiteral()

	for p.curToken.Type == token.ASTERISK {
		p.nextToken()
		expr = &ast.ExpressionNode{
			Left:  expr,
            Operator: "*",
			Right: p.parseLiteral(),
		}
	}

	return expr
}

func (p *Parser) parseLiteral() ast.Expression {
	if p.curToken.Type != token.INT {
		p.errors = append(p.errors, "wrong int")
		return nil
	}

	val, err := strconv.Atoi(p.curToken.Literal)
	if err != nil {
		p.errors = append(p.errors, "error converting int")
		return nil
	}
	node := &ast.IntegerLiteral{
		Token: p.curToken,
		Value: int64(val),
	}

	p.nextToken()

	return node
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

func (p *Parser) peekError(t token.TokenType) {
	p.errors = append(p.errors, fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type))
}
