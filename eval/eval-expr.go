package eval

import (
	"github.com/abhinav-0401/rosso/ast"
	"github.com/abhinav-0401/rosso/env"
	"github.com/abhinav-0401/rosso/object"
	"github.com/abhinav-0401/rosso/token"
)

func evalBinaryExpr(binaryExpr *ast.BinaryExpr, env *env.Env) object.Object {
	var lhs = Eval(binaryExpr.Left, env)
	var rhs = Eval(binaryExpr.Right, env)
	var operator = binaryExpr.Operator

	if lhs.Type() == object.INT && rhs.Type() == object.INT {
		l := lhs.(*object.NumLitObject)
		r := rhs.(*object.NumLitObject)
		switch operator.Type {
		case token.PLUS:
			return &object.NumLitObject{Kind: object.INT, Value: l.Value + r.Value}
		case token.MINUS:
			return &object.NumLitObject{Kind: object.INT, Value: l.Value - r.Value}
		case token.ASTERISK:
			return &object.NumLitObject{Kind: object.INT, Value: l.Value * r.Value}
		case token.SLASH:
			return &object.NumLitObject{Kind: object.INT, Value: l.Value / r.Value}
		}
	}
	return &object.NumLitObject{Kind: object.INT, Value: 0}
}

func evalIdent(ident *ast.Ident, env *env.Env) object.Object {
	// look up the value of the var by name and return it, scope auto resolved in eval.go
	return env.LookupVar(ident.Symbol)
}
