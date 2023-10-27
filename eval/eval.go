package eval

import (
	"fmt"
	"log"
	"os"

	"github.com/abhinav-0401/rosso/ast"
	"github.com/abhinav-0401/rosso/env"
	"github.com/abhinav-0401/rosso/object"
)

var LoopCount = 0

func Eval(astNode ast.Stmt, env *env.Env) object.Object {
	var obj object.Object
	switch node := astNode.(type) {
	case *ast.NumLit:
		return &object.NumLitObject{Kind: object.Int, Value: node.Value}
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.Ident:
		return evalIdent(node, env)
	case *ast.BinaryExpr:
		return evalBinaryExpr(node, env)
	case *ast.VarDecl:
		return evalVarDecl(node, env)
	case *ast.ExprStmt:
		return evalExprStmt(node, env)
	case *ast.PrintStmt:
		return evalPrintStmt(node, env)
	case *ast.IfExpr:
		return evalIfExpr(node, env)
	case *ast.BlockStmt:
		return evalBlockStmt(node, env)
	case *ast.LoopExpr:
		LoopCount++
		return evalLoopExpr(node, env)
	case *ast.BreakStmt:
		// Eval(*ast.Program) will directly call this
		log.Fatal("Error: break statement outside enclosing loop")
		// return evalBreakStmt(node, env)
	default:
		fmt.Printf("This AST Node has not yet been set up for evaluation\n")
		os.Exit(1)
	}
	return obj
}
