package services

import (
	"errors"
	"restaurant-management/pkg/domain"
	"restaurant-management/pkg/models"
)

type OrderItemService struct {
	repo domain.OrderItemRepoInterface
}

func OrderItemServiceInstance(orderItemRepo domain.OrderItemRepoInterface) domain.OrderItemServiceInterface {
	return &OrderItemService{
		repo: orderItemRepo,
	}
}

func (service *OrderItemService) CreateOrderItemService(orderItem *models.OrderItem) error {
	err := service.repo.CreateOrderItem(orderItem)
	return err
}

func (service *OrderItemService) GetOrderItemService(ID uint) ([]models.OrderItem, error) {
	items, err := service.repo.GetOrderItems(ID)
	return items, err
}

func (service *OrderItemService) DeleteOrderItemService(ID uint) error {
	if err := service.repo.DeleteOrderItem(ID); err != nil {
		return errors.New("order item deletion was unsuccessful")
	}
	return nil
}

func (service *OrderItemService) UpdateOrderItemService(item *models.OrderItem) (*models.OrderItem, error) {
	items, err := service.repo.UpdateOrderItem(item)
	if err != nil {
		return nil, errors.New("order item update was unsuccesful")
	}
	return items, nil
}
