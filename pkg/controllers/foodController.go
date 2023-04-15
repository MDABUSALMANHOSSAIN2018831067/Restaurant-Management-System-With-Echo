package controllers

import (
	"fmt"
	"net/http"
	"restaurant-management/pkg/consts"
	"restaurant-management/pkg/domain"
	"restaurant-management/pkg/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FoodController struct {
	foodService domain.FoodServiceInterface
}

func SetFoodService(foodService *domain.FoodServiceInterface) *FoodController {
	return &FoodController{
		foodService: *foodService,
	}
}

func (foodController *FoodController) CreateFood(e echo.Context) error {
	reqFood := &models.Food{}
	if err := e.Bind(reqFood); err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidInput)
	}

	if err := foodController.foodService.CreateFoodService(reqFood); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, "Food was created success")
}

func (foodController *FoodController) GetFoods(e echo.Context) error {
	tempFoodID := e.QueryParam("_id")
	foodID, err := strconv.ParseInt(tempFoodID, 0, 0)
	if err != nil && tempFoodID != "" {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}

	food, err := foodController.foodService.GetFoodService(uint(foodID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, food)
}

func (foodController *FoodController) UpdateFood(c echo.Context) error {
	var food = &models.Food{}
	if err := c.Bind(food); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id := c.Param("_id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, consts.InvalidID)
	}
	oldres, err := foodController.foodService.GetFoodService(uint(ID))
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	food.ID = uint(ID)

	checkedFood := UpdateFoodField(food, oldres)

	res, err := foodController.foodService.UpdateFoodService(checkedFood)
	if err != nil || res == nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "update was success")
}

func UpdateFoodField(food *models.Food, old_Food []models.Food) *models.Food {
	if food.Name == "" {
		food.Name = old_Food[0].Name
	}
	if food.Price == 0 {
		food.Price = old_Food[0].Price
	}
	return food
}

func (foodController *FoodController) DeleteFood(e echo.Context) error {
	tempID := e.Param("_id")
	ID, err := strconv.ParseUint(tempID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}

	_, err = foodController.foodService.GetFoodService(uint(ID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if err := foodController.foodService.DeleteFoodService(uint(ID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, "Food was deleted success")
}
