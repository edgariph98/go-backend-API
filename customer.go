package main

// base class
type Customer struct {
	name    string
	age     uint
	tickets []Ticket
}

//  member function
func (customer Customer) age_in_n_years(n int) int {
	return int(customer.age) + n
}
