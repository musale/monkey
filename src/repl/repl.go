package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/musale/monkey/src/lexer"
	"github.com/musale/monkey/src/token"
)

// PROMPT will appear at the start of a statement in the REPL
const PROMPT = ">> "

// Start will fire up a repl
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if line == "exit" {
			fmt.Fprintf(out, "Exiting the REPL...")
			return
		}
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
