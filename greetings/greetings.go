package greetings

import (
	"errors"
	"fmt"

	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// a slice(array) of greetings
	formats := []string{
		"Hi %v. Welcome!",
		"Oi %v. Prazher em conhecer você",
		"Hola %v. Mucho gusto y bienvenido",
	}

	return formats[rand.Intn(len(formats))]
}
