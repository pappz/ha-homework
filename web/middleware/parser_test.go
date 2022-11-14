package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	errMsgInvalidResp = `{"Message":"invalid input"}`
)

type sampleResponse struct {
	Payload string `json:"payload"`
}

type sampleRequest struct {
	Payload string `json:"payload"`
}

func (s sampleRequest) Validate() error {
	if s.Payload != "sample" {
		return errors.New("invalid input")
	}
	return nil
}

func emptyHandler() func(http.ResponseWriter, *http.Request) {
	hfn := func(ri *RequestInfo) (ResponseData, error) {
		return nil, nil
	}
	return Handle(hfn, nil)
}

func jsonAsRequestHandler() func(http.ResponseWriter, *http.Request) {
	hfn := func(ri *RequestInfo) (ResponseData, error) {
		_, ok := ri.Payload.(*sampleRequest)
		if !ok {
			return nil, ErrRespInternalError
		}
		return nil, nil
	}

	dfn := func() Json {
		return &sampleRequest{}
	}
	return Handle(hfn, dfn)
}

func stringResponseHandler(respData string) func(http.ResponseWriter, *http.Request) {
	hfn := func(ri *RequestInfo) (ResponseData, error) {
		_, _ = ri.W.Write([]byte(respData))
		return nil, nil
	}

	return Handle(hfn, nil)
}

func jsonResponseHandler(respData sampleResponse) func(http.ResponseWriter, *http.Request) {
	hfn := func(ri *RequestInfo) (ResponseData, error) {
		return respData, nil
	}

	return Handle(hfn, nil)
}

func TestJsonParser_EmptyHandler(t *testing.T) {
	resp := setupRecord(emptyHandler(), http.MethodGet, nil)
	defer resp.Body.Close()

	err := checkBody(resp.Body, "")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestJsonParser_Handle_JsonRequest(t *testing.T) {
	sr := sampleRequest{"sample"}
	jsonData, _ := json.Marshal(sr)

	resp := setupRecord(jsonAsRequestHandler(), http.MethodPost, bytes.NewBuffer(jsonData))
	defer resp.Body.Close()

	wantCode := 200
	if wantCode != resp.StatusCode {
		t.Fatalf("unexpected response code, want: %d, got: %d\n", wantCode, resp.StatusCode)
	}

	err := checkBody(resp.Body, "")
	if err != nil {
		t.Error(err.Error())
	}
}

// TestJsonParser_InvalidJsonRequest check json validation function
func TestJsonParser_InvalidJsonRequest(t *testing.T) {
	sr := sampleRequest{"invalid"}
	jsonData, _ := json.Marshal(sr)

	resp := setupRecord(jsonAsRequestHandler(), http.MethodPost, bytes.NewBuffer(jsonData))
	defer resp.Body.Close()

	wantCode := 400
	if wantCode != resp.StatusCode {
		t.Fatalf("unexpected response code, want: %d, got: %d\n", wantCode, resp.StatusCode)
	}

	err := checkBody(resp.Body, errMsgInvalidResp)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestJsonParser_StringResponse(t *testing.T) {
	sampleResponseString := "myPayload"
	resp := setupRecord(stringResponseHandler(sampleResponseString), http.MethodPost, nil)
	defer resp.Body.Close()

	wantCode := 200
	if wantCode != resp.StatusCode {
		t.Fatalf("unexpected response code, want: %d, got: %d\n", wantCode, resp.StatusCode)
	}

	err := checkBody(resp.Body, sampleResponseString)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestJsonParser_JsonResponse(t *testing.T) {
	sampleResp := sampleResponse{"myPayload"}
	resp := setupRecord(jsonResponseHandler(sampleResp), http.MethodPost, nil)
	defer resp.Body.Close()

	wantCode := 200
	if wantCode != resp.StatusCode {
		t.Fatalf("unexpected response code, want: %d, got: %d\n", wantCode, resp.StatusCode)
	}

	sampleRespByte, _ := json.Marshal(sampleResp)

	err := checkBody(resp.Body, string(sampleRespByte))
	if err != nil {
		t.Error(err.Error())
	}
}

func setupRecord(handlerFn func(http.ResponseWriter, *http.Request), method string, body io.Reader) *http.Response {
	req := httptest.NewRequest(method, "/", body)
	recorder := httptest.NewRecorder()
	handlerFn(recorder, req)
	return recorder.Result()
}

func checkBody(body io.ReadCloser, expected string) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	if string(data) != expected {
		return errors.New(fmt.Sprintf("unexpected result: '%s', got: '%s'", expected, data))
	}
	return nil
}
