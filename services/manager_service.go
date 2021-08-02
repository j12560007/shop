package services

import (
	"github.com/gin-gonic/gin"
	"shop/entity"
)

func ManagerUploadProductService(ctx *gin.Context) {
	var request entity.TUserInfoRequest
	if err := ctx.BindJSON(&request); err != nil {
		panic(err)
	}

}

func ManagerOrderListService(ctx *gin.Context) {
	
}