package main

import (
	"fmt"
	"lioslang/internal/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome to lioslang Interpreter %s!\n", user.Username)
	fmt.Printf("Testing REPL\n")
	repl.Start(os.Stdin, os.Stdout)
}
