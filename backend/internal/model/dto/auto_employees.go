package dto

type AutoEmployees struct {
	Project      Project        `json:"project"`
	AutoEmployee []AutoEmployee `json:"auto_employee"`
}
