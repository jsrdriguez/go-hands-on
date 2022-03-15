package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jsrdriguez/go-hands-on/config"
	customer "github.com/jsrdriguez/go-hands-on/customers"
	"github.com/jsrdriguez/go-hands-on/employee"
	"github.com/jsrdriguez/go-hands-on/order"
	"github.com/jsrdriguez/go-hands-on/product"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/jsrdriguez/go-hands-on/docs"
)

//@title         Api
// @version      1.0
// @description  This is a sample server celler server.

// @contact.name   API Support
// @contact.url    https://example.com
// @contact.email  support@swagger.io

func main() {
	database := config.InitDB()
	defer database.Close()

	var (
		productRepository  = product.NewRepository(database)
		employeeRepository = employee.NewRepository(database)
		customerRepository = customer.NewRepository(database)
		orderRepository    = order.NewRepository(database)
	)

	var (
		productService  product.Service
		employeeService employee.Service
		customerService customer.Service
		orderService    order.Service
	)

	productService = product.NewService(productRepository)
	employeeService = employee.NewService(employeeRepository)
	customerService = customer.NewService(customerRepository)
	orderService = order.NewService(orderRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))
	r.Mount("/customers", customer.MakeHttpHandler(customerService))
	r.Mount("/order", order.MakeHttpHandler(orderService))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:9000/swagger/doc.json"),
	))

	fmt.Println("localhost:9000")
	http.ListenAndServe(":9000", r)

}
