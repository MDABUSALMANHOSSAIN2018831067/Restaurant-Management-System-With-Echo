package containers

import (
	"fmt"
	"log"
	"restaurant-management/pkg/config"
	"restaurant-management/pkg/connection"
	"restaurant-management/pkg/controllers"
	"restaurant-management/pkg/repositories"
	"restaurant-management/pkg/routes"
	"restaurant-management/pkg/services"
	unittesting "restaurant-management/pkg/unitTesting"

	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {
	config.SetConfig()
	db := connection.GetDB()
	foodrepo := repositories.FoodDBInterface(db)
	foodService := services.FoodServiceInstance(foodrepo)
	foodController := controllers.SetFoodService(&foodService)

	userRepo := repositories.UserDBInterface(db)
	userService := services.UserServiceInstance(userRepo)
	userController := controllers.SetUserService(&userService)
	unittesting.SetUserTestingService(&userService)

	menuRepo := repositories.MenuDBInterface(db)
	menuservice := services.MenuServiceInstance(menuRepo)
	menuController := controllers.SetMenuService(&menuservice)

	orderItemRepo := repositories.OrderItemDBInterface(db)
	orderItemService := services.OrderItemServiceInstance(orderItemRepo)
	orderItemController := controllers.SetOrderItemService(&orderItemService)

	routes.UserRoutes(e, userController)
	routes.FoodRoutes(e, foodController)
	routes.MenuRoutes(e, menuController)
	routes.OrderItemRoutes(e, orderItemController)
	// routes.TableRoutes(e)
	// routes.OrderRoutes(e)
	// routes.InvoiceRoutes(e)

	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))

}
