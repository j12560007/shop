package entity

type TOrderItemInfo struct {
	ProductName string `json:"productname"`
	Price       int    `json:"price"`
	Number      int    `json:"number"`
	Seller      string `json:"seller"`
}

type TOrderNewRequest struct {
	ProductId int    `json:"productid"`
	Number    int    `json:"number"`
	BuyerName string `json:"buyername"`
}

type TOrderNewResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type TOrderListMineRequest struct {
	BuyerName string `json:"buyername"`
}

type TOrderListMineResponse struct {
	Status    int              `json:"status"`
	Message   string           `json:"message"`
	OrderList []TOrderItemInfo `json:"orderlist"`
}
