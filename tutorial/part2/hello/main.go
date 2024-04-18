package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	// Get a greeting message and print it.
	// Deixe a string vazia para retornar um erro
	message, err := greetings.Hello("marcel")

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	idade := greetings.Idade(17)

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message, idade)
}
