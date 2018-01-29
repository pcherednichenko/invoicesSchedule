package customers

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

const separatorTimeSymbol = "-"

var (
	errEmptyEmail    = fmt.Errorf("empty email")
	errEmptyText     = fmt.Errorf("empty text")
	errEmptySchedule = fmt.Errorf("empty schedule")
)

// Customers notice that they are used type alias, which appeared in the version 1.9
type Customers = []Customer

// Customer with params email, text, schedule
type Customer struct {
	Email    string
	Text     string
	Schedule []time.Duration
}

// LinesFromCSV open csv and return lines from file
func LinesFromCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return csv.NewReader(f).ReadAll()
}

// FromLines loads information about customers and returns slice of customers
func FromLines(lines [][]string) (Customers, error) {
	var err error
	// to doesn't make a relocation of memory with every append()
	// we will set the number of elements in advance and will set them by index
	customers := make(Customers, len(lines)-1)
	t := 0
	for i, line := range lines {
		if i == 0 {
			// skip header
			continue
		}
		// in perfect system we should check email and etc, but here just simple check
		if len(line[0]) == 0 {
			return nil, errEmptyEmail
		}
		if len(line[1]) == 0 {
			return nil, errEmptyText
		}
		if len(line[2]) == 0 {
			return nil, errEmptySchedule
		}
		customers[t].Email = line[0]
		customers[t].Text = line[1]
		customers[t].Schedule, err = durationFromWord(line[2])
		if err != nil {
			return nil, err
		}
		t++
	}
	return customers, nil
}

// durationFromWord parse string line and return slice of durations
func durationFromWord(schedule string) ([]time.Duration, error) {
	var err error
	times := strings.Split(schedule, separatorTimeSymbol)
	durations := make([]time.Duration, len(times))
	for i, t := range times {
		durations[i], err = time.ParseDuration(t)
		if err != nil {
			return nil, err
		}
	}
	return durations, nil
}
