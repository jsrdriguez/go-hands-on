package product

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getProductByIDRequest := kithttp.NewServer(
		makeGetProductByIdEndPoint(s),
		getProductByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)

	r.Method(http.MethodGet, "/{id}", getProductByIDRequest)

	return r
}

func getProductByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	return getProductByIDRequest{
		ProductID: id,
	}, nil
}
