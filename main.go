package main

import "fmt"

func main() {
	var conference_name = "Go conference"
	const conference_tickets = 50
	var remaining_tickets = conference_tickets
	// Format Printing
	fmt.Printf("Welcome to %v! \n", conference_name)
	fmt.Printf("We have %v remaining tickets \n", remaining_tickets)

	var name string
	var age int
	// User Input

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter Your age: ")
	fmt.Scan(&age)
	fmt.Printf("Hello %v with age %v, Welcome! \n", name, age)
}
