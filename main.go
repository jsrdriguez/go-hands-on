package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jsrdriguez/go-hands-on/config"
	"github.com/jsrdriguez/go-hands-on/product"
)

func main() {
	database := config.InitDB()
	defer database.Close()

	var productRepository = product.NewRepository(database)
	var productService product.Service

	productService = product.NewService(productRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))

	http.ListenAndServe(":9000", r)

}
