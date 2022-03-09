package product

type Service interface {
	GetProductById(param *getProductByIDRequest) (*Product, error)
	GetProducts(param *getProductRequest) (*ProductList, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *updateProductRequest) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
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
	if err != nil {
		panic(err)
	}

	totalProducts, err := s.repo.GetTotalProducts()
	if err != nil {
		panic(err)
	}

	return &ProductList{Data: products, TotalRecords: totalProducts}, nil
}
