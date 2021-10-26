package main

import (
    "os"
    "forth/repl"
)

func main(){
    repl.Start(os.Stdin, os.Stdout)
}
