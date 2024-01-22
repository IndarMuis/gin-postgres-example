package entity

type Book struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	AuthorId      uint   `json:"author_id"`
	Category      string `json:"category"`
	PublishedYear string `json:"published_year"`
}
