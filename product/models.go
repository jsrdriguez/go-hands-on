package product

type Product struct {
	Id           int     `json:"id"`
	ProductCode  string  `json:"product_code"`
	ProductName  string  `json:"product_name"`
	Description  string  `json:"description"`
	StandardCost float64 `json:"standardCost"`
	ListPrice    float64 `json:"listPrice"`
	Category     string  `json:"category"`
}

type ProductList struct {
	Data         []*Product `json:"data"`
	TotalRecords int        `json:"totalRecords"`
}
