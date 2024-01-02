package object

import (
	"fmt"
)

type NumLitObject struct {
	Kind  ObjectType
	Value int
}

func (nlo *NumLitObject) Type() ObjectType {
	return nlo.Kind
}

func (nlo *NumLitObject) Debug() {
	fmt.Printf("{Kind: Integer, Value: %v}]\n", nlo.Value)
}
