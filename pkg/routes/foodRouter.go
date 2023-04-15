package routes

import (
	"restaurant-management/pkg/controllers"
	"restaurant-management/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func FoodRoutes(e *echo.Echo, foodController *controllers.FoodController) {
	food := e.Group("/restaurantmanagemensystem", middleware.Authentication, middleware.IsAdmin)

	food.POST("/food", foodController.CreateFood)
	food.GET("/food", foodController.GetFoods)
	food.PUT("/food/:_id", foodController.UpdateFood)
	food.DELETE("/food/:_id", foodController.DeleteFood)
}
