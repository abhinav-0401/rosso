package object

type ObjectType string

const (
	Int    = "integer"
	Bool   = "boolean"
	Nil    = "nil"
	Break  = "break"
	Return = "return"
	Proc   = "Proc"
)

type Object interface {
	Type() ObjectType
	Debug()
}
