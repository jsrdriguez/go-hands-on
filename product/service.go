package product

type Service interface {
	GetProductById(param *getProductByIDRequest) (*Product, error)
	GetProducts(param *getProductRequest) (*ProductList, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetProductById(param *getProductByIDRequest) (*Product, error) {
	return s.repo.GetProductById(param.ProductID)
}

func (s *service) GetProducts(param *getProductRequest) (*ProductList, error) {
	products, err := s.repo.GetProducts(param)
	if err != nil {
		panic(err)
	}

	totalProducts, err := s.repo.GetTotalProducts()
	if err != nil {
		panic(err)
	}

	return &ProductList{Data: products, TotalRecords: totalProducts}, nil
}
