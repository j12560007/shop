package dao

import (
	"database/sql"
	"fmt"
	"shop/config"
	"shop/entity"
	"shop/tool"
)

func GetLenTotalOrders() (status int, len int) {
	serverConf := config.ServerConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		serverConf.SqlUser, serverConf.SqlPassword, serverConf.SqlSchema))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT orderid FROM `order`")
	if err != nil {
		status = tool.ResponseFailed
		return
	}

	len = 0

	for rows.Next() {
		len++
	}

	status = tool.ResponseOK
	return
}

func CreateOrder(productId int, number int, buyerName string) (status int) {
	serverConf := config.ServerConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		serverConf.SqlUser, serverConf.SqlPassword, serverConf.SqlSchema))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	status, len := GetLenTotalOrders()
	if status != tool.ResponseOK {
		return
	}
	newOrderId := len + 1

	_, err = db.Exec(
		"INSERT INTO `order` (orderid, productid, number, buyername, status) VALUES(?, ?, ?, ?, 1)",
		newOrderId,
		productId,
		number,
		buyerName,
	)

	if err != nil {
		status = tool.ResponseFailed
	} else {
		status = tool.ResponseOK
	}
	return
}

func GetOrderInfoListByBuyerName(buyerName string) (status int, resultList []entity.TOrderItemInfo) {
	serverConf := config.ServerConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		serverConf.SqlUser, serverConf.SqlPassword, serverConf.SqlSchema))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT p.productname, p.price, o.number, p.seller "+
		"FROM product AS p INNER JOIN `order` as o "+
		"WHERE p.productid = o.productid AND o.buyername = ?", buyerName)
	if err != nil {
		status = tool.ResponseFailed
		return
	}

	for rows.Next() {
		var info entity.TOrderItemInfo

		err = rows.Scan(&info.ProductName, &info.Price, &info.Number, &info.Seller)
		if err != nil {
			status = tool.ResponseFailed
			return
		}

		resultList = append(resultList, info)
	}

	status = tool.ResponseOK
	return
}
