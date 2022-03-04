package main

import (
	. "booking-app/customer"
	"fmt"
	"strconv"
)

// Functions and arrays
func generateTickets(n int) []Ticket {
	var tickets []Ticket
	var aliasTicket = "ticket"
	for i := 0; i < n; i++ {
		tickets = append(tickets, Ticket{
			Id:    i,
			Alias: aliasTicket + strconv.Itoa(i),
		})
	}
	return tickets
}

// iterating over each ticket item and print the alias
func printTickets(tickets []Ticket) {
	for _, ticket := range tickets {
		fmt.Printf("Ticket alias: %v\n", ticket.Alias)
	}
}
func main() {
	var conference_name = "Go conference"
	const conference_tickets = 50
	var remaining_tickets = conference_tickets
	// Format Printing
	fmt.Printf("Welcome to %v! \n", conference_name)
	fmt.Printf("We have %v remaining tickets \n", remaining_tickets)
	// instantiating Class
	var customer Customer
	// user input
	fmt.Print("Enter your name: ")
	fmt.Scanln(&customer.Name)
	fmt.Print("Enter Your age: ")
	fmt.Scan(&customer.Age)
	fmt.Printf("Hello %v with age %v, Welcome! \n", customer.Name, customer.Age)
	fmt.Print("How many years in adavance you want to check you will be: ")
	var n int
	fmt.Scan(&n)
	fmt.Printf("You will be %v in %v years!\n", customer.Age_in_n_years(n), n)
	fmt.Print("How many random tickets do you want to generate: ")
	fmt.Scan(&n)
	customer.Tickets = generateTickets(n)
	fmt.Printf("your tickets are of type %T\n", customer.Tickets)
	if n <= 2 {
		var total_tickets = n
		fmt.Printf("you need more than 2 tickets, you currently generated %v, how many extra do you want: ", n)
		fmt.Scan(&n)
		customer.Tickets = generateTickets(total_tickets + n)
	} else if n == 3 {
		fmt.Printf("You cannot generate exactly 3 tickets, enter new number of tickets: ")
		fmt.Scan(&n)
		customer.Tickets = generateTickets(n)
	}
	fmt.Printf("We have generated %v tickets \n", len(customer.Tickets))
	printTickets(customer.Tickets)

}
