package customer

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/jsrdriguez/go-hands-on/helpers"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	r.Method(http.MethodPost, "/paginated", kithttp.NewServer(
		makeGetCustomers(s),
		getCustomerRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	return r
}

func getCustomerRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getCustomerRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helpers.Catch(err)

	return request, nil
}
