package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/abhinav-0401/rosso/env"
	"github.com/abhinav-0401/rosso/eval"
	"github.com/abhinav-0401/rosso/object"
	"github.com/abhinav-0401/rosso/parser"
	"github.com/abhinav-0401/rosso/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		fmt.Printf("Hello %s! This is the Rosso programming language!\n",
			user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}

	if len(os.Args) > 2 {
		log.Fatal("Error: unsupported arguments")
	}

	var fileName = os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}
	reader := bufio.NewReader(file)
	src, _ := reader.ReadString(0)

	var parser = parser.New()
	var ast = parser.ProduceAst(src)
	var e = env.New(nil)
	e.DeclareVar("PI", &object.NumLitObject{Kind: object.Int, Value: 4}, true)
	e.DeclareVar("true", &object.BoolLitObject{Kind: object.Bool, Value: true}, true)
	e.DeclareVar("false", &object.BoolLitObject{Kind: object.Bool, Value: false}, true)
	e.DeclareVar("nil", &object.NilLitObject{Kind: object.Nil, Value: nil}, true)
	eval.Eval(ast, e)
	// fmt.Println(value)
	// valuePretty, _ := json.MarshalIndent(value, "", "    ")
	// fmt.Printf("%+v\n", string(valuePretty))
}
