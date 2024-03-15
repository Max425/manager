package core

import "time"

type Project struct {
	ID           int       `db:"id"`
	CompanyID    int       `db:"company_id"`
	Name         string    `db:"name"`
	Stages       string    `db:"stages"`
	Image        string    `db:"image"`
	Description  string    `db:"description"`
	CurrentStage int       `db:"current_stage"`
	Deadline     time.Time `db:"deadline"`
	Status       int16     `db:"status"`
	Complexity   int16     `db:"complexity"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
