package routes

import (
	"restaurant-management/pkg/controllers"
	"restaurant-management/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func OrderItemRoutes(e *echo.Echo, orderItemController *controllers.OrderItemController) {
	orderItem := e.Group("/restaurantmanagemensystem")
	orderItem.POST("/orderItems", orderItemController.CreateOrderItem, middleware.Authentication, middleware.IsAdmin)
	orderItem.GET("/orderItems", orderItemController.GetOrderItems)
	orderItem.DELETE("/orderItems/:_id", orderItemController.DeleteOrderItem, middleware.Authentication, middleware.IsAdmin)
	orderItem.PUT("/orderItems/:_id", orderItemController.UpdateOrderItem, middleware.Authentication, middleware.IsAdmin)
}
