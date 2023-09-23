package main

import (
	"fmt"
	"os"

	"github.com/cian911/monkey-language/repl"
)

func main() {
	fmt.Println("Welcome to my monkey progreamming language.")
	repl.Start(os.Stdin, os.Stdout)
}
