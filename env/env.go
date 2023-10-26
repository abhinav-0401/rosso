package env

import (
	"fmt"
	"log"
	"os"

	"github.com/abhinav-0401/rosso/object"
)

type Env struct {
	Parent *Env
	Vars   map[string]object.Object
	Consts map[string]bool
}

func New(parentEnv *Env) *Env {
	return &Env{Parent: parentEnv, Vars: make(map[string]object.Object), Consts: make(map[string]bool)}
}

func (e *Env) DeclareVar(name string, value object.Object, isConst bool) object.Object {
	// we want to check if the var has already been declared
	if _, ok := e.Vars[name]; ok {
		fmt.Printf("Error: cannot declare %v as the variable already exists in scope\n", name)
		os.Exit(1)
	}

	// otherwise just set the variable equal to the value
	e.Vars[name] = value
	if isConst {
		e.Consts[name] = true
	}
	return value
}

func (e *Env) AssignVar(name string, value object.Object) object.Object {
	// make sure that it exists
	// first we resolve the env that the var belongs to
	var env = e.resolve(name)

	// if it is a constant, throw an error
	if e.Consts[name] {
		log.Fatal("Error: cannot assign to a constant\n")
	}
	env.Vars[name] = value

	return value
}

func (e *Env) LookupVar(name string) object.Object {
	var env = e.resolve(name)
	return env.Vars[name]
}

func (e *Env) resolve(name string) *Env {
	// check if the current env has this var
	if _, ok := e.Vars[name]; ok {
		return e
	}

	if e.Parent == nil {
		fmt.Printf("Error: cannot assign variable %v as it does not exist", name)
		os.Exit(1)
	}

	// otherwise, call resolve on the parent
	return e.Parent.resolve(name)
}
