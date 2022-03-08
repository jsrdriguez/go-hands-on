package product

import "database/sql"

type Repository interface {
	GetProductById(productId int) (*Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &repository{db: database}
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
	if err != nil {
		panic(err)
	}

	return product, nil
}
