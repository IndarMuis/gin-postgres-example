package dto

type AuthorRequest struct {
	Name string `json:"name" binding:"required"`
}

type AuthorResponse struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
