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

func makeGetCustomers(s Service) endpoint.Endpoint {
	getCustomers := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCustomerRequest)

		results, err := s.getCustomers(&req)
		helpers.Catch(err)

		return results, nil
	}

	return getCustomers
}
