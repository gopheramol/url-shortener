package usecase

type ShortURL struct {
	ID        int    `json:"id"`
	FullURL   string `json:"full_url"`
	ShortCode string `json:"short_code"`
}

type encodeRequest struct {
	URL string `json:"url"`
}

type encodeResponse struct {
	ShortURL string `json:"short_url"`
}

type decodeRequest struct {
	ShortURL string `json:"short_url"`
}

type decodeResponse struct {
	URL string `json:"url"`
}
