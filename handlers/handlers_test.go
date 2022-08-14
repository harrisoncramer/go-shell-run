package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatusHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestJobHandler(t *testing.T) {
	type JobType struct {
		Jobs []string `json:"jobs"`
	}

	data, err := json.Marshal(JobType{
		Jobs: []string{"ls -la", "pwd"},
	})

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/jobs", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(JobHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `{"running": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestJobHandlerNoJobs(t *testing.T) {
	type NoJobs struct{}

	data, err := json.Marshal(NoJobs{})

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/jobs", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(JobHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Result().StatusCode; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	expectedBody := "Must provide jobs."
	body, err := io.ReadAll(rr.Body)
	if strings.TrimSpace(string(body)) != strings.TrimSpace(expectedBody) {
		t.Errorf("handler returned wrong body: got %s want %s", body, expectedBody)
	}
}
