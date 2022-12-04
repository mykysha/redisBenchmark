package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v9"
	"github.com/mykysha/redisBenchmark/domain"
)

// Client is a client for a redis database.
type Client struct {
	client *redis.Client
}

func NewClient(connStr string) *Client {
	client := redis.NewClient(&redis.Options{
		Addr: connStr,
	})

	return &Client{client: client}
}

// CreateBook adds a new book to the database with unique id.
func (c *Client) CreateBook(book domain.Book) (int64, error) {
	ctx := context.TODO()

	id, err := c.client.Incr(ctx, "id").Result()
	if err != nil {
		return 0, fmt.Errorf("failed to increment id: %w", err)
	}

	err = c.client.Set(ctx, "book"+strconv.Itoa(int(id)), book, 0).Err()
	if err != nil {
		return 0, fmt.Errorf("failed to create book: %w", err)
	}

	return id, nil
}

// GetBook returns a book from the database.
func (c *Client) GetBook(id int64) (domain.Book, error) {
	ctx := context.TODO()

	cmd := c.client.Get(ctx, "book"+strconv.Itoa(int(id)))

	err := cmd.Err()
	if err != nil {
		return domain.Book{}, fmt.Errorf("failed to get book: %w", err)
	}

	var book domain.Book

	err = cmd.Scan(&book)
	if err != nil {
		return domain.Book{}, fmt.Errorf("failed to get book: %w", err)
	}

	return book, nil
}

// UpdateBook updates a book in the database.
func (c *Client) UpdateBook(book domain.Book) error {
	ctx := context.TODO()

	err := c.client.Set(ctx, "book"+strconv.Itoa(int(book.ID)), book, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}

	return nil
}

// DeleteBook deletes a book from the database.
func (c *Client) DeleteBook(id int64) error {
	ctx := context.TODO()

	err := c.client.Del(ctx, "book"+strconv.Itoa(int(id))).Err()
	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	return nil
}

// CreateAuthor adds a new author to the database with unique id.
func (c *Client) CreateAuthor(author domain.Author) (int64, error) {
	ctx := context.TODO()

	id, err := c.client.Incr(ctx, "id").Result()
	if err != nil {
		return 0, fmt.Errorf("failed to increment id: %w", err)
	}

	err = c.client.Set(ctx, "author"+strconv.Itoa(int(id)), author, 0).Err()
	if err != nil {
		return 0, fmt.Errorf("failed to create author: %w", err)
	}

	return id, nil
}

// GetAuthor returns an author from the database.
func (c *Client) GetAuthor(id int64) (domain.Author, error) {
	ctx := context.TODO()

	cmd := c.client.Get(ctx, "author"+strconv.Itoa(int(id)))

	err := cmd.Err()
	if err != nil {
		return domain.Author{}, fmt.Errorf("failed to get author: %w", err)
	}

	var author domain.Author

	err = cmd.Scan(&author)
	if err != nil {
		return domain.Author{}, fmt.Errorf("failed to get author: %w", err)
	}

	return author, nil
}

// UpdateAuthor updates an author in the database.
func (c *Client) UpdateAuthor(author domain.Author) error {
	ctx := context.TODO()

	err := c.client.Set(ctx, "author"+strconv.Itoa(int(author.ID)), author, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to update author: %w", err)
	}

	return nil
}

// DeleteAuthor deletes an author from the database.
func (c *Client) DeleteAuthor(id int64) error {
	ctx := context.TODO()

	err := c.client.Del(ctx, "author"+strconv.Itoa(int(id))).Err()
	if err != nil {
		return fmt.Errorf("failed to delete author: %w", err)
	}

	return nil
}

// Close closes the client.
func (c *Client) Close() error {
	err := c.client.Close()
	if err != nil {
		return fmt.Errorf("failed to close client: %w", err)
	}

	return nil
}
