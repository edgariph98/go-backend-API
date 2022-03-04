package customer

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	member functions and data members accessed on another package need to be capitalized at the beginnig to be used and exported outside
*/
type Ticket struct {
	Id    int
	Alias string
}

// base class
type Customer struct {
	Name    string
	Age     uint
	Tickets []Ticket
	// map of alias ticket to cost int
	TicketMapCost map[string]int
}

//  member function
func (customer Customer) Age_in_n_years(n int) int {
	return int(customer.Age) + n
}

func (customer Customer) Send_All_tickets(waitGroup *sync.WaitGroup) {
	for _, ticket := range customer.Tickets {
		// kicking off go routine
		waitGroup.Add(1)
		go sendTicket(ticket, waitGroup)
	}
	fmt.Println("Tickets sent kickoff complete")
}

func sendTicket(ticket Ticket, wg *sync.WaitGroup) {

	var seconds = rand.Intn(20)
	time.Sleep(time.Duration(seconds) * time.Second)
	var logStatement = fmt.Sprintf("Ticket with Alias: {%v}, sent", ticket.Alias)
	fmt.Println("\n##############################")
	fmt.Print(logStatement)
	fmt.Println("\n##############################")
	// completing go routine
	wg.Done()
}
