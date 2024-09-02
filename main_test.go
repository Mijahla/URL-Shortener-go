package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHandleFrm tests the handleFrm function
func TestHandleFrm(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleFrm)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "<h2>URL Shortener</h2>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Handler did not return expected body: got %v want %v", rr.Body.String(), expected)
	}
}

// TestHandleShort tests the handleShort function
func TestHandleShort(t *testing.T) {
	form := strings.NewReader("url=http://example.com")
	req, err := http.NewRequest(http.MethodPost, "/shorturl", form)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleShort)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Shortened URL:"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Handler did not return expected body: got %v want %v", rr.Body.String(), expected)
	}
}

// TestHandleRdirect tests the handleRdirect function
func TestHandleRdirect(t *testing.T) {
	// First, add a mapping to the urlMap
	shortKey := "abc123"
	originalURL := "http://example.com"
	urlMap[shortKey] = originalURL

	req, err := http.NewRequest(http.MethodGet, "/short/"+shortKey, nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleRdirect)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMovedPermanently {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusMovedPermanently)
	}

	if location := rr.Header().Get("Location"); location != originalURL {
		t.Errorf("Handler returned wrong location header: got %v want %v", location, originalURL)
	}
}

// TestHandleRdirectNotFound tests the handleRdirect function when the short key is not found
func TestHandleRdirectNotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/short/unknownKey", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleRdirect)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
