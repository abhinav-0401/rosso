package object

import (
	"fmt"
)

type NilLitObject struct {
	Kind  ObjectType
	Value interface{} // hack to represent the null type for now
}

func (nlo *NilLitObject) Type() ObjectType {
	return nlo.Kind
}

func (nlo *NilLitObject) Debug() {
	fmt.Printf(" {Kind: Nil, Value: %v}\n", nlo.Value)
}
