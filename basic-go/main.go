package main

import (
	"basic-go/arrays_concept"
	"basic-go/greetings"
	"basic-go/loops"
	"fmt"
	"log"
)

func main() {
	var a = 5
	var p = &a // p holds variable a's memory address
	fmt.Printf("Address of var a: %p\n", p)
	fmt.Printf("Value of var a: %v\n", *p)
	arrays_concept.Exec()
	loops.Exec()
	message, err := greetings.Hello("Mayank")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(message)
}
