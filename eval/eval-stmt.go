package eval

import (
	"github.com/abhinav-0401/rosso/ast"
	"github.com/abhinav-0401/rosso/env"
	"github.com/abhinav-0401/rosso/object"
)

func evalProgram(program *ast.Program, env *env.Env) object.Object {
	var lastEvaluated object.Object = &object.NumLitObject{Kind: object.INT, Value: 0}

	for _, stmt := range program.Body {
		lastEvaluated = Eval(stmt, env)
	}

	return lastEvaluated
}

func evalVarDecl(decl *ast.VarDecl, env *env.Env) object.Object {
	var declValue = Eval(decl.Value, env)
	env.DeclareVar(decl.Symbol, declValue, decl.IsConstant)
	return declValue
}
