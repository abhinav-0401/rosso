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
		return p.parseIdentExpr()
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
	case token.PROC:
		// this could either be just a proc literal or a proc IIFE
		fmt.Println("inside this case thingy")
		var procExpr = p.parseProcExpr()
		fmt.Println("procExpr: ", procExpr)
		fmt.Println(p.at())
		if p.at().Type == token.LPAREN { // CallExpr
			return p.parseCallExpr(procExpr)
		}
		return procExpr
	default:
		fmt.Print("eof? ")
		os.Exit(1)
	}

	return expr
}

func (p *Parser) parseExpr() ast.Expr {
	return p.parseComparisonExpr()
}

func (p *Parser) parseIdentExpr() ast.Expr {
	// this could also just either be an ident, an assignment or a proc call
	var symbolToken = p.eat()
	var ident = &ast.Ident{Kind: ast.IdentNode, Symbol: symbolToken.Literal}
	switch p.at().Type {
	case token.LPAREN: // CallExpr
		return p.parseCallExpr(ident)
	case token.ASSIGN:
		log.Fatal("Error: AST node for assignment not set up for parsing yet")
	}
	return ident
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
		body = append(body, p.parseStmt(false))
	}
	p.eat() // }
	return &ast.BlockStmt{Kind: ast.BlockExprNode, Body: body}
}

func (p *Parser) parseProcExpr() *ast.ProcLit {
	p.eat()
	p.expect(token.LPAREN, "")
	var params = p.parseParams()
	p.expect(token.RPAREN, "")
	var body = p.parseBlockExpr()
	return &ast.ProcLit{Kind: ast.ProcLitNode, Params: params, Body: body}
}

func (p *Parser) parseCallExpr(procExpr ast.ProcExpr) *ast.CallExpr {
	p.eat() // (
	var args = p.parseArgs()
	p.expect(token.RPAREN, "")
	return &ast.CallExpr{Kind: ast.CallExprNode, Proc: procExpr, Args: args}
}

func (p *Parser) parseParams() []*ast.Ident {
	var params = []*ast.Ident{}
	if p.at().Type == token.RPAREN {
		return params
	}
	params = append(params, p.parseIdent())

	for p.at().Type == token.COMMA {
		p.eat()
		params = append(params, p.parseIdent())
	}
	return params
}

func (p *Parser) parseArgs() []ast.Expr {
	var args = []ast.Expr{}
	if p.at().Type == token.RPAREN {
		return args
	}
	args = append(args, p.parseExpr())
	fmt.Println("inside parseArgs")

	for p.at().Type == token.COMMA {
		p.eat()
		args = append(args, p.parseExpr())
	}
	return args
}

func (p *Parser) parseIdent() *ast.Ident {
	return &ast.Ident{Kind: ast.IdentNode, Symbol: p.eat().Literal}
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
	case token.RETURN:
		return p.parseReturnStmt()
	default:
		return p.parseExprStmt(false)
	}
}
