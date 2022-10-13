package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/pappz/ha-homework/service"
)

var (
	errMsgInvalidCoordinate = fmt.Sprintf(`{"Message":"%s"}`, errInvalidCoordinate.Error())
	errMsgInvalidVelocity   = fmt.Sprintf(`{"Message":"%s"}`, errInvalidVelocity.Error())
	testServer              *httptest.Server
)

func TestMain(m *testing.M) {
	serviceService := service.NewSector(1)
	testServer = httptest.NewServer(Router(serviceService))

	code := m.Run()

	testServer.Close()
	os.Exit(code)
}

func TestLocation_sampleData(t *testing.T) {
	requestData := []byte(`{
				"x": "123.12",
				"y": "456.56",
				"z": "789.89",
				"vel": "20.0"
			}`)
	expectedResult := `{"loc":1389.57}`

	resp, err := doRequest(http.MethodPost, testServer.URL+"/databank", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatalf("client request error: %s\n", err.Error())
	}
	if 200 != resp.StatusCode {
		t.Fatalf("received error response response: %d\n", resp.StatusCode)
	}

	err = checkBody(resp.Body, expectedResult)
	if err != nil {
		t.Error(err.Error())
	}
	_ = resp.Body.Close()
}

func TestLocation_missingCoords(t *testing.T) {
	cases := [][]byte{
		[]byte(`{
				"y": "456.56",
				"z": "789.89",
				"vel": "20.0"
			}`),
		[]byte(`{
				"x": "123.12",
				"z": "789.89",
				"vel": "20.0"
			}`),
		[]byte(`{
				"x": "123.12",
				"y": "456.56",
				"vel": "20.0"
			}`),
	}

	for i := 0; i < len(cases); i++ {
		resp, err := doRequest(http.MethodPost, testServer.URL+"/databank", bytes.NewBuffer(cases[i]))
		if err != nil {
			t.Fatalf("client request error: %s\n", err.Error())
		}
		if 400 != resp.StatusCode {
			t.Fatalf("received incorrect response code: %d\n", resp.StatusCode)
		}

		err = checkBody(resp.Body, errMsgInvalidCoordinate)
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		_ = resp.Body.Close()
	}
}

func TestLocation_missingVel(t *testing.T) {
	requestData := []byte(`{
				"x": "123.12",
				"y": "456.56",
				"z": "789.89"
			}`)

	resp, err := doRequest(http.MethodPost, testServer.URL+"/databank", bytes.NewBuffer(requestData))
	if err != nil {
		t.Fatalf("client request error: %s\n", err.Error())
	}
	if 400 != resp.StatusCode {
		t.Fatalf("received incorrect response code: %d\n", resp.StatusCode)
	}

	err = checkBody(resp.Body, errMsgInvalidVelocity)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	_ = resp.Body.Close()
}

func TestLocation_wrongTypes(t *testing.T) {
	serviceService := service.NewSector(1)
	testServer := httptest.NewServer(Router(serviceService))
	defer testServer.Close()

	cases := [][]byte{
		[]byte(`{
				"x": abc,
				"y": "456.56",
				"z": "789.89",
				"vel": "20.0"
			}`),
		[]byte(`{
				"x": "123.12",
				"y": abc,
				"z": "789.89",
				"vel": "20.0"
			}`),
		[]byte(`{
				"x": "123.12",
				"y": "456.56",
				"z": abc,
				"vel": "20.0"
			}`),
		[]byte(`{
				"x": "123.12",
				"y": "456.56",
				"z": "789.89",
				"vel": abc
			}`),
	}

	for i := 0; i < len(cases); i++ {
		resp, err := doRequest(http.MethodPost, testServer.URL+"/databank", bytes.NewBuffer(cases[i]))
		if err != nil {
			t.Fatalf("client request error: %s\n", err.Error())
		}
		if 400 != resp.StatusCode {
			t.Fatalf("received incorrect response code: %d\n", resp.StatusCode)
		}
		_ = resp.Body.Close()
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

func checkBody(body io.ReadCloser, expected string) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	if string(data) != expected {
		return errors.New(fmt.Sprintf("unexpected result: %s", data))
	}
	return nil
}
