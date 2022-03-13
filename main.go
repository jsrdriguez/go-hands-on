package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jsrdriguez/go-hands-on/config"
	customer "github.com/jsrdriguez/go-hands-on/customers"
	"github.com/jsrdriguez/go-hands-on/employee"
	"github.com/jsrdriguez/go-hands-on/product"
)

func main() {
	database := config.InitDB()
	defer database.Close()

	var (
		productRepository  = product.NewRepository(database)
		employeeRepository = employee.NewRepository(database)
		customerRepository = customer.NewRepository(database)
	)

	var (
		productService  product.Service
		employeeService employee.Service
		customerService customer.Service
	)

	productService = product.NewService(productRepository)
	employeeService = employee.NewService(employeeRepository)
	customerService = customer.NewService(customerRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))
	r.Mount("/customers", customer.MakeHttpHandler(customerService))

	fmt.Println("localhost:9000")
	http.ListenAndServe(":9000", r)

}
