package employee

import (
	"database/sql"

	"github.com/jsrdriguez/go-hands-on/helpers"
)

type Repository interface {
	GetEmployees(params *getEmployeesRequest) ([]*Employee, error)
	GetTotalEmployees() (int64, error)
	GetEmployeeById(params *getEmployeesByIdRequest) (*Employee, error)
	GetBestEmployee(params *getEmployeesBestRequest) (*BestEmployee, error)
	AddEmployee(params *addEmployeesRequest) (int64, error)
	UpdateEmployee(params *updateEmployeesRequest) (int64, error)
	DeleteEmployee(params *deleteEmployeesRequest) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) DeleteEmployee(params *deleteEmployeesRequest) (int64, error) {
	const sql = `DELETE FROM employees WHERE id = ?`

	result, err := r.db.Exec(sql, &params.EmployeeId)
	helpers.Catch(err)

	count, err := result.RowsAffected()

	return count, err
}

func (r *repository) UpdateEmployee(params *updateEmployeesRequest) (int64, error) {
	const sql = `UPDATE employees SET 
	first_name = ?, 
	last_name = ?, 
	company = ?, 
	email_address = ?, 
	job_title = ?, 
	business_phone = ?, 
	home_phone = ?, 
	mobile_phone = ?, 
	fax_number = ?, 
	address = ? WHERE id = ?`

	_, err := r.db.Exec(sql,
		&params.FirstName,
		&params.LastName,
		&params.Company,
		&params.EmailAddress,
		&params.JobTitle,
		&params.BusinessPhone,
		&params.HomePhone,
		&params.MobilePhone,
		&params.FaxNumber,
		&params.Address,
		&params.ID,
	)
	helpers.Catch(err)

	return params.ID, err
}

func (r *repository) AddEmployee(params *addEmployeesRequest) (int64, error) {
	const sql = `INSERT INTO 
	employees (first_name, last_name, company, email_address, job_title, business_phone, home_phone, mobile_phone, fax_number, address)
	VALUES(?,?,?,?,?,?,?,?,?,?)`

	result, err := r.db.Exec(sql,
		&params.FirstName,
		&params.LastName,
		&params.Company,
		&params.EmailAddress,
		&params.JobTitle,
		&params.BusinessPhone,
		&params.HomePhone,
		&params.MobilePhone,
		&params.FaxNumber,
		&params.Address,
	)

	helpers.Catch(err)

	id, err := result.LastInsertId()
	helpers.Catch(err)

	return id, err
}

func (r *repository) GetBestEmployee(params *getEmployeesBestRequest) (*BestEmployee, error) {
	const sql = `SELECT e.id, count(e.id) as totalVentas, e.first_name, e.last_name
							FROM orders o
							INNER JOIN employees e ON o.employee_id = e.id
							GROUP BY o.employee_id
							ORDER BY totalVentas desc
							LIMIT 1`

	row := r.db.QueryRow(sql)
	bestemployee := &BestEmployee{}

	err := row.Scan(
		&bestemployee.ID,
		&bestemployee.TotalVentas,
		&bestemployee.FirstName,
		&bestemployee.LastName,
	)
	helpers.Catch(err)

	return bestemployee, err
}

func (r *repository) GetEmployeeById(params *getEmployeesByIdRequest) (*Employee, error) {
	const sql = `SELECT id, first_name, last_name, company, email_address, job_title, business_phone, 
							home_phone, COALESCE(mobile_phone, ''), fax_number, address 
							FROM employees
							WHERE id = ?`

	row := r.db.QueryRow(sql, params.EmployeeId)
	employee := &Employee{}

	err := row.Scan(
		&employee.ID, &employee.FirstName, &employee.LastName, &employee.Company, &employee.EmailAddress,
		&employee.JobTitle, &employee.BusinessPhone, &employee.HomePhone, &employee.MobilePhone,
		&employee.FaxNumber, &employee.Address,
	)

	helpers.Catch(err)

	return employee, err
}

func (r *repository) GetEmployees(params *getEmployeesRequest) ([]*Employee, error) {
	const sql = `SELECT id, first_name, last_name, company, email_address, job_title, business_phone, 
							 home_phone, COALESCE(mobile_phone, ''), fax_number, address
							 FROM employees
							 LIMIT ? OFFSET ?`

	rows, err := r.db.Query(sql, &params.Limit, &params.Offset)
	helpers.Catch(err)

	var employees []*Employee

	for rows.Next() {
		employee := &Employee{}

		err = rows.Scan(
			&employee.ID, &employee.FirstName, &employee.LastName, &employee.Company, &employee.EmailAddress,
			&employee.JobTitle, &employee.BusinessPhone, &employee.HomePhone, &employee.MobilePhone,
			&employee.FaxNumber, &employee.Address,
		)
		helpers.Catch(err)

		employees = append(employees, employee)
	}

	return employees, err
}

func (r *repository) GetTotalEmployees() (int64, error) {
	const sql = "SELECT COUNT(*) FROM employees"
	var total int64

	row := r.db.QueryRow(sql)
	err := row.Scan(&total)
	helpers.Catch(err)

	return total, nil
}
