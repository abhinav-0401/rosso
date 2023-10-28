package parser

import (
	"fmt"
	"log"

	"github.com/abhinav-0401/rosso/ast"
	"github.com/abhinav-0401/rosso/token"
)

func (p *Parser) parseVarDeclStmt() ast.Stmt {
	var qualifier token.Token = p.eat()
	var isConst bool
	if qualifier.Type == token.CONST {
		isConst = true
	}

	var symbol token.Token = p.expect(token.IDENT, "")
	if p.at().Type != token.SEMICOLON {
		p.expect(token.ASSIGN, "") // don't need the "=" either
		var expr = p.parseExpr()
		p.expect(token.SEMICOLON, "")
		return &ast.VarDecl{Kind: ast.VarDeclNode, IsConstant: isConst, Symbol: symbol.Literal, Value: expr}
	} else {
		if isConst {
			log.Fatal("Error: constants must be assigned a value when being declared")
		}
		p.expect(token.SEMICOLON, "") // eat the semicolon
		return &ast.VarDecl{Kind: ast.VarDeclNode, IsConstant: isConst, Symbol: symbol.Literal, Value: nil}
	}
}

func (p *Parser) parseExprStmt(expectSemicolon bool) ast.Stmt {
	var expr = p.parseExpr()
	var exprStmt = &ast.ExprStmt{Kind: ast.ExprStmtNode, Node: expr}
	if (!expectSemicolon) && (p.at().Type == token.RBRACE) { // BlockExpr, last Expr can be an Expr
		return expr
	}
	if (expr.ExprKind() == ast.IfExprNode) || (expr.ExprKind() == ast.LoopExprNode) || (expr.ExprKind() == ast.BlockExprNode) {
		return exprStmt
	}
	p.expect(token.SEMICOLON, "")
	return exprStmt
}

func (p *Parser) parsePrintStmt() ast.Stmt {
	p.eat()
	var value = p.parseExpr()
	p.expect(token.SEMICOLON, "")
	return &ast.PrintStmt{Kind: ast.PrintStmtNode, Value: value}
}

func (p *Parser) parseBreakStmt() ast.Stmt {
	p.eat()
	fmt.Println(p.at())
	if p.at().Type == token.SEMICOLON {
		p.eat()
		return &ast.BreakStmt{Kind: ast.BreakStmtNode, Value: nil}
	}
	var value = p.parseExpr()
	p.expect(token.SEMICOLON, "")
	return &ast.BreakStmt{Kind: ast.BreakStmtNode, Value: value}
}

func (p *Parser) parseReturnStmt() *ast.ReturnStmt {
	p.eat()
	fmt.Println(p.at())
	if p.at().Type == token.SEMICOLON {
		p.eat()
		return &ast.ReturnStmt{Kind: ast.ReturnStmtNode, Value: nil}
	}
	var value = p.parseExpr()
	p.expect(token.SEMICOLON, "")
	return &ast.ReturnStmt{Kind: ast.ReturnStmtNode, Value: value}
}

func (p *Parser) parseStmt(expectSemicolon bool) ast.Stmt {
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
		return p.parseExprStmt(expectSemicolon)
	}
}

func (p *Parser) parseStmtInBlock() ast.Stmt {
	switch p.at().Type {
	case token.LET, token.CONST:
		return p.parseVarDeclStmt()
	case token.PRINT:
		return p.parsePrintStmt()
	case token.LBRACE:
		return p.parseBlockExpr()
	default:
		return p.parseExpr()
	}
}
