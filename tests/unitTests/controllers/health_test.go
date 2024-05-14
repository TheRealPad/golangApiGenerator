package controllers

import (
	"httpServer/src/controller/health"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()
	health.Health(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "This is the health endpoint\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestShowHtml(t *testing.T) {
	_, err := http.NewRequest("GET", "/example", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	health.ShowHtml(rr, "../../../src/controller/health/health.html", nil)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
