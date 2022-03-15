package order

import "github.com/jsrdriguez/go-hands-on/helpers"

type Service interface {
	GetOrderById(params *getOrderByIdRequest) (*OrderItem, error)
	GetOrders(params *getOrdersRequest) (*OrderList, error)
	InsertOrder(params *addOrderRequest) (int64, error)
	UpdateOrder(params *addOrderRequest) (int64, error)
	DeleteOrderDetail(params *deleteOrderDetailRequest) (int64, error)
	DeleteOrder(params *deleteOrderRequest) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) DeleteOrderDetail(params *deleteOrderDetailRequest) (int64, error) {
	return s.repo.DeleteOrderDetail(params)
}

func (s *service) DeleteOrder(params *deleteOrderRequest) (int64, error) {
	return s.repo.DeleteOrder(params)
}

func (s *service) UpdateOrder(params *addOrderRequest) (int64, error) {
	orderId, err := s.repo.UpdateOrder(params)
	helpers.Catch(err)

	for _, detail := range params.OrderDetails {
		detail.OrderID = orderId
		if detail.ID == 0 {
			_, err := s.repo.InsertOrderDetailt(&detail)
			helpers.Catch(err)
		} else {
			_, err := s.repo.UpdateOrderDetailt(&detail)
			helpers.Catch(err)
		}
	}

	return orderId, nil
}

func (s *service) InsertOrder(params *addOrderRequest) (int64, error) {
	orderId, err := s.repo.InsertOrder(params)
	helpers.Catch(err)

	for _, detail := range params.OrderDetails {
		detail.OrderID = orderId
		_, err := s.repo.InsertOrderDetailt(&detail)
		helpers.Catch(err)
	}

	return orderId, nil
}

func (s *service) GetOrders(params *getOrdersRequest) (*OrderList, error) {
	return s.repo.GetOrders(params)
}

func (s *service) GetOrderById(params *getOrderByIdRequest) (*OrderItem, error) {
	return s.repo.GetOrderById(params)
}
