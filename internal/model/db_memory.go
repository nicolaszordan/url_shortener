package model

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/fnv"
	"time"
)

var (
	ErrNotFound = errors.New("short URL not found")
)

type MemoryDB struct {
	urls map[ShortCode]*ShortURL
	id   int
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		urls: make(map[ShortCode]*ShortURL),
		id:   0,
	}
}

func (db *MemoryDB) CreateShortURL(url string) (*ShortURL, error) {
	shortCode := ShortCode(generateShortCode(url))
	shortURL := &ShortURL{
		ID:        db.generateID(),
		URL:       url,
		ShortCode: shortCode,
		CreatedAt: getCurrentTime(),
		UpdatedAt: getCurrentTime(),
	}

	db.urls[shortCode] = shortURL
	return shortURL, nil
}

func (db *MemoryDB) GetOriginalURL(shortCode ShortCode) (*ShortURL, error) {
	shortURL, exists := db.urls[shortCode]
	if !exists {
		return nil, ErrNotFound
	}
	return shortURL, nil
}

func (db *MemoryDB) UpdateShortURL(shortCode ShortCode, newURL string) (*ShortURL, error) {
	shortURL, exists := db.urls[shortCode]
	if !exists {
		return nil, ErrNotFound
	}

	shortURL.URL = newURL
	shortURL.UpdatedAt = getCurrentTime()
	return shortURL, nil
}

func (db *MemoryDB) DeleteShortURL(shortCode ShortCode) error {
	_, exists := db.urls[shortCode]
	if !exists {
		return ErrNotFound
	}

	delete(db.urls, shortCode)
	return nil
}

func (db *MemoryDB) generateID() string {
	db.id++
	return fmt.Sprintf("%d", db.id)
}

func generateShortCode(url string) ShortCode {
	h := fnv.New64a()
	h.Write([]byte(url))
	hash := h.Sum64()

	byteHash := make([]byte, 8)
	binary.BigEndian.PutUint64(byteHash, hash)

	base64URL := base64.URLEncoding.EncodeToString(byteHash)
	return ShortCode(base64URL[:8]) // Use only first 8 characters for short code
}

func getCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}
