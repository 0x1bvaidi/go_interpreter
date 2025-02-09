package repl

import (
	"bufio"
	// buffered io--> package providing efficient reading and writing methods for input files
	// bufio reads it as a chunk rather than byte-by-byte reducing overhead
	"fmt"
	"github.com/0x1bvaidi/go_interpreter/lexer"
	"github.com/0x1bvaidi/go_interpreter/token"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in) // NewScanner reads the input line-by-line or token-by-token from io.Reader
	// reading user stdin from CLI

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan() // Scan() method of bufio iterate over input, each call to Scan() advances the scanner to next token
		if !scanned {
			return
		}

		line := scanner.Text() //.Text() retrieves the current token as a string
		l := lexer.New(line) // line is given as input to the lexer where it acts as an input to the 'New' func and
		// returns a struct that will analyze the input
		// l stores the new lexer instance

		// first time using c style for loop
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("+%v\n", tok)
		}
	}
}
