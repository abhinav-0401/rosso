package object

import (
	"fmt"
)

type BoolLitObject struct {
	Kind  ObjectType
	Value bool
}

func (blo *BoolLitObject) Type() ObjectType {
	return blo.Kind
}

func (blo *BoolLitObject) Debug() {
	fmt.Printf(" {Kind: Boolean, Value: %v}\n", blo.Value)
}
