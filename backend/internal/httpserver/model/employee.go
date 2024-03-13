package model

type Employee struct {
	ID        int     `db:"id" json:"id"`
	CompanyID int     `db:"company_id" json:"company_id"`
	Name      string  `db:"name" json:"name"`
	Position  string  `db:"position" json:"position,omitempty"`
	Mail      string  `db:"mail" json:"mail,omitempty"`
	Password  string  `db:"password" json:"-"`
	Salt      string  `db:"salt" json:"-"`
	Image     string  `db:"image" json:"image,omitempty"`
	Rating    float64 `db:"rating" json:"rating,omitempty"`
}
