package repositories

import (
	"errors"
	"restaurant-management/pkg/domain"
	"restaurant-management/pkg/models"

	"gorm.io/gorm"
)

type OrderItemRepo struct {
	DB *gorm.DB
}

func OrderItemDBInterface(db *gorm.DB) domain.OrderItemRepoInterface {
	return &OrderItemRepo{
		DB: db,
	}

}
func (itemRepo *OrderItemRepo) CreateOrderItem(orderItem *models.OrderItem) error {
	if err := itemRepo.DB.Where("id = ?", orderItem.MenuID).First(&models.Menu{}).Error; err != nil {
		return err
	}
	if err := itemRepo.DB.Where("quantity = ? AND menu_id = ?", orderItem.Quantity, orderItem.MenuID).
		First(&orderItem).Error; err == nil {
		return errors.New("order item already exists")
	}
	if err := itemRepo.DB.Create(&orderItem).Error; err != nil {
		return err
	}
	return nil
}

func (itemRepo *OrderItemRepo) GetOrderItems(ID uint) ([]models.OrderItem, error) {
	var orderItem []models.OrderItem
	// if err := DB.Where("id = ?", orderItem[0].FoodID).First(&models.Food{}).Error; err != nil {
	// 	return nil, errors.New("food id not found")
	// }
	query := itemRepo.DB.Preload("Menu").Preload("Menu.Food")
	if ID != 0 {
		if err := query.Where("id = ?", ID).First(&orderItem).Error; err != nil {
			return nil, errors.New("orderitem id not found")
		}
	} else {
		if err := query.Find(&orderItem).Error; err != nil {
			return nil, err
		}
	}
	return orderItem, nil
}

func (itemRepo *OrderItemRepo) DeleteOrderItem(ID uint) error {
	var item models.OrderItem
	if err := itemRepo.DB.
		Where("id = ?", ID).Delete(&item).
		Error; err != nil {
		return err
	}
	return nil
}

func (itemRepo *OrderItemRepo) UpdateOrderItem(item *models.OrderItem) (*models.OrderItem, error) {
	if err := itemRepo.DB.Where("id = ?", item.MenuID).First(&models.Menu{}).Error; err != nil {
		return nil, errors.New("menu id not found")
	}
	if err := itemRepo.DB.Model(&item).Save(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}
