package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int
}

type getProductRequest struct {
	Limit  int
	Offset int
}

type getAddProductRequest struct {
	ProductCode  string
	ProductName  string
	Description  string
	StandardCost string
	ListPrice    string
	Category     string
}

type updateProductRequest struct {
	ID           int64
	ProductCode  string
	ProductName  string
	Description  string
	StandardCost string
	ListPrice    string
	Category     string
}

type deleteProductRequest struct {
	ProductId int
}

func makeDeleteProductsEndPoint(s Service) endpoint.Endpoint {
	deleteProductByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductRequest)

		productId, err := s.DeleteProduct(&req)
		if err != nil {
			panic(nil)
		}

		return productId, nil

	}

	return deleteProductByIdEndPoint
}

func makeUpdateProductsEndPoint(s Service) endpoint.Endpoint {
	updateProductByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateProductRequest)

		productId, err := s.UpdateProduct(&req)
		if err != nil {
			panic(nil)
		}

		return productId, nil

	}

	return updateProductByIdEndPoint
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

func makGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductRequest)

		result, err := s.GetProducts(&req)
		if err != nil {
			panic(nil)
		}

		return result, nil

	}

	return getProductsByIdEndPoint
}

func makAddProductsEndPoint(s Service) endpoint.Endpoint {
	addProductByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)

		productId, err := s.InsertProduct(&req)
		if err != nil {
			panic(nil)
		}

		return productId, nil

	}

	return addProductByIdEndPoint
}
