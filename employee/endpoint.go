package employee

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/jsrdriguez/go-hands-on/helpers"
)

type getEmployeesRequest struct {
	Limit  int
	Offset int
}

type getEmployeesByIdRequest struct {
	EmployeeId int
}

type getEmployeesBestRequest struct {
}

func makeGetBestEmployee(s Service) endpoint.Endpoint {

	getBestEmployee := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesBestRequest)

		result, err := s.GetBestEmployee(&req)
		helpers.Catch(err)

		return result, nil
	}

	return getBestEmployee
}

func makeGetEmployeeIdEndpoint(s Service) endpoint.Endpoint {

	getEmployeeId := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesByIdRequest)

		result, err := s.GetEmployeeById(&req)
		helpers.Catch(err)

		return result, nil
	}

	return getEmployeeId
}

func makeGetEmployeesEndpoint(s Service) endpoint.Endpoint {

	getEmployeesEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesRequest)

		result, err := s.GetEmployees(&req)
		helpers.Catch(err)

		return result, nil
	}

	return getEmployeesEndpoint
}
