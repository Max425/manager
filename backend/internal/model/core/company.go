package core

import (
	"fmt"
	"time"
)

type Company struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Positions   string    `db:"positions"`
	Image       string    `db:"image"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func NewCompany(id int, name, positions, image, description string) (*Company, error) {
	if name == "" {
		return nil, fmt.Errorf("%w: company name is required", ErrRequired)
	}

	return &Company{
		ID:          id,
		Name:        name,
		Positions:   positions,
		Image:       image,
		Description: description,
	}, nil
}
