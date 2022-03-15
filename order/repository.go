package order

import (
	"database/sql"
	"fmt"

	"github.com/jsrdriguez/go-hands-on/helpers"
)

type Repository interface {
	GetOrderById(params *getOrderByIdRequest) (*OrderItem, error)
	GetOrders(params *getOrdersRequest) (*OrderList, error)
	InsertOrder(params *addOrderRequest) (int64, error)
	InsertOrderDetailt(params *addOrderDetailRequest) (int64, error)

	UpdateOrder(params *addOrderRequest) (int64, error)
	UpdateOrderDetailt(params *addOrderDetailRequest) (int64, error)
	DeleteOrderDetail(params *deleteOrderDetailRequest) (int64, error)

	DeleteOrder(params *deleteOrderRequest) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) DeleteOrder(params *deleteOrderRequest) (int64, error) {
	const sqlOrder = `DELETE FROM orders WHERE id = ?`
	const sqlOrderDetail = `DELETE FROM order_details WHERE order_id = ?`

	_, err := r.db.Exec(sqlOrderDetail, params.OrderId)
	helpers.Catch(err)

	result_a, err := r.db.Exec(sqlOrder, params.OrderId)
	helpers.Catch(err)

	count, err := result_a.RowsAffected()

	return count, nil
}

func (r *repository) DeleteOrderDetail(params *deleteOrderDetailRequest) (int64, error) {
	const sql = `DELETE FROM order_details WHERE id = ?`

	result, err := r.db.Exec(sql, params.OrderDetailId)
	helpers.Catch(err)

	deleted, err := result.RowsAffected()

	return deleted, nil
}

func (r *repository) UpdateOrderDetailt(params *addOrderDetailRequest) (int64, error) {
	const sql = `UPDATE order_details SET quantity = ?, unit_price = ? WHERE id = ?`

	_, err := r.db.Exec(sql, params.Quantity, params.UnitPrice, params.ID)
	helpers.Catch(err)

	return params.ID, nil
}

func (r *repository) UpdateOrder(params *addOrderRequest) (int64, error) {
	const sql = `UPDATE orders SET customer_id = ? WHERE id = ?`

	_, err := r.db.Exec(sql, params.CustomerID, params.ID)
	helpers.Catch(err)

	return params.ID, nil
}

func (r *repository) InsertOrder(params *addOrderRequest) (int64, error) {
	const sql = `INSERT INTO orders(customer_id, order_date) VALUES(?, ?)`

	result, err := r.db.Exec(sql, &params.CustomerID, &params.OrderDate)
	helpers.Catch(err)

	id, err := result.LastInsertId()
	helpers.Catch(err)

	return id, nil
}

func (r *repository) InsertOrderDetailt(params *addOrderDetailRequest) (int64, error) {
	const sql = `INSERT INTO order_details(order_id, product_id, quantity, unit_price) VALUES(?, ?, ?, ?)`

	result, err := r.db.Exec(sql, &params.OrderID, &params.ProductID, &params.Quantity, &params.UnitPrice)
	helpers.Catch(err)

	detailtId, err := result.LastInsertId()
	helpers.Catch(err)

	return detailtId, nil
}

func (r *repository) GetOrders(params *getOrdersRequest) (*OrderList, error) {
	var filter string

	if params.Status != nil {
		filter += fmt.Sprintf(" AND o.status_id = %v ", params.Status.(float64))
	}

	if params.DateFrom != nil && params.DateTo == nil {
		filter += fmt.Sprintf(" AND o.order_date >= '%v' ", params.DateFrom.(string))
	}

	if params.DateFrom == nil && params.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date <= '%v' ", params.DateTo.(string))
	}

	if params.DateFrom != nil && params.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date BETWEEN '%v' and '%v' ", params.DateFrom.(string), params.DateTo.(string))
	}

	var sql = `SELECT o.id, o.customer_id, o.order_date, o.status_id, os.status_name, 
	CONCAT(c.first_name,' ', c.last_name) as customer
	FROM orders o
	INNER JOIN orders_status os ON o.status_id = os.id
	INNER JOIN customers c ON o.customer_id = c.id
	WHERE 1=1 `

	sql = sql + filter + " ORDER BY o.id DESC LIMIT ? OFFSET ? "

	results, err := r.db.Query(sql, &params.Limit, &params.Offset)
	helpers.Catch(err)

	var orders []*OrderItem

	for results.Next() {
		item := OrderItem{}

		err := results.Scan(
			&item.ID,
			&item.CustomerId,
			&item.OrderDate,
			&item.StatusId,
			&item.StatusName,
			&item.Customer,
		)
		helpers.Catch(err)

		details, _ := GetOrderDetail(r, &item.ID)
		item.Data = details

		orders = append(orders, &item)
	}

	var sqlTotal = "SELECT COUNT(*) FROM orders WHERE 1=1" + filter
	row := r.db.QueryRow(sqlTotal)

	var total int64
	err = row.Scan(&total)
	helpers.Catch(err)

	return &OrderList{
		Data:         orders,
		TotalRecords: total,
	}, nil
}

func (r *repository) GetOrderById(params *getOrderByIdRequest) (*OrderItem, error) {
	const sqlOrder = `SELECT o.id, o.customer_id, o.order_date, o.status_id, os.status_name, 
	CONCAT(c.first_name,' ', c.last_name) as customer, c.company, c.address, c.business_phone as phone, c.city
	FROM orders o
	INNER JOIN orders_status os ON o.status_id = os.id
	INNER JOIN customers c ON o.customer_id = c.id
	WHERE o.id = ?`

	row := r.db.QueryRow(sqlOrder, &params.OrderId)
	order := &OrderItem{}

	err := row.Scan(
		&order.ID,
		&order.CustomerId,
		&order.OrderDate,
		&order.StatusId,
		&order.StatusName,
		&order.Customer,
		&order.Company,
		&order.Address,
		&order.Phone,
		&order.City,
	)
	helpers.Catch(err)

	details, err := GetOrderDetail(r, &params.OrderId)
	helpers.Catch(err)

	order.Data = details

	return order, nil
}

func GetOrderDetail(r *repository, orderId *int64) ([]*OrderDetailItem, error) {
	const sqlDetails = `SELECT order_id, od.id, quantity, unit_price, p.product_name, product_id
	FROM order_details od
	INNER JOIN products p ON od.product_id = p.id
	WHERE od.order_id = ?`

	results, err := r.db.Query(sqlDetails, &orderId)
	helpers.Catch(err)

	var orderDetails []*OrderDetailItem

	for results.Next() {
		orderDetail := &OrderDetailItem{}

		err := results.Scan(
			&orderDetail.OrderId,
			&orderDetail.ID,
			&orderDetail.Quantity,
			&orderDetail.UnitPrice,
			&orderDetail.ProductName,
			&orderDetail.ProductId,
		)
		helpers.Catch(err)

		orderDetails = append(orderDetails, orderDetail)
	}

	return orderDetails, nil
}
