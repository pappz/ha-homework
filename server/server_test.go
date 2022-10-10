package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_Health(t *testing.T) {
	s := NewServer()
	testServer := httptest.NewServer(s.router())
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
	if healthResponse != string(actual) {
		t.Errorf("invalid response msg: '%s'\n", string(actual))
	}

}
