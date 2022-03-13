package employee

import "github.com/jsrdriguez/go-hands-on/helpers"

type Service interface {
	GetEmployees(params *getEmployeesRequest) (*EmployeeList, error)
	GetEmployeeById(params *getEmployeesByIdRequest) (*Employee, error)
	GetBestEmployee(params *getEmployeesBestRequest) (*BestEmployee, error)
	AddEmployee(params *addEmployeesRequest) (int64, error)
	UpdateEmployee(params *updateEmployeesRequest) (int64, error)
	DeleteEmployee(params *deleteEmployeesRequest) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) DeleteEmployee(params *deleteEmployeesRequest) (int64, error) {
	return s.repo.DeleteEmployee(params)
}

func (s service) UpdateEmployee(params *updateEmployeesRequest) (int64, error) {
	return s.repo.UpdateEmployee(params)
}

func (s service) AddEmployee(params *addEmployeesRequest) (int64, error) {
	return s.repo.AddEmployee(params)
}

func (s service) GetBestEmployee(params *getEmployeesBestRequest) (*BestEmployee, error) {
	return s.repo.GetBestEmployee(params)
}

func (s service) GetEmployeeById(params *getEmployeesByIdRequest) (*Employee, error) {
	return s.repo.GetEmployeeById(params)
}

func (s service) GetEmployees(params *getEmployeesRequest) (*EmployeeList, error) {
	employess, err := s.repo.GetEmployees(params)
	helpers.Catch(err)

	totalEmployees, err := s.repo.GetTotalEmployees()
	helpers.Catch(err)

	return &EmployeeList{
		Data:         employess,
		TotalRecords: totalEmployees,
	}, nil
}
