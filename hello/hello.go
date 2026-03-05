package main

import (
	"fmt"

	"example.com/greetings"
)

func showHello(name string) {
	message, err := greetings.HelloWithErrpr(name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message)
	}

}

func main() {
	// Get a greeting message and print it.
	showHello("Gladys")
	showHello("")

}
