package customers

import (
	"testing"
	"time"
)

var (
	email    = "paulcheredn@gmail.com"
	text     = "test text"
	schedule = "4s-15s"
)

// TestNormalLinesFromFile
func TestNormalFromLines(t *testing.T) {
	lines := [][]string{
		{"header email", "header text", "header schedule"},
		{email, text, schedule},
	}
	c, err := FromLines(lines)
	if err != nil {
		t.Error(err)
	}
	if c[0].Email != email {
		t.Fail()
	}
	if c[0].Text != text {
		t.Fail()
	}
	if c[0].Schedule[0].String() != "4s" {
		t.Fail()
	}
}

// TestEmptyEmail email field is empty
func TestEmptyEmail(t *testing.T) {
	lines := [][]string{
		{"header email", "header text", "header schedule"},
		{"", text, schedule},
	}
	_, err := FromLines(lines)
	if err != errEmptyEmail {
		t.Error(err)
	}
}

// TestParseDuration test that duration parsed right
func TestParseDuration(t *testing.T) {
	durations, err := durationFromWord("2s-5s-10h")
	if err != nil {
		t.Error(err)
	}
	if durations[0] != time.Second*2 {
		t.Fail()
	}
	if durations[1] != time.Second*5 {
		t.Fail()
	}
	if durations[2] != time.Hour*10 {
		t.Fail()
	}
}

// TestWrongDuration test that wrong durations not will be parsed
func TestWrongDuration(t *testing.T) {
	durations, err := durationFromWord("test")
	if err == nil {
		t.Error(err)
	}
	if len(durations) > 0 {
		t.Fail()
	}
}
