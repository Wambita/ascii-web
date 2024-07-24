package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	handler "server/handlers"
)

// TestHome tests the Home function
func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler.HomeHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Home() returned status code %d, expected %d", w.Code, http.StatusOK)
	}
}

func TestAboutHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler.AboutHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Home() returned status code %d, expected %d", w.Code, http.StatusOK)
	}
}

func TestInstructionHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler.InstructionsHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Home() returned status code %d, expected %d", w.Code, http.StatusOK)
	}
}
