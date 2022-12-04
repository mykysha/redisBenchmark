package domain

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	ID       int64
	AuthorID int64
	Year     int
	Pages    int
	Title    string
	Genre    string
}

func (b Book) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(b)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal book: %w", err)
	}

	return data, nil
}

func (b *Book) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, b)
	if err != nil {
		return fmt.Errorf("failed to unmarshal book: %w", err)
	}

	return nil
}
