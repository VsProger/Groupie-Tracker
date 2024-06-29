package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func mockFetchArtists() ([]Artist, error) {
	return []Artist{
		{ID: 1, Name: "Artist 1"},
		{ID: 2, Name: "Artist 2"},
	}, nil
}

func initTestTemplates() {
	var err error
	basePath, _ := os.Getwd()
	tpl, err = template.ParseGlob(filepath.Join(basePath, "../templates", "*"))
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	initTestTemplates()
	os.Exit(m.Run())
}

func TestHome(t *testing.T) {
	originalFetchArtists := FetchArtists
	FetchArtists = mockFetchArtists
	defer func() { FetchArtists = originalFetchArtists }()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Home)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `<title>Groupie-Tracker</title>`
	if !contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestArtistHandler(t *testing.T) {
	originalFetchArtists := FetchArtists
	FetchArtists = mockFetchArtists
	defer func() { FetchArtists = originalFetchArtists }()

	req, err := http.NewRequest("GET", "/artist/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ArtistHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `<title>Artist 1 - Artist Details</title>`
	if !contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestInvalidArtistID(t *testing.T) {
	req, err := http.NewRequest("GET", "/artist/abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ArtistHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestArtistNotFound(t *testing.T) {
	originalFetchArtists := FetchArtists
	FetchArtists = mockFetchArtists
	defer func() { FetchArtists = originalFetchArtists }()

	req, err := http.NewRequest("GET", "/artist/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ArtistHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestNonGETMethod(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Home)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestArtistIDWithSpecialChars(t *testing.T) {
	req, err := http.NewRequest("GET", "/artist/1@#", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ArtistHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestArtistIDTooLong(t *testing.T) {
	req, err := http.NewRequest("GET", "/artist/12345678901", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ArtistHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestArtistIDFractional(t *testing.T) {
	req, err := http.NewRequest("GET", "/artist/1.5", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ArtistHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func contains(body, substr string) bool {
	return strings.Contains(body, substr)
}
