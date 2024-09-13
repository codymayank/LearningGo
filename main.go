package main

import (
	"fmt"
	"log"
	"pdd-projects/arrays_concept"
	"pdd-projects/greetings"
	"pdd-projects/loops"
)

func main() {
	arrays_concept.Exec()
	loops.Exec()
	message, err := greetings.Hello("Mayank")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(message)
}
