package product

import (
	"database/sql"

	"github.com/jsrdriguez/go-hands-on/helpers"
)

type Repository interface {
	GetProductById(productId int) (*Product, error)
	GetProducts(params *getProductRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *updateProductRequest) (int64, error)
	DeleteProduct(params *deleteProductRequest) (int64, error)
	GetBetSellers() ([]*ProductTop, error)
	GetTotalVentas() (float64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &repository{db: database}
}

func (r *repository) GetTotalVentas() (float64, error) {
	sql := "SELECT SUM(od.quantity * od.unit_price) vendido FROM order_details od"
	var total float64

	row := r.db.QueryRow(sql)
	err := row.Scan(&total)

	helpers.Catch(err)

	return total, nil
}

func (r *repository) GetBetSellers() ([]*ProductTop, error) {
	sql := `SELECT od.product_id, p.product_name, SUM(od.quantity * od.unit_price) AS vendido 
					FROM order_details od
					INNER JOIN products p ON od.product_id = p.id
					GROUP BY od.product_id
					ORDER BY vendido desc LIMIT 10`

	rows, err := r.db.Query(sql)
	helpers.Catch(err)

	var products []*ProductTop

	for rows.Next() {
		product := &ProductTop{}

		err = rows.Scan(
			&product.ID,
			&product.ProductName,
			&product.Vendidos,
		)

		helpers.Catch(err)

		products = append(products, product)
	}

	return products, nil
}

func (r *repository) DeleteProduct(params *deleteProductRequest) (int64, error) {
	const sql = `DELETE FROM products WHERE id = ?`

	result, err := r.db.Exec(sql, params.ProductId)
	helpers.Catch(err)

	count, _ := result.RowsAffected()

	return count, nil
}

func (r *repository) UpdateProduct(params *updateProductRequest) (int64, error) {
	const sql = `UPDATE products
							 SET 
							 product_code = ?, 
							 product_name = ?, 
							 description = ?, 
							 standard_cost = ?, 
							 list_price = ?, 
							 category = ?
							 WHERE id = ?`

	_, err := r.db.Exec(sql,
		params.ProductCode,
		params.ProductName,
		params.Description,
		params.StandardCost,
		params.ListPrice,
		params.Category,
		params.ID,
	)
	helpers.Catch(err)

	id := params.ID

	return id, nil
}

func (r *repository) InsertProduct(params *getAddProductRequest) (int64, error) {
	const sql = `INSERT INTO 
							 products(product_code, product_name, description, standard_cost, list_price, category)
							 VALUES(?,?,?,?,?,?)`

	result, err := r.db.Exec(sql,
		params.ProductCode,
		params.ProductName,
		params.Description,
		params.StandardCost,
		params.ListPrice,
		params.Category,
	)
	helpers.Catch(err)

	id, _ := result.LastInsertId()

	return id, nil
}

func (r *repository) GetProductById(productId int) (*Product, error) {
	sql := `SELECT id, product_code, product_name, COALESCE(description, ''), standard_cost, list_price,
					category FROM products WHERE id=?`

	row := r.db.QueryRow(sql, productId)
	product := &Product{}

	err := row.Scan(
		&product.Id,
		&product.ProductCode,
		&product.ProductName,
		&product.Description,
		&product.StandardCost,
		&product.ListPrice,
		&product.Category,
	)
	helpers.Catch(err)

	return product, nil
}

func (r *repository) GetProducts(params *getProductRequest) ([]*Product, error) {
	sql := `SELECT id, product_code, product_name, COALESCE(description, ''), standard_cost, list_price,
					category 
					FROM products 
					ORDER BY id LIMIT ? OFFSET ?`

	rows, err := r.db.Query(sql, params.Limit, params.Offset)
	helpers.Catch(err)

	var products []*Product

	for rows.Next() {
		product := &Product{}

		err = rows.Scan(
			&product.Id,
			&product.ProductCode,
			&product.ProductName,
			&product.Description,
			&product.StandardCost,
			&product.ListPrice,
			&product.Category,
		)

		helpers.Catch(err)

		products = append(products, product)
	}

	return products, nil
}

func (r *repository) GetTotalProducts() (int, error) {
	sql := "SELECT COUNT(*) FROM products"
	var total int

	row := r.db.QueryRow(sql)
	err := row.Scan(&total)

	helpers.Catch(err)

	return total, nil
}
