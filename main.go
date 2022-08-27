package main

import (
	"fmt"
	"github.com/rustatian/InterpreterInGo/repl"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is MPL!!!\n", usr.Username)
	repl.Start(os.Stdin, os.Stdout)
}