package clients

import "github.com/mykysha/redisBenchmark/domain"

type Client interface {
	CreateBook(book domain.Book) (int64, error)
	GetBook(id int64) (domain.Book, error)
	UpdateBook(book domain.Book) error
	DeleteBook(id int64) error
	CreateAuthor(author domain.Author) (int64, error)
	GetAuthor(id int64) (domain.Author, error)
	UpdateAuthor(author domain.Author) error
	DeleteAuthor(id int64) error
	Close() error
}
