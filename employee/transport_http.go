package employee

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jsrdriguez/go-hands-on/helpers"

	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	r.Method(http.MethodPost, "/paginated", kithttp.NewServer(
		makeGetEmployeesEndpoint(s),
		getEmployeesRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	r.Method(http.MethodGet, "/{id}", kithttp.NewServer(
		makeGetEmployeeIdEndpoint(s),
		getEmployeeIdRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	r.Method(http.MethodDelete, "/{id}", kithttp.NewServer(
		makeDeleteEmployee(s),
		getDeleteEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	r.Method(http.MethodGet, "/best", kithttp.NewServer(
		makeGetBestEmployee(s),
		getBestEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	r.Method(http.MethodPost, "/", kithttp.NewServer(
		makeInsertEmployee(s),
		getAddEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	r.Method(http.MethodPut, "/", kithttp.NewServer(
		makeUpdateEmployee(s),
		getUpdateEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse,
	))

	return r
}

func getDeleteEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	return deleteEmployeesRequest{
		EmployeeId: id,
	}, nil
}

func getUpdateEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateEmployeesRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	helpers.Catch(err)

	return request, nil
}

func getAddEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addEmployeesRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	helpers.Catch(err)

	return request, nil
}

func getBestEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return getEmployeesBestRequest{}, nil
}

func getEmployeeIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	return getEmployeesByIdRequest{
		EmployeeId: id,
	}, nil
}

func getEmployeesRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeesRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	helpers.Catch(err)

	return request, nil
}
