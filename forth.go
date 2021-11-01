package main

import (
	"bufio"
	"fmt"
	"forth/evaluator"
	"forth/lexer"
	"forth/parser"
	"io"
	"os"
)


const PROMPT = ">>"

func startRepl(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    environment := &evaluator.Environment{Variables: make(map[string]interface{})}

    for {
        fmt.Fprintf(out,PROMPT)
        scanned := scanner.Scan()

        if !scanned {
            return
        }

        line := scanner.Text()

        l := lexer.New(line)
        tokens := l.Lex()

        p := parser.New(tokens)
        ast, err := p.Parse()
        if err {continue}

        e := evaluator.New(ast,environment)
        err = e.Evaluate()
        if err {continue}
    }
}


func main(){
    startRepl(os.Stdin, os.Stdout)
}
