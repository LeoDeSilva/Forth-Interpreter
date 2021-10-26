package main

import (
    "forth/lexer"
    "bufio"
    "fmt"
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

        for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {
            fmt.Printf("%+v\n", tok)
        }

    }
}


func main(){
    startRepl(os.Stdin, os.Stdout)
}
