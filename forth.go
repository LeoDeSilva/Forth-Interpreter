package main

import (
	"bufio"
	"fmt"
	"forth/evaluator"
	"forth/lexer"
	"forth/parser"
	"io"
	"os"
	"strings"
)


func startRepl(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    environment := &evaluator.Environment{Functions:make(map[string]parser.ProgramNode), Variables: make(map[int]int), Stack: make([]interface{},0), Identifiers: make(map[string]int)}

    for {
        fmt.Fprintf(out,">>")
        scanned := scanner.Scan()

        if !scanned {
            return
        }

        line := scanner.Text()

        l := lexer.New(strings.TrimSpace(line))
        tokens := l.Lex()

        p := parser.New(tokens)
        ast, err := p.Parse()
        if err {continue}

        e := evaluator.New(ast,environment)
        err = e.Evaluate()
        if err {continue}

        fmt.Println(environment.Stack)
    }
}


func main(){
    startRepl(os.Stdin, os.Stdout)
}
