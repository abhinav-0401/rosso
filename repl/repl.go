package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/abhinav-0401/rosso/eval"
	"github.com/abhinav-0401/rosso/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		parser := parser.New()
		program := parser.ProduceAst(line)
		value := eval.Eval(program)

		// programPretty, _ := json.MarshalIndent(program, "", "    ")
		valuePretty, _ := json.MarshalIndent(value, "", "    ")

		// fmt.Printf("%+v\n", string(programPretty))
		fmt.Printf("%+v\n", string(valuePretty))
	}
}
