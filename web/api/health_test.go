package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pappz/ha-homework/service"
)

func TestHealth(t *testing.T) {
	serviceService := service.NewSector(56)
	testServer := httptest.NewServer(Router(serviceService))
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
