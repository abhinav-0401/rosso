package ast

import (
	"github.com/abhinav-0401/rosso/token"
)

type NodeType string

const (
	ProgramNode    = "Program"
	NumLitNode     = "NumericLiteral"
	IdentNode      = "IdentNode"
	BinaryExprNode = "BinaryExprNode"
	VarDeclNode    = "VarDeclNode"
	IfExprNode     = "IfExprNode"
	ExprStmtNode   = "ExprStmtNode"
	PrintStmtNode  = "PrintStmtNode"
	BlockStmtNode  = "BlockStmtNode"
	BlockExprNode  = "BlockExprNode"
	LoopExprNode   = "LoopExprNode"
	BreakStmtNode  = "BreakStmtNode"
	ProcLitNode    = "ProcLitExprNode"
	CallExprNode   = "CallExprNode"
	ReturnStmtNode = "ReturnStmtNode"
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

type ProcExpr interface {
	Stmt
	Expr
	ProcExprKind() NodeType
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

// a ReturnStmt is just a BreakStmt, but insteak of evalLoopExpr catching it, evalProcExpr will catch it :)
// copy-paste code ftw
type ReturnStmt struct {
	Kind  NodeType
	Value Expr
}

func (rs *ReturnStmt) StmtKind() NodeType {
	return ReturnStmtNode
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

func (i *Ident) ProcExprKind() NodeType {
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

type ProcLit struct {
	Kind   NodeType
	Params []*Ident
	Body   *BlockStmt
}

func (ple *ProcLit) StmtKind() NodeType {
	return ProcLitNode
}

func (ple *ProcLit) ExprKind() NodeType {
	return ProcLitNode
}

func (ple *ProcLit) ProcExprKind() NodeType {
	return ProcLitNode
}

type CallExpr struct {
	Kind NodeType
	Proc ProcExpr // Either an Ident or a ProcLit, since both can be called
	Args []Expr
}

func (ce *CallExpr) StmtKind() NodeType {
	return CallExprNode
}

func (ce *CallExpr) ExprKind() NodeType {
	return CallExprNode
}
