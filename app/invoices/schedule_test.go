package invoices

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/pcherednichenko/invoicesSchedule/app/customers"
)

// TestBasicSendInvoices check basic work of sending invoices
func TestBasicSendInvoices(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"paid":true}`))
		if r.Method != "POST" {
			t.Fail()
		}
		bodyReq, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.FailNow()
		}
		if err := r.Body.Close(); err != nil {
			t.FailNow()
		}
		req := &messageRequest{}
		if err := json.Unmarshal(bodyReq, req); err != nil {
			t.FailNow()
		}
		if req.Email != "test@email.com" {
			t.Fail()
		}
		if req.Text != "test" {
			t.Fail()
		}
	}))
	sendInvoices(testServer.URL, customers.Customer{"test@email.com", "test", []time.Duration{time.Microsecond}})
}
