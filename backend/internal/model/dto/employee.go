package dto

type Employee struct {
	ID        int     `json:"id"`
	CompanyID int     `json:"company_id"`
	Name      string  `json:"name"`
	Position  string  `json:"position,omitempty"`
	Mail      string  `json:"mail,omitempty"`
	Password  string  `json:"-"`
	Salt      string  `json:"-"`
	Image     string  `json:"image,omitempty"`
	Rating    float64 `json:"rating,omitempty"`
}
