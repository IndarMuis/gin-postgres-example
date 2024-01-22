package repository

import (
	"database/sql"
	"github.com/IndarMuis/gin-postgres-example.git/src/model/dto"
	"github.com/IndarMuis/gin-postgres-example.git/src/model/entity"
)

type BookRepository interface {
	FindAll() ([]*entity.Book, error)
	FindById(id uint) (*entity.Book, error)
	FindBookDetails(id uint) (*dto.BookDetailsResponse, error)
	Save(book *entity.Book, authorId int) (*entity.Book, error)
	Update(book *entity.Book) (*entity.Book, error)
	Delete(book *entity.Book) (*entity.Book, error)
}

type BookRepositoryImpl struct {
	*sql.DB
}

func (bookRepository *BookRepositoryImpl) FindBookDetails(id uint) (*dto.BookDetailsResponse, error) {
	db := bookRepository.DB

	var book entity.Book
	var author entity.Author
	if err := db.QueryRow("SELECT b.id, b.title, b.category, b.published_year, a.id, a.name FROM book AS b "+
		"INNER JOIN author AS a ON b.author_id = a.id WHERE b.id = $1", id).
		Scan(&book.ID, &book.Title, &book.Category, &book.PublishedYear, &author.ID, &author.Name); err != nil {
		panic(err)
	}

	return &dto.BookDetailsResponse{
		ID:            book.ID,
		Title:         book.Title,
		Author:        &author,
		Category:      book.Category,
		PublishedYear: book.PublishedYear,
	}, nil
}

func (bookRepository *BookRepositoryImpl) FindAll() ([]*entity.Book, error) {
	db := bookRepository.DB

	query, err := db.Query("SELECT b.id, b.title, b.author_id, b.category, b.published_year FROM book b INNER JOIN author a ON a.id = b.author_id")
	if err != nil {
		return nil, err
	}

	var books []*entity.Book
	for query.Next() {
		var book entity.Book
		err := query.Scan(&book.ID, &book.Title, &book.AuthorId, &book.Category, &book.PublishedYear)
		if err != nil {
			panic(err)
		}
		books = append(books, &book)
	}

	return books, nil
}

func (bookRepository *BookRepositoryImpl) FindById(id uint) (*entity.Book, error) {
	db := bookRepository.DB

	var book entity.Book
	if err := db.QueryRow("SELECT id, title, author_id, category, published_year FROM book WHERE id = $1", id).
		Scan(&book.ID, &book.Title, &book.AuthorId, &book.Category, &book.PublishedYear); err != nil {
		return new(entity.Book), nil
	}

	return &book, nil
}

func (bookRepository *BookRepositoryImpl) Save(book *entity.Book, authorId int) (*entity.Book, error) {
	// init DB
	db := bookRepository.DB

	// get author
	var author entity.Author
	row := db.QueryRow("SELECT * FROM author WHERE id = $1", authorId)
	err := row.Scan(&author.ID, &author.Name)
	if err != nil {
		return nil, err
	}

	// insert new book
	var lastInsertId int
	query := "INSERT INTO book(title, author_id, category, published_year) VALUES($1, $2, $3, $4) RETURNING id"
	err = db.QueryRow(query, book.Title, author.ID, book.Category, book.PublishedYear).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	// build response
	book.ID = uint(lastInsertId)
	book.AuthorId = author.ID
	return book, nil
}

func (bookRepository *BookRepositoryImpl) Update(book *entity.Book) (*entity.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (bookRepository *BookRepositoryImpl) Delete(book *entity.Book) (*entity.Book, error) {
	//TODO implement me
	panic("implement me")
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &BookRepositoryImpl{DB: db}
}
