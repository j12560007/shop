package services

import (
	"net/http"
	"shop/dao"
	"shop/entity"
	"shop/tool"

	"github.com/gin-gonic/gin"
)

func ProductListAllService(ctx *gin.Context) {

	status, resultList := dao.GetProductInfoList()

	var response entity.TProductListAllResponse
	response.Status = status
	response.ProductList = resultList

	if response.Status != tool.ResponseOK {
		response.Message = "商品列表取得失敗！"
	} else {
		response.Message = "商品列表取得完成！"
	}

	ctx.JSON(http.StatusOK, response)
}

func ProductNewService(ctx *gin.Context) {
	var request entity.TProductNewRequest
	if err := ctx.BindJSON(&request); err != nil {
		panic(err)
	}

	status := dao.NewProduct(request.ProductName, request.Price, request.TotalNumber, request.Seller)

	var response entity.TProductNewResponse
	response.Status = status

	if response.Status != tool.ResponseOK {
		response.Message = "商品新增失敗！"
	} else {
		response.Message = "商品新增完成！"
	}

	ctx.JSON(http.StatusOK, response)
}

func ProductRemoveService(ctx *gin.Context) {
	var request entity.TProductRemoveRequest
	if err := ctx.BindJSON(&request); err != nil {
		panic(err)
	}

	status := dao.RemoveProductByProductId(request.ProductId, request.Seller)

	var response entity.TProductRemoveResponse
	response.Status = status

	if response.Status != tool.ResponseOK {
		response.Message = "商品下架失敗！商品ID或是賣家不符合！"
	} else {
		response.Message = "商品下架完成！"
	}

	ctx.JSON(http.StatusOK, response)
}
