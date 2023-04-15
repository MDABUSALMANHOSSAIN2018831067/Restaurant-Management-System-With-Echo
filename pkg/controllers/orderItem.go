package controllers

import (
	"fmt"
	"net/http"
	"restaurant-management/pkg/consts"
	"restaurant-management/pkg/models"
	"restaurant-management/pkg/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateOrderItem(c echo.Context) error {
	reqOderItem := &models.OrderItem{}
	if err := c.Bind(reqOderItem); err != nil {
		return c.JSON(http.StatusBadRequest, consts.InvalidInput)
	}
	if err := repositories.CreateOrderItem(reqOderItem); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "OrderItem was created success")
}

func GetOrderItems(c echo.Context) error {
	tempID := c.QueryParam("_id")
	ID, err := strconv.ParseInt(tempID, 0, 0)
	if err != nil && tempID != "" {
		return c.JSON(http.StatusBadRequest, consts.InvalidID)
	}
	orderItem, err := repositories.GetOrderItems(uint(ID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, orderItem)
}

func DeleteOrderItem(e echo.Context) error {
	tempID := e.Param("_id")
	ID, err := strconv.ParseUint(tempID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}
	_, err = repositories.GetOrderItems(uint(ID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := repositories.DeleteOrderItem(uint(ID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Order item was deleted success")
}

func UpdateOrderItem(c echo.Context) error {
	var item = &models.OrderItem{}
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id := c.Param("_id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, consts.InvalidID)
	}
	oldres, err := repositories.GetOrderItems(uint(ID))
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	item.ID = uint(ID)

	checkedItem := UpdateOrderItemField(item, oldres)

	res, err := repositories.UpdateOrderItem(checkedItem)
	if err != nil || res == nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "update was success")
}

func UpdateOrderItemField(item *models.OrderItem, old_Item []models.OrderItem) *models.OrderItem {
	if item.Quantity == 0 {
		item.Quantity = old_Item[0].Quantity
	}
	if item.FoodID == 0 {
		item.FoodID = old_Item[0].FoodID
	}
	return item
}
