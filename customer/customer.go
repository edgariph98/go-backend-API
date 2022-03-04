package customer

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
