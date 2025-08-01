package model

type ShortCode string

type ShortURL struct {
	ID          string    `json:"id"`
	URL         string    `json:"url"`
	ShortCode   ShortCode `json:"short_code"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	AccessCount int       `json:"access_count"`
}

type Controller interface {
	CreateShortURL(url string) (*ShortURL, error)
	GetOriginalURL(shortCode ShortCode) (*ShortURL, error)
	UpdateShortURL(shortCode ShortCode, newURL string) (*ShortURL, error)
	DeleteShortURL(shortCode ShortCode) error
}
