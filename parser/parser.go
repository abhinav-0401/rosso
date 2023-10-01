package parser

import (
	"fmt"
	"os"
	"strconv"

	"github.com/abhinav-0401/rosso/ast"
	"github.com/abhinav-0401/rosso/lexer"
	"github.com/abhinav-0401/rosso/token"
)

type Parser struct {
	Tokens  []token.Token
	current int
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) parseAdditiveExpr() ast.Expr {
	var left = p.parseMultiplicativeExpr()

	for p.at().Type == token.PLUS || p.at().Type == token.MINUS {
		var operator = p.eat()
		var right = p.parseMultiplicativeExpr()
		var binaryExpr = &ast.BinaryExpr{Kind: ast.BinaryExprNode, Left: left, Right: right, Operator: operator}
		left = binaryExpr
	}

	return left
}

func (p *Parser) parseMultiplicativeExpr() ast.Expr {
	var left = p.parsePrimaryExpr()

	for p.at().Type == token.ASTERISK || p.at().Type == token.SLASH {
		var operator = p.eat()
		var right = p.parsePrimaryExpr()
		var binaryExpr = &ast.BinaryExpr{Kind: ast.BinaryExprNode, Left: left, Right: right, Operator: operator}
		left = binaryExpr
	}

	return left
}

func (p *Parser) parsePrimaryExpr() ast.Expr {
	var tok = p.at()
	var expr ast.Expr

	switch tok.Type {
	case token.IDENT:
		return &ast.Ident{Kind: ast.IdentNode, Symbol: p.eat().Literal}
	case token.INT:
		numLit, _ := strconv.Atoi(p.eat().Literal)
		return &ast.NumLit{Kind: ast.NumLitNode, Value: numLit}
	case token.LPAREN:
		p.eat()
		expr := p.parseExpr()
		p.eat()
		return expr
	default:
		fmt.Print("eof?")
		os.Exit(1)
	}

	return expr
}

func (p *Parser) parseExpr() ast.Expr {
	return p.parseAdditiveExpr()
}

func (p *Parser) parseStmt() ast.Stmt {
	return p.parseExpr()
}

func (p *Parser) ProduceAst(src string) *ast.Program {
	var program = &ast.Program{
		Kind: ast.ProgramNode,
	}

	lexer := lexer.New(src)
	p.Tokens = lexer.Tokenise()
	fmt.Printf("tokens in the parser: %v\n", p.Tokens)

	// parse until you hit the EOF token
	for p.notEof() {
		program.Body = append(program.Body, p.parseStmt())
	}

	return program
}

func (p *Parser) notEof() bool {
	return p.Tokens[p.current].Type != token.EOF
}

func (p *Parser) at() token.Token {
	return p.Tokens[p.current]
}

func (p *Parser) eat() token.Token {
	prev := p.Tokens[p.current]
	p.current++
	return prev
}
