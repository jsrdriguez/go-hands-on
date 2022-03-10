package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jsrdriguez/go-hands-on/config"
	"github.com/jsrdriguez/go-hands-on/employee"
	"github.com/jsrdriguez/go-hands-on/product"
)

func main() {
	database := config.InitDB()
	defer database.Close()

	var (
		productRepository  = product.NewRepository(database)
		employeeRepository = employee.NewRepository(database)
	)

	var (
		productService  product.Service
		employeeService employee.Service
	)

	productService = product.NewService(productRepository)
	employeeService = employee.NewService(employeeRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))

	http.ListenAndServe(":9000", r)

}
