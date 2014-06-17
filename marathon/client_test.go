package marathon

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Run Get-request against an url and return body
func TestGetBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ExampleTest")
	}))
	defer ts.Close()

	b, err := GetBody(ts.URL)
	if err != nil {
		t.Fatal("Could not get body of test")
	}

	if string(b) == "ExampleTest" {
		t.Fatal("Didn't get ExampleTest as output")
	}
}
