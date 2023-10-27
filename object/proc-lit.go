package object

import (
	"fmt"
)

type ProcLitObject struct {
	Kind ObjectType
}

func (nlo *ProcLitObject) Type() ObjectType {
	return nlo.Kind
}

func (nlo *ProcLitObject) Debug() {
	fmt.Printf(" {Kind: Nil, Value: %v}\n")
}
