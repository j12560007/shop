package services

import (
	"net/http"
	"shop/dao"
	"shop/entity"
	"shop/tool"

	"github.com/gin-gonic/gin"
)

func OrderNewService(ctx *gin.Context) {
	var request entity.TOrderNewRequest
	if err := ctx.BindJSON(&request); err != nil {
		panic(err)
	}

	status := dao.CreateOrder(request.ProductId, request.Number, request.BuyerName)

	var response entity.TOrderNewResponse
	response.Status = status

	if response.Status != tool.ResponseOK {
		response.Message = "訂單建立失敗！"
	} else {
		response.Message = "訂單建立完成！"
	}

	ctx.JSON(http.StatusOK, response)
}

func OrderListMineService(ctx *gin.Context) {
	var request entity.TOrderListMineRequest
	if err := ctx.BindJSON(&request); err != nil {
		panic(err)
	}

	status, resultList := dao.GetOrderInfoListByBuyerName(request.BuyerName)

	var response entity.TOrderListMineResponse
	response.Status = status
	response.OrderList = resultList

	if response.Status != tool.ResponseOK {
		response.Message = "訂單列表取得失敗！"
	} else {
		response.Message = "訂單列表取得完成！"
	}

	ctx.JSON(http.StatusOK, response)
}
