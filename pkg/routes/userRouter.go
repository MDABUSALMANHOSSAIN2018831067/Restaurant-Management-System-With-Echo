package routes

import (
	"restaurant-management/pkg/controllers"
	"restaurant-management/pkg/middleware"

	//"restaurant-management/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, userController *controllers.UserController) {
	user := e.Group("/restaurantmanagemensystem")

	user.POST("/registration", userController.Registration)
	user.POST("/login", userController.Login)
	user.GET("/users", userController.GetUsers)
	user.DELETE("/users/:_id", userController.DeleteUser, middleware.Authentication, middleware.IsAdmin)
	user.PUT("/users/:_id", userController.UpdateUser, middleware.Authentication)

}

// package routes

// import (
// 	"restaurant-management/pkg/controllers"

// 	"github.com/labstack/echo/v4"
// )

// func User(e *echo.Echo) {
// 	sub := e.Group("/user")
// 	sub.POST("/registration", controllers.Registration)
// 	sub.POST("/login", controllers.Login)
// 	sub.GET("/users", controllers.GetAllUsers)
// 	sub.GET("/:id", controllers.GetAUsers)
// 	//sub.PUT("/:id", controllers.UpdateUser)
// 	sub.DELETE("/:id", controllers.DeleteUser)

// }
