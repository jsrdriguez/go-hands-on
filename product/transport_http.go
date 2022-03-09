package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	r.Method(http.MethodGet, "/{id}", kithttp.NewServer(
		makeGetProductByIdEndPoint(s),
		getProductByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	r.Method(http.MethodPost, "/paginated", kithttp.NewServer(
		makGetProductsEndPoint(s),
		getProductsRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	r.Method(http.MethodPost, "/", kithttp.NewServer(
		makAddProductsEndPoint(s),
		addProductRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	r.Method(http.MethodPut, "/", kithttp.NewServer(
		makeUpdateProductsEndPoint(s),
		updateProductRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	r.Method(http.MethodDelete, "/{id}", kithttp.NewServer(
		makeDeleteProductsEndPoint(s),
		deleteProductRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	return r
}

func deleteProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	return deleteProductRequest{
		ProductId: id,
	}, nil
}

func updateProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateProductRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	return request, nil
}

func addProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	return request, nil
}

func getProductsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	return request, nil
}

func getProductByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	return getProductByIDRequest{
		ProductID: id,
	}, nil
}
