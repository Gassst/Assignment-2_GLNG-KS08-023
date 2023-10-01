package routers

import (
	"Assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/orders", controllers.GetAllOrders)
	router.POST("/orders", controllers.CreateOrder)
	router.PATCH("/orders/:orderId", controllers.UpdateOrder)
	router.PUT("/orders/:orderId/items/:itemId", controllers.UpdateItem)
	router.DELETE("/orders/:orderId", controllers.DeleteOrder)

	return router
}
