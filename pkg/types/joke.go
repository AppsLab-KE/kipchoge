package types

type JokeResponse struct {
	CreatedAt  string        `json:"created_at"`
	IconURL    string        `json:"icon_url"`
	ID         string        `json:"id"`
	UpdatedAt  string        `json:"updated_at"`
	URL        string        `json:"url"`
	Value      string        `json:"value"`
}
type JokeRequest struct {
	Value      string        `json:"value"`
	URL        string        `json:"url"`
	IconURL    string        `json:"icon_url"`
}
