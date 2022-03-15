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

func makeDeleteOrderDetailEnpoint(s Service) endpoint.Endpoint {
	deleteOrderDetail := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderDetailRequest)

		result, err := s.DeleteOrderDetail(&req)
		helpers.Catch(err)

		return result, nil
	}

	return deleteOrderDetail
}

func makeDeleteOrderEnpoint(s Service) endpoint.Endpoint {
	deleteOrder := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderRequest)

		result, err := s.DeleteOrder(&req)
		helpers.Catch(err)

		return result, nil
	}

	return deleteOrder
}

func makeUpdateOrderEnpoint(s Service) endpoint.Endpoint {
	UpdateOrder := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)

		result, err := s.UpdateOrder(&req)
		helpers.Catch(err)

		return result, nil
	}

	return UpdateOrder
}

func makeAddOrderEnpoint(s Service) endpoint.Endpoint {
	addOrder := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)

		result, err := s.InsertOrder(&req)
		helpers.Catch(err)

		return result, nil
	}

	return addOrder
}

func makeGetOrdersEndPoint(s Service) endpoint.Endpoint {
	getOrders := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrdersRequest)

		result, err := s.GetOrders(&req)
		helpers.Catch(err)

		return result, nil
	}

	return getOrders
}

func makeGetOrderByIdEndpoint(s Service) endpoint.Endpoint {
	getOrderByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrderByIdRequest)

		result, err := s.GetOrderById(&req)
		helpers.Catch(err)

		return result, nil
	}

	return getOrderByIdEndpoint
}
