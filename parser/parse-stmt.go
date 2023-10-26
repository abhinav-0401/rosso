package parser

import (
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

func (p *Parser) parseExprStmt() ast.Stmt {
	var expr = p.parseExpr()
	p.expect(token.SEMICOLON, "")
	return &ast.ExprStmt{Kind: ast.ExprStmtNode, Node: expr}
}

func (p *Parser) parsePrintStmt() ast.Stmt {
	p.eat()
	var value = p.parseExpr()
	p.expect(token.SEMICOLON, "")
	return &ast.PrintStmt{Kind: ast.PrintStmtNode, Value: value}
}

func (p *Parser) parseStmt() ast.Stmt {
	switch p.at().Type {
	case token.LET, token.CONST:
		return p.parseVarDeclStmt()
	case token.PRINT:
		return p.parsePrintStmt()
	default:
		return p.parseExpr()
	}
}
