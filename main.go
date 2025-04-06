package main

import(
	"fmt"
	"os"
	"os/user"
	"github.com/0x1bvaidi/go_interpreter/repl"
)

func main(){
	user, err := user.Current()
	if err != nil{
		panic(err)
	}
	fmt.Printf("Hello! %s, this is the DADDY programming langauge version 0.1\n", user.Username)
	fmt.Printf("fill free to type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
