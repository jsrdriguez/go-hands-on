package order

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/jsrdriguez/go-hands-on/helpers"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	r.Method(http.MethodGet, "/{id}", kithttp.NewServer(
		makeGetOrderByIdEndpoint(s),
		getOrderByIdRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	r.Method(http.MethodPost, "/paginated", kithttp.NewServer(
		makeGetOrdersEndPoint(s),
		getOrdersRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	r.Method(http.MethodPost, "/", kithttp.NewServer(
		makeAddOrderEnpoint(s),
		addOrderRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	r.Method(http.MethodPut, "/", kithttp.NewServer(
		makeUpdateOrderEnpoint(s),
		updateOrderRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	r.Method(http.MethodDelete, "/detail/{id}", kithttp.NewServer(
		makeDeleteOrderDetailEnpoint(s),
		deleteOrderDetailRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	r.Method(http.MethodDelete, "/{id}", kithttp.NewServer(
		makeDeleteOrderEnpoint(s),
		deleteOrderRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	return r
}

func deleteOrderRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	orderId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	helpers.Catch(err)

	return deleteOrderRequest{
		OrderId: orderId,
	}, nil
}

func deleteOrderDetailRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	orderDetailId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	helpers.Catch(err)

	return deleteOrderDetailRequest{
		OrderDetailId: orderDetailId,
	}, nil
}

func updateOrderRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := addOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helpers.Catch(err)

	return request, nil
}

func addOrderRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := addOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func getOrdersRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getOrdersRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func getOrderByIdRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	orderId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	helpers.Catch(err)

	return getOrderByIdRequest{
		OrderId: orderId,
	}, nil
}
