package api

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	router := mux.NewRouter()
	RegisterHealthHandler(router)
	testServer := httptest.NewServer(router)
	defer testServer.Close()

	resp, err := http.Get(testServer.URL + "/health")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("received non-200 response: %d\n", resp.StatusCode)
	}
	actual, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if HealthResponse != string(actual) {
		t.Errorf("invalid response msg: '%s'\n", string(actual))
	}
}
