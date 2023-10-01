package eval

import (
	"fmt"
	"os"

	"github.com/abhinav-0401/rosso/ast"
	"github.com/abhinav-0401/rosso/object"
	"github.com/abhinav-0401/rosso/token"
)

func Eval(astNode ast.Stmt) object.Object {
	var obj object.Object
	switch node := astNode.(type) {
	case *ast.NumLit:
		return &object.NumLitObject{Kind: object.INT, Value: node.Value}
	case *ast.Program:
		return evalProgram(node)
	case *ast.BinaryExpr:
		return evalBinaryExpr(node)
	default:
		fmt.Printf("This AST Node has not yet been set up for evaluation\n")
		os.Exit(1)
	}
	return obj
}

func evalProgram(program *ast.Program) object.Object {
	var lastEvaluated object.Object = &object.NumLitObject{Kind: object.INT, Value: 0}

	for _, stmt := range program.Body {
		lastEvaluated = Eval(stmt)
	}

	return lastEvaluated
}

func evalBinaryExpr(binaryExpr *ast.BinaryExpr) object.Object {
	var lhs = Eval(binaryExpr.Left)
	var rhs = Eval(binaryExpr.Right)
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
