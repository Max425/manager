package dto

type AutoEmployee struct {
	Position string   `json:"position"`
	Employee Employee `json:"employee"`
	Pin      bool     `json:"pin"`
}
