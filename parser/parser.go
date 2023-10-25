package parser

import (
	"fmt"
	"log"
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
		return p.parseIdent()
	case token.INT:
		numLit, _ := strconv.Atoi(p.eat().Literal)
		return &ast.NumLit{Kind: ast.NumLitNode, Value: numLit}
	case token.LPAREN:
		p.eat()
		expr := p.parseExpr()
		p.eat()
		return expr
	default:
		fmt.Print("eof? ")
		os.Exit(1)
	}

	return expr
}

func (p *Parser) parseIdent() ast.Expr {
	var ident token.Token = p.eat()
	if p.at().Type == token.ASSIGN {
		p.eat()
		var expr = p.parseExpr()
		p.expect(token.SEMICOLON, "")
		return &ast.VarAssign{Kind: ast.VarAssignNode, Symbol: ident.Literal, Value: expr}
	}
	return &ast.Ident{Kind: ast.IdentNode, Symbol: ident.Literal}
}

func (p *Parser) parseExpr() ast.Expr {
	return p.parseAdditiveExpr()
}

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

func (p *Parser) parseStmt() ast.Stmt {
	switch p.at().Type {
	case token.LET, token.CONST:
		return p.parseVarDeclStmt()
	default:
		return p.parseExpr()
	}
}

func (p *Parser) ProduceAst(src string) *ast.Program {
	var program = &ast.Program{
		Kind: ast.ProgramNode,
	}

	lex := lexer.New(src)
	p.Tokens = lex.Tokenise()
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

func (p *Parser) expect(tokenType token.TokenType, msg string) token.Token {
	if p.at().Type != tokenType {
		if msg == "" {
			msg = fmt.Sprintf("Error: expected token %v, found %v\n", tokenType, p.at().Literal)
		}
		fmt.Printf(msg)
		os.Exit(1)
	}
	return p.eat()
}
