package model

type Company struct {
	ID          int      `db:"id" json:"id"`
	Name        string   `db:"name" json:"name"`
	Positions   []string `db:"positions" json:"positions"`
	Image       string   `db:"image" json:"image,omitempty"`
	Description string   `db:"description" json:"description,omitempty"`
}