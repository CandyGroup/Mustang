package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/makmav/Mustang/lexer"
	"github.com/makmav/Mustang/token"
)

const prompt = ">> "


func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(prompt)
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
