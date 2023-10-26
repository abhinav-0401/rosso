package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/abhinav-0401/rosso/ast"
	"github.com/abhinav-0401/rosso/token"
)

func (p *Parser) parseComparisonExpr() ast.Expr {
	var left = p.parseAdditiveExpr()

	for p.at().Type == token.GT || p.at().Type == token.LT || p.at().Type == token.EQ || p.at().Type == token.NOT_EQ {
		var operator = p.eat()
		var right = p.parseAdditiveExpr()
		var binaryExpr = &ast.BinaryExpr{Kind: ast.BinaryExprNode, Left: left, Right: right, Operator: operator}
		left = binaryExpr
	}

	return left
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
	case token.IF: // Expr are Stmts :)
		return p.parseIfExpr()
	default:
		fmt.Print("eof? ")
		os.Exit(1)
	}

	return expr
}

func (p *Parser) parseExpr() ast.Expr {
	return p.parseComparisonExpr()
}

func (p *Parser) parseIfExpr() ast.Expr {
	p.eat()

	var condition = p.parseExpr()
	var thenBranch []ast.Stmt
	var elseBranch ast.Stmt

	p.expect(token.LBRACE, "")
	for p.at().Type != token.RBRACE {
		thenBranch = append(thenBranch, p.parseStmt())
	}
	p.eat() // getting rid of the RBRACE at the end, always gets me whenever i make an interpreter
	if p.at().Type == token.ELSE {
		// gotta handle two things for else
		// if, and simple else { }
		p.eat()
		elseBranch = p.parseStmt()
	}

	str, _ := json.MarshalIndent(ast.IfExpr{Kind: ast.IfExprNode, ThenBranch: thenBranch, ElseBranch: elseBranch, Condition: condition}, "", "    ")
	fmt.Print(string(str))

	return &ast.IfExpr{Kind: ast.IfExprNode, ThenBranch: thenBranch, ElseBranch: elseBranch, Condition: condition}
}
