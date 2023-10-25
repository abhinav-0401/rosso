package ast

import (
	"github.com/abhinav-0401/rosso/token"
)

type NodeType string

const (
	ProgramNode    = "Program"
	NumLitNode     = "NumericLiteral"
	IdentNode      = "Identifier"
	BinaryExprNode = "BinaryExpr"
	VarDeclNode    = "VarDeclNode"
)

type Stmt interface {
	StmtKind() NodeType
}

type Expr interface {
	Stmt
	ExprKind() NodeType // method literally only exists to distinguish this type from Stmt
}

type Program struct {
	Kind NodeType
	Body []Stmt
}

func (p *Program) StmtKind() NodeType {
	return ProgramNode

}

type VarDecl struct {
	Kind       NodeType
	IsConstant bool
	Symbol     string
	Value      Expr
}

func (vd *VarDecl) StmtKind() NodeType {
	return vd.Kind
}

// -----------------------------------------------
//	STATEMENTS ABOVE, EXPRESSIONS BELOW
// -----------------------------------------------

type BinaryExpr struct {
	Kind     NodeType
	Left     Expr
	Right    Expr
	Operator token.Token
}

func (be *BinaryExpr) StmtKind() NodeType {
	return BinaryExprNode
}
func (be *BinaryExpr) ExprKind() NodeType {
	return BinaryExprNode
}

type Ident struct {
	Kind   NodeType
	Symbol string
}

func (i *Ident) StmtKind() NodeType {
	return IdentNode
}
func (i *Ident) ExprKind() NodeType {
	return IdentNode
}

type NumLit struct {
	Kind  NodeType
	Value int
}

func (nl *NumLit) StmtKind() NodeType {
	return NumLitNode
}
func (nl *NumLit) ExprKind() NodeType {
	return NumLitNode
}
