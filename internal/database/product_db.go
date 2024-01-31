package database

import (
	"database/sql"

	"github.com/cDreyer00/devfullcycle/gopai/internal/entity"
)

const queryItems = `SELECT id, name, description, price, category_id, image_url FROM products`

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (pd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := pd.db.Query(queryItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (pd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product
	if err := pd.db.QueryRow(queryItems+" WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL); err != nil {
		return nil, err
	}
	return &product, nil
}

func (pd *ProductDB) GetProductsByCategoryID(categoryID string) ([]*entity.Product, error) {
	rows, err := pd.db.Query(queryItems+` WHERE category_id = ?`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (pd *ProductDB) CreateProduct(product *entity.Product) (string, error) {
	_, err := pd.db.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) VALUES (?, ?, ?, ?, ?, ?)", product.ID, product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)
	if err != nil {
		return "", err
	}
	return product.ID, nil
}
