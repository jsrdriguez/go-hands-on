package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jsrdriguez/go-hands-on/config"
	"github.com/jsrdriguez/go-hands-on/utils"
)

var db *sql.DB

type Product struct {
	ID           int    `json:"id"`
	Product_Code string `json:"product_code"`
	Description  string `json:"description"`
}

func AllProducts(w http.ResponseWriter, r *http.Request) {
	sql := "SELECT id, product_code, COALESCE(description, '') FROM products"

	results, err := db.Query(sql)
	utils.Catch(err)

	products := []*Product{}

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.Product_Code, &product.Description)
		utils.Catch(err)

		products = append(products, product)
	}

	utils.ResponseWidthJSON(w, http.StatusOK, products)
}

func StoreProduct(w http.ResponseWriter, r *http.Request) {
	var product Product

	json.NewDecoder(r.Body).Decode(&product)

	sql := "INSERT products SET product_code = ?, description = ?"

	result, err := db.Prepare(sql)
	utils.Catch(err)

	_, err = result.Exec(product.Product_Code, product.Description)

	utils.Catch(err)

	defer result.Close()

	utils.ResponseWidthJSON(w, http.StatusCreated, map[string]string{"mesage": "successfully created"})
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product

	id := chi.URLParam(r, "id")

	json.NewDecoder(r.Body).Decode(&product)

	sql := "UPDATE products SET product_code = ?, description = ? WHERE id=?"

	result, err := db.Prepare(sql)
	utils.Catch(err)

	_, err = result.Exec(product.Product_Code, product.Description, id)

	utils.Catch(err)

	defer result.Close()

	utils.ResponseWidthJSON(w, http.StatusCreated, map[string]string{"mesage": "update successfully"})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	result, err := db.Prepare("DELETE FROM products WHERE id=?")
	utils.Catch(err)

	_, err = result.Exec(id)

	utils.Catch(err)

	defer result.Close()

	utils.ResponseWidthJSON(w, http.StatusCreated, map[string]string{"mesage": "delete successfully"})
}

func main() {
	db = config.InitDB()

	defer db.Close()

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Word"))
	})

	r.Get("/products", AllProducts)
	r.Post("/products", StoreProduct)
	r.Put("/products/{id}", UpdateProduct)
	r.Delete("/products/{id}", DeleteProduct)

	log.Fatal(http.ListenAndServe(":9000", r))

}
