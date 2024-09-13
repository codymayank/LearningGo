package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
// Here I am returning tuple of string and error
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	if name == "" {
		// errors.New is a standard library function that returns an error message
		return "", errors.New("empy name")
	}
	message := fmt.Sprintf("Hi, %s. Dude!", name)
	sendHello(name)
	// If no error, return nil
	return message, nil
}
