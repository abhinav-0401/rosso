package eval

import (
	"log"

	"github.com/abhinav-0401/rosso/ast"
	"github.com/abhinav-0401/rosso/env"
	"github.com/abhinav-0401/rosso/object"
)

func evalProgram(program *ast.Program, env *env.Env) object.Object {
	var lastEvaluated object.Object = &object.NumLitObject{Kind: object.Int, Value: 0}

	for _, stmt := range program.Body {
		lastEvaluated = Eval(stmt, env)
	}

	return lastEvaluated
}

func evalVarDecl(decl *ast.VarDecl, env *env.Env) object.Object {
	var declValue = Eval(decl.Value, env)
	env.DeclareVar(decl.Symbol, declValue, decl.IsConstant)
	return env.LookupVar("nil")
}

func evalExprStmt(node *ast.ExprStmt, env *env.Env) object.Object {
	var value = Eval(node.Node, env)
	if value.Type() == object.Break {
		return value
	}
	return env.LookupVar("nil") // ExprStmt always return nil
}

func evalPrintStmt(stmt *ast.PrintStmt, env *env.Env) object.Object {
	Eval(stmt.Value, env).Debug()
	return env.LookupVar("nil")
}

func evalBreakStmt(stmt *ast.BreakStmt, e *env.Env) *object.BreakLitObject {
	if LoopCount == 0 {
		log.Fatal("Error: break statement outside enclosing loop")
	}
	LoopCount--
	if stmt.Value != nil {
		return &object.BreakLitObject{Kind: object.Break, Value: Eval(stmt.Value, e)}
	}
	return &object.BreakLitObject{Kind: object.Break, Value: e.LookupVar("nil")}
}
