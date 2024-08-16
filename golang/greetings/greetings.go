package greetings

import "fmt"

// hello returns a greting for the named person
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. welcome!", name)
	return message
}
