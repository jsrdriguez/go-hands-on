package product

import "github.com/jsrdriguez/go-hands-on/helpers"

type Service interface {
	GetProductById(param *getProductByIDRequest) (*Product, error)
	GetProducts(param *getProductRequest) (*ProductList, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *updateProductRequest) (int64, error)
	DeleteProduct(params *deleteProductRequest) (int64, error)
	GetBetSellers() (*ProductTopResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetBetSellers() (*ProductTopResponse, error) {
	products, err := s.repo.GetBetSellers()
	helpers.Catch(err)

	totalVentas, err := s.repo.GetTotalVentas()
	helpers.Catch(err)

	return &ProductTopResponse{
		TotalVentas: totalVentas,
		Data:        products,
	}, nil
}

func (s *service) DeleteProduct(params *deleteProductRequest) (int64, error) {
	return s.repo.DeleteProduct(params)
}

func (s *service) UpdateProduct(params *updateProductRequest) (int64, error) {
	return s.repo.UpdateProduct(params)
}

func (s *service) InsertProduct(params *getAddProductRequest) (int64, error) {
	return s.repo.InsertProduct(params)
}

func (s *service) GetProductById(param *getProductByIDRequest) (*Product, error) {
	return s.repo.GetProductById(param.ProductID)
}

func (s *service) GetProducts(param *getProductRequest) (*ProductList, error) {
	products, err := s.repo.GetProducts(param)
	helpers.Catch(err)

	totalProducts, err := s.repo.GetTotalProducts()
	helpers.Catch(err)

	return &ProductList{Data: products, TotalRecords: totalProducts}, nil
}
