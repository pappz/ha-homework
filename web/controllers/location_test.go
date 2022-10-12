package controllers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pappz/ha-homework/service"
)

func TestLocation(t *testing.T) {
	serviceService := service.NewSector(1)
	testServer := httptest.NewServer(Router(serviceService))
	defer testServer.Close()

	requestData := []byte(`{
				"x": "123.12",
				"y": "456.56",
				"z": "789.89",
				"vel": "20.0"
			}`)
	expectedResult := `{"loc":"1389.57"}`

	resp, err := doRequest(http.MethodPost, testServer.URL+"/sector", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatalf("client request error: %s\n", err.Error())
	}
	if 200 != resp.StatusCode {
		t.Fatalf("received error response response: %d\n", resp.StatusCode)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read body: %v", err)
	}
	if string(data) != expectedResult {
		t.Errorf("unexpected result: %s", data)
	}
}

func doRequest(method string, url string, data *bytes.Buffer) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return client.Do(req)
}
