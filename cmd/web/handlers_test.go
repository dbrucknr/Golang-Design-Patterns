package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllDogBreedsJson(t *testing.T) {
	// Create a Request
	req, _ := http.NewRequest("GET", "/api/dog-breeds", nil)
	// Create a Response Recorder
	rr := httptest.NewRecorder()
	// Create a Handler
	handler := http.HandlerFunc(testApp.GetAllDogBreedsJson)
	// Serve the Handler
	handler.ServeHTTP(rr, req)
	// Check the Response
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}
