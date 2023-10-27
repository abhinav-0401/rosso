package eval

import (
	"fmt"
	"log"

	"github.com/abhinav-0401/rosso/ast"
	"github.com/abhinav-0401/rosso/env"
	"github.com/abhinav-0401/rosso/object"
	"github.com/abhinav-0401/rosso/token"
)

type BlockStmtEvalType int

const (
	Normal BlockStmtEvalType = iota
	Break
	Continue
	Return
)

func evalBinaryExpr(binaryExpr *ast.BinaryExpr, env *env.Env) object.Object {
	var lhs = Eval(binaryExpr.Left, env)
	var rhs = Eval(binaryExpr.Right, env)
	var operator = binaryExpr.Operator

	if lhs.Type() == object.Int && rhs.Type() == object.Int {
		l := lhs.(*object.NumLitObject)
		r := rhs.(*object.NumLitObject)
		switch operator.Type {
		case token.PLUS:
			return &object.NumLitObject{Kind: object.Int, Value: l.Value + r.Value}
		case token.MINUS:
			return &object.NumLitObject{Kind: object.Int, Value: l.Value - r.Value}
		case token.ASTERISK:
			return &object.NumLitObject{Kind: object.Int, Value: l.Value * r.Value}
		case token.SLASH:
			return &object.NumLitObject{Kind: object.Int, Value: l.Value / r.Value}
		case token.GT:
			return &object.BoolLitObject{Kind: object.Bool, Value: l.Value > r.Value}
		case token.LT:
			return &object.BoolLitObject{Kind: object.Bool, Value: l.Value < r.Value}
		case token.EQ:
			return &object.BoolLitObject{Kind: object.Bool, Value: l.Value == r.Value}
		case token.NOT_EQ:
			return &object.BoolLitObject{Kind: object.Bool, Value: l.Value != r.Value}
		}
	}
	return env.LookupVar("nil")
}

func evalIdent(ident *ast.Ident, env *env.Env) object.Object {
	// look up the value of the var by name and return it, scope auto resolved in eval.go
	return env.LookupVar(ident.Symbol)
}

func evalIfExpr(expr *ast.IfExpr, e *env.Env) object.Object {
	// an IfExpr will return nil if not evaluated at all, to prevent weird scenarios like:
	// let x = if 4 < 3 { 3 }
	var conditionValue = Eval(expr.Condition, e)
	if conditionValue.Type() != object.Bool {
		log.Fatal("Error: condition must be of bool type\n")
	}

	cv, _ := conditionValue.(*object.BoolLitObject)
	if cv.Value {
		return evalBlockStmt(expr.ThenBranch, e)
	} else if expr.ElseBranch != nil {
		return Eval(expr.ElseBranch, e)
	} else {
		return e.LookupVar("nil")
	}
}

func evalLoopExpr(loop *ast.LoopExpr, e *env.Env) object.Object {
	var loopValue = evalLoopBlockStmt(loop.Body, e).Value
	return loopValue
}

func evalBlockStmt(block *ast.BlockStmt, e *env.Env) object.Object {
	var lastEvaluated object.Object = e.LookupVar("nil")
	newEnv := &env.Env{Parent: e, Vars: make(map[string]object.Object), Consts: make(map[string]bool)}

	for _, stmt := range block.Body {
		fmt.Println(stmt.StmtKind())
		if stmt.StmtKind() == ast.BreakStmtNode {
			fmt.Println("this is a break statement")
			breakStmt, _ := stmt.(*ast.BreakStmt)
			return evalBreakStmt(breakStmt, newEnv)
		}
		lastEvaluated = Eval(stmt, newEnv)
	}

	return lastEvaluated
}

func evalLoopBlockStmt(block *ast.BlockStmt, e *env.Env) *object.BreakLitObject {
	var lastEvaluated object.Object = e.LookupVar("nil")
	newEnv := &env.Env{Parent: e, Vars: make(map[string]object.Object), Consts: make(map[string]bool)}

	for {
		for _, stmt := range block.Body {
			if stmt.StmtKind() == ast.BreakStmtNode {
				breakStmt, _ := stmt.(*ast.BreakStmt)
				return evalBreakStmt(breakStmt, newEnv)
			}
			lastEvaluated = Eval(stmt, newEnv)
			if lastEvaluated.Type() == object.Break {
				l := lastEvaluated.(*object.BreakLitObject)
				return l
			}
		}
	}
}
