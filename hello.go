package main

import (
	"fmt"

	"milzam.com/greetings"
)

func main() {
	message := greetings.Hello("Milzam")
	fmt.Println(message)
}
