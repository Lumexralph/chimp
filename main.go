// Package main - gets the user of the REPL and start it.
package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/Lumexralph/chimp/repl"
)

func main() {
	// get the currently logged in user
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	t := time.Now()

	// Python 3.6.5 (v3.6.5:f59c0932b4, Mar 28 2018, 05:52:31)
	fmt.Printf("Chimp v1.0.0 (%s) \n", t.Format(time.RFC1123))
	fmt.Printf("Hi %s, welcome to chimp programming language! - Interactive Mode\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
