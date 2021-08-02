package dao

import (
	"database/sql"
	"fmt"
	"shop/config"
	"shop/entity"
	"shop/tool"
)

func GetLenTotalProducts() (status int, len int) {
	serverConf := config.ServerConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		serverConf.SqlUser, serverConf.SqlPassword, serverConf.SqlSchema))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT productid FROM product")
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

func GetProductInfoList() (status int, resultList []entity.TProductItemInfo) {
	serverConf := config.ServerConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		serverConf.SqlUser, serverConf.SqlPassword, serverConf.SqlSchema))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT productname, price, totalnumber, seller FROM product WHERE onsale=?", tool.ProductOnSale)
	if err != nil {
		status = tool.ResponseFailed
		return
	}

	for rows.Next() {
		var info entity.TProductItemInfo

		err = rows.Scan(&info.ProductName, &info.Price, &info.TotalNumber, &info.Seller)
		if err != nil {
			status = tool.ResponseFailed
			return
		}

		resultList = append(resultList, info)
	}

	status = tool.ResponseOK
	return
}

func NewProduct(productName string, price int, totalNumber int, seller string) (status int) {
	serverConf := config.ServerConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		serverConf.SqlUser, serverConf.SqlPassword, serverConf.SqlSchema))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	status, len := GetLenTotalProducts()
	if status != tool.ResponseOK {
		return
	}
	newProductId := len + 1

	_, err = db.Exec(
		"INSERT INTO product (productid, productname, price, totalnumber, seller, onsale) VALUES(?, ?, ?, ?, ?, 1)",
		newProductId,
		productName,
		price,
		totalNumber,
		seller,
	)

	if err != nil {
		status = tool.ResponseFailed
	} else {
		status = tool.ResponseOK
	}

	return
}

func RemoveProductByProductId(productId int, seller string) (status int) {
	serverConf := config.ServerConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		serverConf.SqlUser, serverConf.SqlPassword, serverConf.SqlSchema))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE product SET onsale=-1 WHERE productid=? AND seller=?")
	if err != nil {
		status = tool.ResponseFailed
		return
	}

	_, err = stmt.Exec(productId, seller)
	if err != nil {
		status = tool.ResponseFailed
		return
	}

	status = tool.ResponseOK
	return
}
