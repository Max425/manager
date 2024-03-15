package dto

type Company struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Positions   []string `json:"positions"`
	Image       string   `json:"image,omitempty"`
	Description string   `json:"description,omitempty"`
}
