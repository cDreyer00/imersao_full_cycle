package entity

import "github.com/google/uuid"

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  string  `json:"category_id"`
	ImageURL    string  `json:"image_url"`
}

func NewProduct(name string, description string, price float64, categoryID string, imageUrl string) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Price:       price,
		Description: description,
		CategoryID:  categoryID,
		ImageURL:    imageUrl,
	}
}
