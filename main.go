package main

import (
	"github.com/pcherednichenko/invoicesSchedule/app/customers"
	"github.com/pcherednichenko/invoicesSchedule/app/invoices"
)

// in the future such parameters should be in the config, as well as the url address where the requests are sent
const filename = "customers.csv"

// main the program sends requests for purchase to users according to the schedule from the CSV file
func main() {
	lines, err := customers.LinesFromCSV(filename)
	if err != nil {
		panic(err)
	}
	c, err := customers.FromLines(lines)
	if err != nil {
		panic(err)
	}
	invoices.SendCustomers(c)
}
