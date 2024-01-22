package impl

import (
	"github.com/IndarMuis/gin-postgres-example.git/src/model/dto"
	"github.com/IndarMuis/gin-postgres-example.git/src/model/entity"
	"github.com/IndarMuis/gin-postgres-example.git/src/repository"
	"github.com/IndarMuis/gin-postgres-example.git/src/service"
)

type BookServiceImpl struct {
	bookRepository repository.BookRepository
}

func (bookService *BookServiceImpl) FindBookDetails(id uint) (dto.BookDetailsResponse, error) {
	book, err := bookService.bookRepository.FindBookDetails(id)
	if err != nil {
		return dto.BookDetailsResponse{}, err
	}

	return *book, nil
}

func (bookService *BookServiceImpl) FindAll() ([]dto.BookResponse, error) {
	resultBooks, err := bookService.bookRepository.FindAll()
	if err != nil {
		panic(err)
	}

	var books []dto.BookResponse
	for _, book := range resultBooks {
		books = append(books, dto.BookResponse{
			ID:            book.ID,
			Title:         book.Title,
			AuthorId:      book.AuthorId,
			Category:      book.Category,
			PublishedYear: book.PublishedYear,
		})
	}

	return books, nil
}

func (bookService *BookServiceImpl) FindById(id uint) (dto.BookResponse, error) {
	bookResult, err := bookService.bookRepository.FindById(id)

	if err != nil {
		return dto.BookResponse{}, err
	}

	book := dto.BookResponse{
		ID:            bookResult.ID,
		Title:         bookResult.Title,
		AuthorId:      bookResult.AuthorId,
		Category:      bookResult.Category,
		PublishedYear: bookResult.PublishedYear,
	}

	return book, nil
}

func (bookService *BookServiceImpl) Save(book dto.BookRequest) (dto.BookResponse, error) {
	newBook := entity.Book{
		Title:         book.Title,
		Category:      book.Category,
		PublishedYear: book.PublishedYear,
	}

	result, err := bookService.bookRepository.Save(&newBook, book.AuthorId)
	if err != nil {
		panic(err)
	}

	return dto.BookResponse{
		ID:            result.ID,
		Title:         result.Title,
		AuthorId:      result.AuthorId,
		Category:      result.Category,
		PublishedYear: result.PublishedYear,
	}, nil
}

func (bookService *BookServiceImpl) Update(book dto.BookRequest) (dto.BookResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (bookService *BookServiceImpl) Delete(id uint) (dto.BookResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewBookService(bookRepository repository.BookRepository) service.BookService {
	return &BookServiceImpl{bookRepository: bookRepository}
}
