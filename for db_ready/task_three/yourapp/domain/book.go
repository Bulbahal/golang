package domain

import "context"

// модели и интерфейсы домена
type Book struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	AuthorId        int    `json:"author_id"`
	PublicationYear int    `json:"publication_year"`
}

type BookFilter struct {
	Title  *string
	Author *string
}

type BookRepo interface {
	Create(ctx context.Context, book Book) (int, error)
	GetByID(ctx context.Context, id int) (*Book, error)
	List(ctx context.Context, f BookFilter) ([]Book, error)
	Update(ctx context.Context, book Book) error
	Delete(ctx context.Context, id int) error
}
