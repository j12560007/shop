package routes

import (
	"shop/services"

	"github.com/gin-gonic/gin"
)

func SetRouteer() *gin.Engine {
	router := gin.Default()

	shop := router.Group("/shop")
	{
		manager := shop.Group("/manager")
		{
			manager.POST("/uploadproduct", services.ManagerUploadProductService)
			manager.POST("/listorder", services.ManagerOrderListService)
		}

		user := shop.Group("/user")
		{
			user.POST("/register", services.UserRegisterService)
			user.POST("/login", services.UserLoginService)
			user.GET("/info", services.UserInfoService)
			user.POST("/editinfo", services.UserEditInfoService)
			user.POST("/editpassword", services.UserEditPasswordService)
		}

		product := shop.Group("/product")
		{
			product.GET("/listall", services.ProductListAllService)
			product.POST("/new", services.ProductNewService)
			product.POST("/remove", services.ProductRemoveService)
		}

		order := shop.Group("/order")
		{
			order.POST("/new", services.OrderNewService)
			order.GET("/listmine", services.OrderListMineService)
		}

		// test
		purchase := shop.Group("/purchase")
		{
			purchase.GET("/listproduct", services.PurchaseListProductService)
			// purchase.GET("/cart/list", services.PurchaseListCartService)
			// purchase.POST("/cart/add", services.PurchaseAddToCartService)
			// purchase.POST("/cart/remove", services.PurchaseRemoveFromCartService)
			// purchase.POST("/cart/checkout", services.PurchaseCheckoutFromCartService)
		}
	}

	return router
}
