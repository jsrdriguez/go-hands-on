package product

import "database/sql"

type Repository interface {
	GetProductById(productId int) (*Product, error)
	GetProducts(params *getProductRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &repository{db: database}
}

func (r *repository) InsertProduct(params *getAddProductRequest) (int64, error) {
	const sql = `INSERT INTO 
							 products(product_code, product_name, description, standard_cost, list_price, category)
							 VALUES(?,?,?,?,?,?)`

	result, err := r.db.Exec(sql,
		&params.ProductCode,
		&params.ProductName,
		&params.Description,
		&params.StandardCost,
		&params.ListPrice,
		&params.Category,
	)
	if err != nil {
		panic(err)
	}

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
	if err != nil {
		panic(err)
	}

	return product, nil
}

func (r *repository) GetProducts(params *getProductRequest) ([]*Product, error) {
	sql := `SELECT id, product_code, product_name, COALESCE(description, ''), standard_cost, list_price,
					category 
					FROM products 
					ORDER BY id LIMIT ? OFFSET ?`

	rows, err := r.db.Query(sql, params.Limit, params.Offset)
	if err != nil {
		panic(err)
	}

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

		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *repository) GetTotalProducts() (int, error) {
	sql := "SELECT COUNT(*) FROM products"
	var total int

	row := r.db.QueryRow(sql)
	err := row.Scan(&total)

	if err != nil {
		panic(err)
	}

	return total, nil
}
