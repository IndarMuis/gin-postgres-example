package dto

import "github.com/IndarMuis/gin-postgres-example.git/src/model/entity"

type BookRequest struct {
	Title         string `json:"title" binding:"required"`
	AuthorId      int    `json:"author_id" binding:"required"`
	Category      string `json:"category" binding:"required"`
	PublishedYear string `json:"published_year" binding:"required"`
}

type BookUpdateRequest struct {
	Title         string `json:"title" binding:"required"`
	Category      string `json:"category" binding:"required"`
	PublishedYear string `json:"published_year" binding:"required"`
}

type BookResponse struct {
	ID            uint   `json:"id" binding:"required"`
	Title         string `json:"title" binding:"required"`
	AuthorId      uint   `json:"author_id" binding:"required"`
	Category      string `json:"category,omitempty" binding:"required"`
	PublishedYear string `json:"published_year,omitempty" binding:"required"`
}

type BookDetailsResponse struct {
	ID            uint           `json:"id" binding:"required"`
	Title         string         `json:"title" binding:"required"`
	Author        *entity.Author `json:"author" binding:"required"`
	Category      string         `json:"category,omitempty" binding:"required"`
	PublishedYear string         `json:"published_year,omitempty" binding:"required"`
}
