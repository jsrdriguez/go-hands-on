package customer

import (
	"database/sql"

	"github.com/jsrdriguez/go-hands-on/helpers"
)

type Repository interface {
	getCustomers(params *getCustomerRequest) ([]*Customer, error)
	getTotalCustomers() (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) getCustomers(params *getCustomerRequest) ([]*Customer, error) {
	const sql = `SELECT id, first_name, last_name, address, business_phone, city, company
	FROM customers LIMIT ? OFFSET ?`

	rows, err := r.db.Query(sql, &params.Limit, &params.Offset)
	helpers.Catch(err)

	var customers []*Customer

	for rows.Next() {
		customer := Customer{}

		err := rows.Scan(
			&customer.ID,
			&customer.FirstName,
			&customer.LastName,
			&customer.Address,
			&customer.BusinessPhone,
			&customer.City,
			&customer.Company,
		)
		helpers.Catch(err)

		customers = append(customers, &customer)
	}

	return customers, nil
}

func (r *repository) getTotalCustomers() (int64, error) {
	const sql = `SELECT COUNT(*) FROM customers`
	var total int64

	row := r.db.QueryRow(sql)
	err := row.Scan(&total)
	helpers.Catch(err)

	return total, err
}
