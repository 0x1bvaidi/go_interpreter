package repl

// read, eval, print, loop

import (
	"bufio"
	"fmt"
	"github.com/0x1bvaidi/go_interpreter/lexer"
	"github.com/0x1bvaidi/go_interpreter/token"
	"io"
)

const PROMPT = ">> "

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

		for tok := l.ConsumeToken_Advance(); tok.Type != token.EOF; tok = l.ConsumeToken_Advance() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
