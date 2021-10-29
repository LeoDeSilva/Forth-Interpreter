package main

import (
	"bufio"
	"fmt"
	"forth/lexer"
	"forth/parser"
	"io"
	"os"
)

const PROMPT = ">>"

func startRepl(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    for {
        fmt.Fprintf(out,PROMPT)
        scanned := scanner.Scan()

        if !scanned {
            return
        }

        line := scanner.Text()

        l := lexer.New(line)
        tokens := l.Lex()
        fmt.Println(tokens)

        p := parser.New(tokens)
        ast, err := p.Parse()
        if err {continue}
        fmt.Println(ast)
    }
}


func main(){
    startRepl(os.Stdin, os.Stdout)
}
