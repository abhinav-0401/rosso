package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/abhinav-0401/rosso/eval"

	"github.com/abhinav-0401/rosso/env"
	"github.com/abhinav-0401/rosso/object"
	"github.com/abhinav-0401/rosso/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	var e = env.New(nil)
	e.DeclareVar("PI", &object.NumLitObject{Kind: object.Int, Value: 4}, true)
	e.DeclareVar("true", &object.BoolLitObject{Kind: object.Bool, Value: true}, true)
	e.DeclareVar("false", &object.BoolLitObject{Kind: object.Bool, Value: false}, true)
	e.DeclareVar("nil", &object.NilLitObject{Kind: object.Nil, Value: nil}, true)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		parse := parser.New()
		program := parse.ProduceAst(line)
		programPretty, _ := json.MarshalIndent(program, "", "    ")
		fmt.Printf("%+v\n", string(programPretty))

		value := eval.Eval(program, e)
		valuePretty, _ := json.MarshalIndent(value, "", "    ")
		fmt.Printf("%+v\n", string(valuePretty))
	}
}
