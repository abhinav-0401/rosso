package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

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

		programPretty, _ := json.Marshal(program)
		fmt.Printf("%+v", string(programPretty))
	}
}
