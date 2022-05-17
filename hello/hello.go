package main

import (
	"fmt"
	"log"

	"exemple.org/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	var name string

	fmt.Printf("Give your name: ")
	fmt.Scan(&name)

	message, err := greetings.Hello(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
