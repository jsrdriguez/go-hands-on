package order

type OrderItem struct {
	ID         int64              `json:"id"`
	CustomerId int                `json:"customerId"`
	OrderDate  string             `json:"orderDate"`
	StatusId   string             `json:"statusId"`
	StatusName string             `json:"statusName"`
	Customer   string             `json:"customer"`
	Company    string             `json:"company"`
	Address    string             `json:"address"`
	Phone      string             `json:"phone"`
	City       string             `json:"city"`
	Data       []*OrderDetailItem `json:"data"`
}

type OrderDetailItem struct {
	ID          int     `json:"id"`
	OrderId     int     `json:"orderId"`
	ProductId   int     `json:"productId"`
	Quantity    float64 `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
	ProductName string  `json:"productName"`
}

type OrderList struct {
	Data         []*OrderItem `json:"data"`
	TotalRecords int64        `json:"totalRecords"`
}
