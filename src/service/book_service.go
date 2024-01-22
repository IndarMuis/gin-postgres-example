package service

import "github.com/IndarMuis/gin-postgres-example.git/src/model/dto"

type BookService interface {
	FindAll() ([]dto.BookResponse, error)
	FindById(id uint) (dto.BookResponse, error)
	FindBookDetails(id uint) (dto.BookDetailsResponse, error)
	Save(book dto.BookRequest) (dto.BookResponse, error)
	Update(book dto.BookRequest) (dto.BookResponse, error)
	Delete(id uint) (dto.BookResponse, error)
}
