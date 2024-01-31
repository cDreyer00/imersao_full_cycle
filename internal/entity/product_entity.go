package entity

import "github.com/google/uuid"

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	CategoryID  string
	ImageURL    string
}

func NewProduct(name string, description string, price float64, categoryID string, imageUrl string) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
	}
}
