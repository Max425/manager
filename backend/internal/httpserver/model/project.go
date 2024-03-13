package model

import "time"

type Project struct {
	ID           int       `db:"id" json:"id"`
	CompanyID    int       `db:"company_id" json:"company_id"`
	Name         string    `db:"name" json:"name"`
	Stages       []string  `db:"stages" json:"stages"`
	Image        string    `db:"image" json:"image,omitempty"`
	Description  string    `db:"description" json:"description,omitempty"`
	CurrentStage int       `db:"current_stage" json:"current_stage,omitempty"`
	Deadline     time.Time `db:"deadline" json:"deadline,omitempty"`
	Status       int16     `db:"status" json:"status"`
	Complexity   int16     `db:"complexity" json:"complexity,omitempty"`
}
