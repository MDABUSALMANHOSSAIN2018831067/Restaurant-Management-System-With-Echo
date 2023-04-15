package routes

import (
	"restaurant-management/pkg/controllers"
	"restaurant-management/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func MenuRoutes(e *echo.Echo, menuController *controllers.MenuController) {
	menu := e.Group("/restaurantmanagemensystem", middleware.Authentication, middleware.IsAdmin)

	menu.POST("/menus", menuController.CreateMenu)
	menu.GET("/menus", menuController.GetMenus)
	menu.DELETE("/menus/:_id", menuController.DeleteMenu)
	menu.PUT("/menus/:_id", menuController.UpdateManu)
}
