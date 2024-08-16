package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello returns a greting for the named person
func Hello(name string) (string, error) {
	// nameが空文字ならエラーを返す
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
