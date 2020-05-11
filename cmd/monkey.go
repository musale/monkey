package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/musale/monkey/pkg/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hello %s!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
