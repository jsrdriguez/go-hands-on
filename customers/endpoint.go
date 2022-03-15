package customer

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/jsrdriguez/go-hands-on/helpers"
)

type getCustomerRequest struct {
	Limit  int
	Offset int
}

// @Sumary Lista de Clientes
// @Tags Customers
// @Accept json
// @Produce json
// @Param request body customer.getCustomerRequest true "User Data"
// @Success 200 {object} customer.CustomerList "ok"
// @Router /customers/paginated [post]
func makeGetCustomers(s Service) endpoint.Endpoint {
	getCustomers := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCustomerRequest)

		results, err := s.getCustomers(&req)
		helpers.Catch(err)

		return results, nil
	}

	return getCustomers
}
