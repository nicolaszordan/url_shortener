package model

type ShortURL struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	ShortCode string `json:"short_code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
