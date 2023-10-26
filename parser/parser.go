package parser

import (
	"fmt"
	"os"

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
