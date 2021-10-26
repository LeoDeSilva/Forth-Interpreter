package main

import (
	"bufio"
	"fmt"
	"forth/lexer"
	"io"
	"os"
)

const PROMPT = ">>"

func startRepl(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    for {
        fmt.Printf(PROMPT)
        scanned := scanner.Scan()

        if !scanned {
            return
        }

        line := scanner.Text()

        l := lexer.New(line)
        tokens := l.Lex()
        fmt.Println(tokens)
    }
}


func main(){
    startRepl(os.Stdin, os.Stdout)
}
