package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("asumi kana")
	fmt.Println(message)
}
