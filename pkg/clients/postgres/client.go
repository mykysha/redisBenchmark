package postgres

import (
	"database/sql"
	"fmt"

	// Import the postgres driver.
	_ "github.com/lib/pq"
	"github.com/mykysha/redisBenchmark/domain"
)

// Client is a client for the postgres database.
type Client struct {
	conn *sql.DB
}

// NewClient creates a new client for the postgres database.
func NewClient(host, port, user, pass, dbName, sslmode string) (*Client, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pass, dbName, sslmode)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %w", err)
	}

	return &Client{conn: conn}, nil
}

// CreateBook adds a new book to the database.
func (c *Client) CreateBook(book domain.Book) (int64, error) {
	var id int64

	err := c.conn.QueryRow(
		"INSERT INTO books (title, genre, author_id, year, pages) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		book.Title, book.Genre, book.AuthorID, book.Year, book.Pages).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create book: %w", err)
	}

	return id, nil
}

// GetBook returns a book from the database.
func (c *Client) GetBook(id int64) (domain.Book, error) {
	var book domain.Book

	err := c.conn.QueryRow(
		"SELECT id, title, genre, author_id, year, pages FROM books WHERE id = $1",
		id).
		Scan(&book.ID, &book.Title, &book.Genre, &book.AuthorID, &book.Year, &book.Pages)
	if err != nil {
		return domain.Book{}, fmt.Errorf("failed to get book: %w", err)
	}

	return book, nil
}

// UpdateBook updates a book in the database.
func (c *Client) UpdateBook(book domain.Book) error {
	_, err := c.conn.Exec(
		"UPDATE books SET title = $1, genre = $2, author_id = $3, year = $4, pages = $5 WHERE id = $6",
		book.Title, book.Genre, book.AuthorID, book.Year, book.Pages, book.ID)
	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}

	return nil
}

// DeleteBook deletes a book from the database.
func (c *Client) DeleteBook(id int64) error {
	_, err := c.conn.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	return nil
}

// CreateAuthor adds a new author to the database.
func (c *Client) CreateAuthor(author domain.Author) (int64, error) {
	var id int64

	err := c.conn.QueryRow(
		"INSERT INTO authors (name, surname, birth_country) VALUES ($1, $2, $3) RETURNING id",
		author.Name, author.Surname, author.BirthCountry).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create author: %w", err)
	}

	return id, nil
}

// GetAuthor returns an author from the database.
func (c *Client) GetAuthor(id int64) (domain.Author, error) {
	var author domain.Author

	err := c.conn.QueryRow(
		"SELECT id, name, surname, birth_country FROM authors WHERE id = $1",
		id).
		Scan(&author.ID, &author.Name, &author.Surname, &author.BirthCountry)
	if err != nil {
		return domain.Author{}, fmt.Errorf("failed to get author: %w", err)
	}

	return author, nil
}

// UpdateAuthor updates an author in the database.
func (c *Client) UpdateAuthor(author domain.Author) error {
	_, err := c.conn.Exec(
		"UPDATE authors SET name = $1, surname = $2, birth_country = $3 WHERE id = $4",
		author.Name, author.Surname, author.BirthCountry, author.ID)
	if err != nil {
		return fmt.Errorf("failed to update author: %w", err)
	}

	return nil
}

// DeleteAuthor deletes an author from the database.
func (c *Client) DeleteAuthor(id int64) error {
	_, err := c.conn.Exec("DELETE FROM authors WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete author: %w", err)
	}

	return nil
}

// Close closes the connection to the database.
func (c *Client) Close() error {
	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	return nil
}
