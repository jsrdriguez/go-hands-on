package customer

type Customer struct {
	ID            int    `json:"id"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Address       string `json:"address"`
	BusinessPhone string `json:"businessPhone"`
	City          string `json:"city"`
	Company       string `json:"company"`
}

type CustomerList struct {
	Data         []*Customer `json:"data"`
	TotalRecords int64       `json:"totalRecords"`
}
