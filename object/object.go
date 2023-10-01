package object

type ObjectType string

const (
	INT = "integer"
)

type Object interface {
	Type() ObjectType
	Debug()
}
