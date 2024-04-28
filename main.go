package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xsni1/go-monkey-interpreter/lexer"
	"github.com/xsni1/go-monkey-interpreter/token"
)

func main() {
    s := bufio.NewScanner(os.Stdin)
    for {
        fmt.Printf(">>> ")
        if !s.Scan() {
            return
        }

        l := lexer.New(s.Text())
        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
            fmt.Println(tok)
        }
    }
}
