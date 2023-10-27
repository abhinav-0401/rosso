package ast

import (
	"github.com/abhinav-0401/rosso/token"
)

type NodeType string

const (
	ProgramNode     = "Program"
	NumLitNode      = "NumericLiteral"
	IdentNode       = "IdentNode"
	BinaryExprNode  = "BinaryExprNode"
	VarDeclNode     = "VarDeclNode"
	IfExprNode      = "IfExprNode"
	ExprStmtNode    = "ExprStmtNode"
	PrintStmtNode   = "PrintStmtNode"
	BlockStmtNode   = "BlockStmtNode"
	BlockExprNode   = "BlockExprNode"
	LoopExprNode    = "LoopExprNode"
	BreakStmtNode   = "BreakStmtNode"
	ProcLitExprNode = "ProcLitExprNode"
)

type Stmt interface {
	StmtKind() NodeType
}

type Expr interface {
	Stmt
	ExprKind() NodeType // method literally only exists to distinguish this type from Stmt
}

type ConditionalStmt interface {
	Stmt
	ConditionalKind() NodeType
}

type Program struct {
	Kind NodeType
	Body []Stmt
}

func (p *Program) StmtKind() NodeType {
	return ProgramNode

}

type BlockStmt struct {
	Kind   NodeType
	Body   []Stmt
	Last   Expr
	IsLoop bool
	// Value  Expr
}

func (bs *BlockStmt) StmtKind() NodeType {
	return BlockExprNode

}

func (bs *BlockStmt) ExprKind() NodeType { // this makes the struct an Expr
	return BlockExprNode
}

func (bs *BlockStmt) ConditionalKind() NodeType {
	return BlockExprNode

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

type BreakStmt struct {
	Kind  NodeType
	Value Expr
}

func (bs *BreakStmt) StmtKind() NodeType {
	return BreakStmtNode
}

type ExprStmt struct {
	Kind NodeType
	Node Expr
}

func (es *ExprStmt) StmtKind() NodeType {
	return ExprStmtNode
}

type PrintStmt struct {
	Kind  NodeType
	Value Expr
}

func (ps *PrintStmt) StmtKind() NodeType {
	return PrintStmtNode
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

type IfExpr struct {
	Kind       NodeType
	ThenBranch *BlockStmt
	ElseBranch ConditionalStmt
	Condition  Expr
}

func (ie *IfExpr) StmtKind() NodeType {
	return IfExprNode
}

func (ie *IfExpr) ExprKind() NodeType {
	return IfExprNode
}

func (ie *IfExpr) ConditionalKind() NodeType {
	return IfExprNode
}

type LoopExpr struct {
	Kind NodeType
	Body *BlockStmt
}

func (le *LoopExpr) StmtKind() NodeType {
	return LoopExprNode
}

func (le *LoopExpr) ExprKind() NodeType {
	return LoopExprNode
}

type ProcLitExpr struct {
	Kind   NodeType
	Params []*Ident
	Body   *BlockStmt
}

func (ple *ProcLitExpr) StmtKind() NodeType {
	return LoopExprNode
}

func (ple *ProcLitExpr) ExprKind() NodeType {
	return LoopExprNode
}

// type CallExpr struct {
// 	Kind NodeType
// 	Proc ProcExpr // Either an Ident or a ProcLitExpr, since both can be called

// }
