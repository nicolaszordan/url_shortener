package model

import (
	"errors"
	"testing"
)

func TestMemoryDB(t *testing.T) {
	db := NewMemoryDB()

	// Test CreateShortURL
	shortURL, err := db.CreateShortURL("https://www.example.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if shortURL.URL != "https://www.example.com" {
		t.Errorf("expected URL to be 'https://www.example.com', got '%s'", shortURL.URL)
	}

	// Test GetOriginalURL
	retrievedURL, err := db.GetOriginalURL(shortURL.ShortCode)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if retrievedURL.URL != shortURL.URL {
		t.Errorf("expected retrieved URL to match created URL, got '%s'", retrievedURL.URL)
	}

	// Test UpdateShortURL
	newShortURL, err := db.UpdateShortURL(shortURL.ShortCode, "https://www.updated-example.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if newShortURL.URL != "https://www.updated-example.com" {
		t.Errorf("expected updated URL to be 'https://www.updated-example.com', got '%s'", newShortURL.URL)
	}

	// Test DeleteShortURL
	err = db.DeleteShortURL(shortURL.ShortCode)
	if err != nil {
		t.Fatalf("expected no error on delete, got %v", err)
	}

	// Test GetOriginalURL after deletion
	_, err = db.GetOriginalURL(shortURL.ShortCode)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound after deletion, got %v", err)
	}
}
