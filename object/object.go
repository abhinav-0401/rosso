package object

type ObjectType string

const (
	Int  = "integer"
	Bool = "boolean"
	Nil  = "nil"
)

type Object interface {
	Type() ObjectType
	Debug()
}
