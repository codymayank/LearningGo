package greetings

import "fmt"

func sendHello(name string) {
	message := "Hi, " + name + ". Message is sent!"
	fmt.Println(message)
}
