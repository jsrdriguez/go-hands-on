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

// @Sumary Elimina un Empleado by Id
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path int true "Employee Id"
// @Success 200 {object} employee.deleteEmployeesRequest "ok"
// @Router /employees/{id} [delete]
func makeDeleteEmployee(s Service) endpoint.Endpoint {
	deleteEmployees := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteEmployeesRequest)

		result, err := s.DeleteEmployee(&req)
		helpers.Catch(err)

		return result, err
	}

	return deleteEmployees
}

// @Sumary Actualizar Empleados
// @Tags Employee
// @Accept json
// @Produce json
// @Param request body employee.updateEmployeesRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /employees/ [put]
func makeUpdateEmployee(s Service) endpoint.Endpoint {
	updateEmployee := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateEmployeesRequest)

		result, err := s.UpdateEmployee(&req)
		helpers.Catch(err)

		return result, nil
	}

	return updateEmployee
}

// @Sumary Insertar Empleados
// @Tags Employee
// @Accept json
// @Produce json
// @Param request body employee.addEmployeesRequest true "User Data"
// @Success 200 {integer} int "ok"
// @Router /employees/ [post]
func makeInsertEmployee(s Service) endpoint.Endpoint {
	insertEmployee := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addEmployeesRequest)

		result, err := s.AddEmployee(&req)
		helpers.Catch(err)

		return result, nil
	}

	return insertEmployee
}

// @Sumary Mejores Empleado
// @Tags Employee
// @Accept json
// @Produce json
// @Success 200 {object} employee.BestEmployee "ok"
// @Router /employees/best [get]
func makeGetBestEmployee(s Service) endpoint.Endpoint {

	getBestEmployee := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesBestRequest)

		result, err := s.GetBestEmployee(&req)
		helpers.Catch(err)

		return result, nil
	}

	return getBestEmployee
}

// @Sumary Empleado by Id
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path int true "Employee Id"
// @Success 200 {object} employee.Employee "ok"
// @Router /employees/{id} [get]
func makeGetEmployeeIdEndpoint(s Service) endpoint.Endpoint {

	getEmployeeId := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesByIdRequest)

		result, err := s.GetEmployeeById(&req)
		helpers.Catch(err)

		return result, nil
	}

	return getEmployeeId
}

// @Sumary Lista de Empleados
// @Tags Employee
// @Accept json
// @Produce json
// @Param request body employee.getEmployeesRequest true "User Data"
// @Success 200 {object} employee.EmployeeList "ok"
// @Router /employees/paginated [post]
func makeGetEmployeesEndpoint(s Service) endpoint.Endpoint {

	getEmployeesEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesRequest)

		result, err := s.GetEmployees(&req)
		helpers.Catch(err)

		return result, nil
	}

	return getEmployeesEndpoint
}
