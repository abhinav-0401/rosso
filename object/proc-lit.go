package object

import (
	"fmt"

	"github.com/abhinav-0401/rosso/ast"
	// "github.com/abhinav-0401/rosso/object/env"
)

type ProcLitObject struct {
	Kind   ObjectType
	Params []*ast.Ident
	Body   *ast.BlockStmt
	Env    Enver
}

type Enver interface {
	Idk() // to circumvent the, frankly stupid, "circular dependency" error when there is none
}

func (nlo *ProcLitObject) Type() ObjectType {
	return nlo.Kind
}

func (nlo *ProcLitObject) Debug() {
	fmt.Printf("{Kind: Nil, Value:}\n")
}
