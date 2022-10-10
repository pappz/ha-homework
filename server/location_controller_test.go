package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLocation(t *testing.T) {
	s := NewServer()
	testServer := httptest.NewServer(s.router())
	defer testServer.Close()

	requestData := []byte(`{
			"x": "123.12",
			"y": "456.56",
			"z": "789.89"",
			"vel": "20.0"
		}`)

	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, testServer.URL+"/sector", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("client request error: %s\n", err.Error())
	}
	if 200 != resp.StatusCode {
		t.Fatalf("received error response response: %d\n", resp.StatusCode)
	}

}
