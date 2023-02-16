package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	//"example.com/randomgreetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	//message, err := greetings.Hello("Arnab")
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)

	//message, err := randomgreetings.Hello("Arnab")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	//fmt.Println(message)
	fmt.Println(messages)
}
