package invoices

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pcherednichenko/invoicesSchedule/app/customers"
	"github.com/pcherednichenko/invoicesSchedule/app/logger"
)

const baseURL = "http://localhost:9090/messages"

// SendCustomers send information to customers
func SendCustomers(c customers.Customers) {
	var wg sync.WaitGroup
	for _, customer := range c {
		wg.Add(1)
		go func(c customers.Customer) {
			defer wg.Done()
			sendInvoices(baseURL, c)
		}(customer)
	}
	wg.Wait()
	logger.LogInfo("All invoices for customers have been sent!")
}

// sendInvoices to the customer in accordance with the schedule
func sendInvoices(url string, customer customers.Customer) {
	var (
		done     = make(chan bool)
		n        = len(customer.Schedule)
		lastStep bool
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i, delay := range customer.Schedule {
		timer := time.NewTimer(delay)
		lastStep = i == n-1
		go func(ctx context.Context, lastStep bool) {
			<-timer.C
			select {
			case <-ctx.Done():
				return
			default:
				paid, err := sendMessage(url, customer.Email, customer.Text)
				if err != nil {
					logger.LogError(err)
					return
				}
				// if paid, then send a signal done, or if this is the last step in the schedule
				if paid || lastStep {
					done <- true
					return
				}
			}
		}(ctx, lastStep)
	}

	<-done
	logger.LogInfo(fmt.Sprintf("All invoices for customer: '%s' have been sent", customer.Email))
}
