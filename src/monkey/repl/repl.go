package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/src/monkey/lexer"
	"monkey/src/monkey/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NexToken(); tok.Type != token.EOF; tok = l.NexToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
