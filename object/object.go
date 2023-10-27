package object

type ObjectType string

const (
	Int   = "integer"
	Bool  = "boolean"
	Nil   = "nil"
	Break = "break"
)

type Object interface {
	Type() ObjectType
	Debug()
}
