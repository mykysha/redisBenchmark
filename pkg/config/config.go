// Package config is responsible for getting needed configs.
package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Repository is an interface for getting configs.
type Repository interface {
	SetConfigFile(path string) error
	GetString(key string) string
	GetDuration(key string) time.Duration
	GetInt(key string) int
	GetStringArray(key string) []string
}

// Reader is the default implementation of the Repository interface.
type Reader struct {
	reader *viper.Viper
}

// NewReader creates a new instance of the Reader.
func NewReader() *Reader {
	return &Reader{reader: viper.New()}
}

// NewReader creates a new instance of the Reader and sets it to the config file.
func NewReaderWithPath(path string) (*Reader, error) {
	reader := &Reader{reader: viper.New()}

	err := reader.SetConfigFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to set config file: %w", err)
	}

	return reader, nil
}

// SetConfigFile defines path and name of the desired config file.
func (r *Reader) SetConfigFile(path string) error {
	r.reader.SetConfigFile(path)

	err := r.reader.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	return nil
}

// GetString reads string with the specified key from the config file.
func (r *Reader) GetString(key string) string {
	return r.reader.GetString(key)
}

// GetDuration reads time.Duration with the specified key from the config file.
func (r *Reader) GetDuration(key string) time.Duration {
	return r.reader.GetDuration(key)
}

// GetInt reads int with the specified key from the config file.
func (r *Reader) GetInt(key string) int {
	return r.reader.GetInt(key)
}

// GetStringArray reads a slice of strings from the config file.
func (r *Reader) GetStringArray(key string) []string {
	return r.reader.GetStringSlice(key)
}
