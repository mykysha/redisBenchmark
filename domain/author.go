package domain

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	ID           int64
	Name         string
	Surname      string
	BirthCountry string
}

func (a Author) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal author: %w", err)
	}

	return data, nil
}

func (a *Author) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, a)
	if err != nil {
		return fmt.Errorf("failed to unmarshal author: %w", err)
	}

	return nil
}
