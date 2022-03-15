package order

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/jsrdriguez/go-hands-on/helpers"
)

type getOrderByIdRequest struct {
	OrderId int64
}

type getOrdersRequest struct {
	Limit    int
	Offset   int
	Status   interface{}
	DateFrom interface{}
	DateTo   interface{}
}

type addOrderRequest struct {
	ID           int64
	OrderDate    string
	CustomerID   int
	OrderDetails []addOrderDetailRequest
}

type addOrderDetailRequest struct {
	ID        int64
	OrderID   int64
	ProductID int64
	Quantity  int64
	UnitPrice float64
}

type deleteOrderDetailRequest struct {
	OrderDetailId int64
}

type deleteOrderRequest struct {
	OrderId int64
}

// @Sumary Eliminar OrderDetail by Id
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "OrderDetail Id"
// @Success 200 {integer} int "ok"
// @Router /detail/{id} [delete]
func makeDeleteOrderDetailEnpoint(s Service) endpoint.Endpoint {
	deleteOrderDetail := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderDetailRequest)

		result, err := s.DeleteOrderDetail(&req)
		helpers.Catch(err)

		return result, nil
	}

	return deleteOrderDetail
}

// @Sumary Eliminar Order by Id
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "Order Id"
// @Success 200 {integer} int "ok"
// @Router /Order/{id} [delete]
func makeDeleteOrderEnpoint(s Service) endpoint.Endpoint {
	deleteOrder := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderRequest)

		result, err := s.DeleteOrder(&req)
		helpers.Catch(err)

		return result, nil
	}

	return deleteOrder
}

// @Sumary Update Order
// @Tags Order
// @Accept json
// @Produce json
// @Param request body order.addOrderRequest true "Order Data"
// @Success 200 {integer} int "ok"
// @Router /order/ [put]
func makeUpdateOrderEnpoint(s Service) endpoint.Endpoint {
	UpdateOrder := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)

		result, err := s.UpdateOrder(&req)
		helpers.Catch(err)

		return result, nil
	}

	return UpdateOrder
}

// @Sumary Insertar Order
// @Tags Order
// @Accept json
// @Produce json
// @Param request body order.addOrderRequest true "Order Data"
// @Success 200 {integer} int "ok"
// @Router /order/ [post]
func makeAddOrderEnpoint(s Service) endpoint.Endpoint {
	addOrder := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)

		result, err := s.InsertOrder(&req)
		helpers.Catch(err)

		return result, nil
	}

	return addOrder
}

// @Sumary Lista de Ordenes
// @Tags Order
// @Accept json
// @Produce json
// @Param request body order.getOrdersRequest true "User Data"
// @Success 200 {object} order.OrderList "ok"
// @Router /order/paginated [post]
func makeGetOrdersEndPoint(s Service) endpoint.Endpoint {
	getOrders := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrdersRequest)

		result, err := s.GetOrders(&req)
		helpers.Catch(err)

		return result, nil
	}

	return getOrders
}

// @Sumary Order by Id
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "OrderId"
// @Success 200 {object} order.OrderItem "ok"
// @Router /order/{id} [get]
func makeGetOrderByIdEndpoint(s Service) endpoint.Endpoint {
	getOrderByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrderByIdRequest)

		result, err := s.GetOrderById(&req)
		helpers.Catch(err)

		return result, nil
	}

	return getOrderByIdEndpoint
}
