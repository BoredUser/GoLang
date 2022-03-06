package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Rahul")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Rahul", "Tanmay", "GoLang"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}

}
