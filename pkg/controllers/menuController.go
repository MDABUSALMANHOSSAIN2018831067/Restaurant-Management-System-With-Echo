package controllers

import (
	"fmt"
	"net/http"
	"restaurant-management/pkg/consts"
	"restaurant-management/pkg/domain"
	"restaurant-management/pkg/models"
	"restaurant-management/pkg/types"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MenuController struct {
	menuService domain.MenuServiceInterface
}

func SetMenuService(menuService *domain.MenuServiceInterface) *MenuController {
	return &MenuController{
		menuService: *menuService,
	}
}
func (menuController *MenuController) CreateMenu(c echo.Context) error {

	reqMenu := &models.Menu{}
	if err := c.Bind(reqMenu); err != nil {
		return c.JSON(http.StatusBadRequest, consts.InvalidInput)
	}
	if err := menuController.menuService.CreateMenuService(reqMenu); err != nil {
		return c.JSON(http.StatusInternalServerError, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}

	return c.JSON(http.StatusCreated, "menu was created success")

}

func (menuController *MenuController) GetMenus(c echo.Context) error {
	tempID := c.QueryParam("_id")
	ID, err := strconv.ParseInt(tempID, 0, 0)
	if err != nil && tempID != "" {
		return c.JSON(http.StatusBadRequest, consts.InvalidID)
	}

	menu, err := menuController.menuService.GetMenuService(uint(ID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}

	return c.JSON(http.StatusOK, menu)
}

func (menuController *MenuController) DeleteMenu(e echo.Context) error {
	tempID := e.Param("_id")
	ID, err := strconv.ParseUint(tempID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}
	_, err = menuController.menuService.GetMenuService(uint(ID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}
	if err := menuController.menuService.DeleteMenuService(uint(ID)); err != nil {
		return e.JSON(http.StatusInternalServerError, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}
	return e.JSON(http.StatusOK, "Menu was deleted success")
}

func (menuController *MenuController) UpdateManu(c echo.Context) error {
	var menu = &models.Menu{}
	if err := c.Bind(menu); err != nil {
		return c.JSON(http.StatusBadRequest, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}
	id := c.Param("_id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, consts.InvalidID)
	}
	oldres, err := menuController.menuService.GetMenuService(uint(ID))
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusOK, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}
	menu.ID = uint(ID)

	checkedMenu := UpdateMenuField(menu, oldres)

	res, err := menuController.menuService.UpdateManuService(checkedMenu)
	if err != nil || res == nil {
		return c.JSON(http.StatusInternalServerError, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}
	return c.JSON(http.StatusOK, "update was success")
}

func UpdateMenuField(menu *models.Menu, old_Menu []models.Menu) *models.Menu {
	if menu.Category == "" {
		menu.Category = old_Menu[0].Category
	}
	if menu.FoodID == 0 {
		menu.FoodID = old_Menu[0].FoodID
	}
	return menu
}
