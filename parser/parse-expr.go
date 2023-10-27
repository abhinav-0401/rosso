package parser

import (
	"fmt"
	"log"
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
	case token.LOOP:
		return p.parseLoopExpr()
	case token.LBRACE:
		return p.parseBlockExpr()
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
	var thenBranch = &ast.BlockStmt{Kind: ast.BlockExprNode, Body: nil}
	var elseBranch ast.ConditionalStmt

	thenBranch = p.parseBlockExpr()
	if p.at().Type == token.ELSE {
		// gotta handle two things for else
		// if, and simple else { }
		p.eat()
		switch p.at().Type {
		case token.LBRACE: // BlockStmt
			elseBranch = p.parseBlockExpr()
		case token.IF:
			elseBranch = p.parseIfExpr().(*ast.IfExpr)
		default:
			log.Fatalf("Error: else must be followed by another if expression or a block statement")
		}
	}

	return &ast.IfExpr{Kind: ast.IfExprNode, ThenBranch: thenBranch, ElseBranch: elseBranch, Condition: condition}
}

func (p *Parser) parseLoopExpr() ast.Expr {
	p.eat()
	var body = p.parseBlockExpr()
	body.IsLoop = true
	return &ast.LoopExpr{Kind: ast.LoopExprNode, Body: body}
}

// actually an expr
func (p *Parser) parseBlockExpr() *ast.BlockStmt {
	p.eat() // {
	var body []ast.Stmt
	for p.at().Type != token.RBRACE { // this is not the logic i want
		body = append(body, p.parseBlockExprStmt())
	}
	p.eat() // }
	return &ast.BlockStmt{Kind: ast.BlockExprNode, Body: body}
}

// helper to parseBlockExpr()
func (p *Parser) parseBlockExprStmt() ast.Stmt {
	switch p.at().Type {
	case token.LET, token.CONST:
		return p.parseVarDeclStmt()
	case token.PRINT:
		return p.parsePrintStmt()
	case token.BREAK:
		return p.parseBreakStmt()
	default:
		return p.parseExprStmt(false)
	}
}
