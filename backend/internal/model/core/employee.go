package core

import "time"

type Employee struct {
	ID                   int       `db:"id"`
	CompanyID            int       `db:"company_id"`
	Name                 string    `db:"name"`
	Position             string    `db:"position"`
	Mail                 string    `db:"mail"`
	Password             string    `db:"password"`
	Salt                 string    `db:"salt"`
	Image                string    `db:"image"`
	Rating               float64   `db:"rating"`
	ActiveProjectsCount  int       `db:"active_projects_count"`
	OverdueProjectsCount int       `db:"overdue_projects_count"`
	TotalProjectsCount   int       `db:"total_projects_count"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
}
