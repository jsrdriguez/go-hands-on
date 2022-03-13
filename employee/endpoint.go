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

type addEmployeesRequest struct {
	Address       string
	BusinessPhone string
	Company       string
	EmailAddress  string
	FaxNumber     string
	FirstName     string
	HomePhone     string
	JobTitle      string
	LastName      string
	MobilePhone   string
}

type updateEmployeesRequest struct {
	ID            int64
	Address       string
	BusinessPhone string
	Company       string
	EmailAddress  string
	FaxNumber     string
	FirstName     string
	HomePhone     string
	JobTitle      string
	LastName      string
	MobilePhone   string
}

type deleteEmployeesRequest struct {
	EmployeeId int
}

func makeDeleteEmployee(s Service) endpoint.Endpoint {
	deleteEmployees := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteEmployeesRequest)

		result, err := s.DeleteEmployee(&req)
		helpers.Catch(err)

		return result, err
	}

	return deleteEmployees
}

func makeUpdateEmployee(s Service) endpoint.Endpoint {
	updateEmployee := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateEmployeesRequest)

		result, err := s.UpdateEmployee(&req)
		helpers.Catch(err)

		return result, nil
	}

	return updateEmployee
}

func makeInsertEmployee(s Service) endpoint.Endpoint {
	insertEmployee := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addEmployeesRequest)

		result, err := s.AddEmployee(&req)
		helpers.Catch(err)

		return result, nil
	}

	return insertEmployee
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
