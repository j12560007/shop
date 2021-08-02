package entity

type TProductItemInfo struct {
	ProductName string `json:"productname"`
	Price       int    `json:"price"`
	TotalNumber int    `json:"totalnumber"`
	Seller      string `json:"seller"`
}

type TProductListAllResponse struct {
	Status      int                `json:"status"`
	Message     string             `json:"message"`
	ProductList []TProductItemInfo `json:"productlist"`
}

type TProductNewRequest struct {
	Seller      string `json:"seller"`
	ProductName string `json:"productname"`
	Price       int    `json:"price"`
	TotalNumber int    `json:"totalnumber"`
}

type TProductNewResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type TProductRemoveRequest struct {
	ProductId int    `json:"productid"`
	Seller    string `json:"seller"`
}

type TProductRemoveResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
