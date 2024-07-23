package types

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Text        string `json:"text"`
	Issue       string `json:"issue"`
}
