package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/0x1bvaidi/go_interpreter/repl"
)

func main(){
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! this is the minilang!\n", user.Username)
	fmt.Printf("feel free!!\n")
	repl.Start(os.Stdin, os.Stdout)
}
