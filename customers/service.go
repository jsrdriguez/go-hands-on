package customer

import "github.com/jsrdriguez/go-hands-on/helpers"

type Service interface {
	getCustomers(params *getCustomerRequest) (*CustomerList, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) getCustomers(params *getCustomerRequest) (*CustomerList, error) {
	customers, err := s.repo.getCustomers(params)
	helpers.Catch(err)

	totalRecords, err := s.repo.getTotalCustomers()
	helpers.Catch(err)

	return &CustomerList{
		Data:         customers,
		TotalRecords: totalRecords,
	}, nil
}
