package schema

import (
	"errors"
	"strings"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (t *Todo) Validate() error {
	if strings.TrimSpace(t.Title) == "" {
		return errors.New("Title is required")
	}
	return nil
}
