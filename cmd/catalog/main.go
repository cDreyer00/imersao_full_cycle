package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/cDreyer00/devfullcycle/gopai/internal/database"
	"github.com/cDreyer00/devfullcycle/gopai/internal/service"
	"github.com/cDreyer00/devfullcycle/gopai/internal/webserver"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	productDB := database.NewProductDB(db)

	categoryService := service.NewCategoryService(*categoryDB)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Get("/categories", webCategoryHandler.GetCategories)
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/products", webProductHandler.GetProducts)
	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product/category/{id}", webProductHandler.GetProductsByCategoryID)
	c.Post("/product", webProductHandler.CreateProduct)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", c)
}
