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
	e.DeclareVar("x", &object.NumLitObject{Kind: object.INT, Value: 4}, false)
	e.AssignVar("x", &object.NumLitObject{Kind: object.INT, Value: 10})

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
