package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int
}

func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {

	getProductByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)

		product, err := s.GetProductById(&req)

		if err != nil {
			panic(nil)
		}

		return product, nil
	}

	return getProductByIdEndPoint
}
