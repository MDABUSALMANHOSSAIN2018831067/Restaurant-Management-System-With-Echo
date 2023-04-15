package domain

import "restaurant-management/pkg/models"

type FoodRepoInterface interface {
	CreateFood(food *models.Food) error
	GetFoods(foodID uint) ([]models.Food, error)
	UpdateFood(food *models.Food) (*models.Food, error)
	DeleteFood(ID uint) error
}

type FoodServiceInterface interface {
	CreateFoodService(food *models.Food) error
	GetFoodService(ID uint) ([]models.Food, error)
	UpdateFoodService(food *models.Food) (*models.Food, error)
	DeleteFoodService(ID uint) error
}

type UserRepoInterface interface {
	Registration(user *models.User) error
	Login(email string) (*models.User, error)
	GetUsers(ID uint) ([]models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(ID uint) error
}

type UserServiceInterface interface {
	RegistrationService(user *models.User) error
	LoginService(email string) (*models.User, error)
	GetUserService(ID uint) ([]models.User, error)
	UpdateUserService(user *models.User) (*models.User, error)
	DeleteUserService(ID uint) error
}

type MenuRepoInterface interface{
	CreateMenu(menu *models.Menu) error
	GetMenus(ID uint) ([]models.Menu, error)
	DeleteMenu(ID uint) error
	UpdateManu(menu *models.Menu) (*models.Menu, error)
}

type MenuServiceInterface interface{
	CreateMenuService(menu *models.Menu) error
	GetMenuService(ID uint) ([]models.Menu, error)
	DeleteMenuService(ID uint) error
	UpdateManuService(menu *models.Menu) (*models.Menu, error)
}