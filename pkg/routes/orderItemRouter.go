package routes

import (
	"restaurant-management/pkg/controllers"
	"restaurant-management/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func OrderItemRoutes(e *echo.Echo) {
	orderItem := e.Group("/restaurantmanagemensystem")
	orderItem.POST("/orderItems", controllers.CreateOrderItem, middleware.Authentication, middleware.IsAdmin)
	orderItem.GET("/orderItems", controllers.GetOrderItems)
	orderItem.DELETE("/orderItems/:_id", controllers.DeleteOrderItem, middleware.Authentication, middleware.IsAdmin)
	orderItem.PUT("/orderItems/:_id", controllers.UpdateOrderItem, middleware.Authentication, middleware.IsAdmin)
}
