package eval

import (
	"fmt"
	"os"

	"github.com/abhinav-0401/rosso/ast"
	"github.com/abhinav-0401/rosso/env"
	"github.com/abhinav-0401/rosso/object"
)

func Eval(astNode ast.Stmt, env *env.Env) object.Object {
	var obj object.Object
	switch node := astNode.(type) {
	case *ast.NumLit:
		return &object.NumLitObject{Kind: object.INT, Value: node.Value}
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.Ident:
		return evalIdent(node, env)
	case *ast.BinaryExpr:
		return evalBinaryExpr(node, env)
	case *ast.VarDecl:
		return evalVarDecl(node, env)
	case *ast.VarAssign:
		return evalVarAssign(node, env)
	default:
		fmt.Printf("This AST Node has not yet been set up for evaluation\n")
		os.Exit(1)
	}
	return obj
}
