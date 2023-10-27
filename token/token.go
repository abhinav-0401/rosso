package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	PROC  = "PROC"
	LET   = "LET"
	CONST = "CONST"
	// TRUE     = "TRUE"
	// FALSE    = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
	PRINT  = "PRINT"
	LOOP   = "LOOP"
	BREAK  = "BREAK"
)

var keywords = map[string]TokenType{
	"proc":  PROC,
	"let":   LET,
	"const": CONST,
	// "true":   TRUE,
	// "false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"print":  PRINT,
	"loop":   LOOP,
	"break":  BREAK,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
