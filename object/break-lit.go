package object

import (
	"fmt"
)

type BreakLitObject struct {
	Kind  ObjectType
	Value Object
}

func (blo *BreakLitObject) Type() ObjectType {
	return blo.Kind
}

func (blo *BreakLitObject) Debug() {
	fmt.Printf(" {Kind: BreakLitObject, Value: %v}\n", blo.Value)
}
