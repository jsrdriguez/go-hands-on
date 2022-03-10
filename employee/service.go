package employee

import "github.com/jsrdriguez/go-hands-on/helpers"

type Service interface {
	GetEmployees(params *getEmployeesRequest) (*EmployeeList, error)
	GetEmployeeById(params *getEmployeesByIdRequest) (*Employee, error)
	GetBestEmployee(params *getEmployeesBestRequest) (*BestEmployee, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
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
