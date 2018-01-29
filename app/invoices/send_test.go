package invoices

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestServerResponseOk check basic work with ok response
func TestServerResponseOk(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	result, err := sendMessage(testServer.URL, "test@email.com", "test")
	if err != nil {
		t.Error(err)
	}
	if result != false {
		t.Fail()
	}
}

// TestServerPaid check response with paid true
func TestServerPaid(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"paid":true}`))
	}))
	result, err := sendMessage(testServer.URL, "test@email.com", "test")
	if err != nil {
		t.Error(err)
	}
	if result != true {
		t.Fail()
	}
}

// TestServerReturnBadCode check that with 404 we get error
func TestServerReturnBadCode(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{}`))
	}))
	_, err := sendMessage(testServer.URL, "test@email.com", "test")
	if err == nil {
		t.Fail()
	}
}
