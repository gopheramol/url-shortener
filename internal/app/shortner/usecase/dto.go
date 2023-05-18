package usecase

type ShortURL struct {
	URL      string `json:"url,omitempty"`
	ShortURL string `json:"short_url,omitempty"`
}
