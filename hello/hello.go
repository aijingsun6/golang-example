package main

import (
	"fmt"
	"reflect"

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

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays // this constant is not exported
)

type (
	A1 = string
	A2 = A1
)

type (
	B1 string
	B2 B1
	B3 []B1
	B4 B3
)

func underlyingTypes() {
	var b B4
	fmt.Printf("%v %T", b, b)
	var b1 B1
	var b2 B2
	fmt.Println(reflect.TypeOf(b1))
	fmt.Println(reflect.TypeOf(b2))
	fmt.Println(reflect.TypeOf(b1) == reflect.TypeOf(b2))

}

func main() {
	// Get a greeting message and print it.
	showHello("Gladys")
	showHello("")
	fmt.Println(Sunday)
	underlyingTypes()

}
