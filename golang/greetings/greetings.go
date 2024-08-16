package greetings

import (
	"errors"
	"fmt"
)

// hello returns a greting for the named person
func Hello(name string) (string, error) {
	// nameが空文字ならエラーを返す
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. welcome!", name)
	return message, nil
}
