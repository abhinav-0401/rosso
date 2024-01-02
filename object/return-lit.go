package object

import (
	"fmt"
)

type ReturnLitObject struct {
	Kind  ObjectType
	Value Object
}

func (rlo *ReturnLitObject) Type() ObjectType {
	return rlo.Kind
}

func (rlo *ReturnLitObject) Debug() {
	fmt.Printf(" {Kind: ReturnLitObject, Value: %v}\n", rlo.Value)
}
