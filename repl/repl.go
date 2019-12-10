// Package repl has the program that will take input
// send it to the interpreter to be evaluated and print
// the result of the evaluation. Just like REPL in Python,
// JavaScript runtimes, etc.
package repl

import (
	"bufio"
	"io"

	"fmt"
	"github.com/Lumexralph/chimp/lexer"
	"github.com/Lumexralph/chimp/token"
)

const PROMPT = ">>"

// Start - read from the input source until encountering a newline,
// take the just read line and pass it to an instance of our lexer
// and finally print all the tokens the lexer gives us until we encounter EOF.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Println(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
