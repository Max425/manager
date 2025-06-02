package dto

type Employee struct {
	ID                   int    `json:"id"`
	CompanyID            int    `json:"company_id"`
	Name                 string `json:"name"`
	Position             string `json:"position,omitempty"`
	Mail                 string `json:"mail,omitempty"`
	Password             string `json:"-"`
	Salt                 string `json:"-"`
	Image                string `json:"image,omitempty"`
	Rating               []int  `json:"rating,omitempty"`
	ActiveProjectsCount  int    `json:"active_projects_count"`
	OverdueProjectsCount int    `json:"overdue_projects_count"`
	TotalProjectsCount   int    `json:"total_projects_count"`
}
