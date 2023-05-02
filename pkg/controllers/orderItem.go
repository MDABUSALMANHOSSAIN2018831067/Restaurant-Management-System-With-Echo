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

type OrderItemController struct {
	orderItemService domain.OrderItemServiceInterface
}

func SetOrderItemService(orderItemService *domain.OrderItemServiceInterface) *OrderItemController {
	return &OrderItemController{
		orderItemService: *orderItemService,
	}
}
func (orderItemController *OrderItemController) CreateOrderItem(c echo.Context) error {
	reqOderItem := &models.OrderItem{}
	if err := c.Bind(reqOderItem); err != nil {
		return c.JSON(http.StatusBadRequest, consts.InvalidInput)
	}
	if err := orderItemController.orderItemService.CreateOrderItemService(reqOderItem); err != nil {
		return c.JSON(http.StatusInternalServerError, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}

	return c.JSON(http.StatusCreated, "OrderItem was created success")
}

func (orderItemController *OrderItemController) GetOrderItems(c echo.Context) error {
	tempID := c.QueryParam("_id")
	ID, err := strconv.ParseInt(tempID, 0, 0)
	if err != nil && tempID != "" {
		return c.JSON(http.StatusBadRequest, consts.InvalidID)
	}
	orderItem, err := orderItemController.orderItemService.GetOrderItemService(uint(ID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}

	return c.JSON(http.StatusOK, orderItem)
}

func (orderItemController *OrderItemController) DeleteOrderItem(e echo.Context) error {
	tempID := e.Param("_id")
	ID, err := strconv.ParseUint(tempID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}
	_, err = orderItemController.orderItemService.GetOrderItemService(uint(ID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}
	if err := orderItemController.orderItemService.DeleteOrderItemService(uint(ID)); err != nil {
		return e.JSON(http.StatusInternalServerError, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}
	return e.JSON(http.StatusOK, "Order item was deleted success")
}

func (orderItemController *OrderItemController) UpdateOrderItem(c echo.Context) error {
	var item = &models.OrderItem{}
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, consts.InvalidInput)
	}
	id := c.Param("_id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, consts.InvalidID)
	}
	oldres, err := orderItemController.orderItemService.GetOrderItemService(uint(ID))
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusOK, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}
	item.ID = uint(ID)

	checkedItem := UpdateOrderItemField(item, oldres)

	res, err := orderItemController.orderItemService.UpdateOrderItemService(checkedItem)
	if err != nil || res == nil {
		return c.JSON(http.StatusInternalServerError, &types.CustomError{
			Message: err.Error(),
			Err:     err,
		})
	}
	return c.JSON(http.StatusOK, "update was success")
}

func UpdateOrderItemField(item *models.OrderItem, old_Item []models.OrderItem) *models.OrderItem {
	if item.Quantity == 0 {
		item.Quantity = old_Item[0].Quantity
	}
	if item.MenuID == 0 {
		item.MenuID = old_Item[0].MenuID
	}
	return item
}
