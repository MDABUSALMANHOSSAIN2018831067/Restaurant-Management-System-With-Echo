package repositories

import (
	"errors"
	"restaurant-management/pkg/connection"
	"restaurant-management/pkg/models"
)

func CreateOrderItem(orderItem *models.OrderItem) error {
	DB := connection.GetDB()
	if err := DB.Where("id = ?", orderItem.FoodID).First(&models.Food{}).Error; err != nil {
		return errors.New("food id not found")
	}
	if err := DB.Where("quantity = ? AND food_id = ?", orderItem.Quantity, orderItem.FoodID).
		First(&orderItem).Error; err == nil {
		return errors.New("order item already exists")
	}
	if err := DB.Create(&orderItem).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderItems(ID uint) ([]models.OrderItem, error) {
	DB := connection.GetDB()
	var orderItem []models.OrderItem
	// if err := DB.Where("id = ?", orderItem[0].FoodID).First(&models.Food{}).Error; err != nil {
	// 	return nil, errors.New("food id not found")
	// }
	query := DB.Preload("Food")
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

func DeleteOrderItem(ID uint) error {
	DB := connection.GetDB()
	var item models.OrderItem
	if err := DB.
		Where("id = ?", ID).Delete(&item).
		Error; err != nil {
		return err
	}
	return nil
}

func UpdateOrderItem(item *models.OrderItem) (*models.OrderItem, error) {
	DB := connection.GetDB()
	if err := DB.Where("id = ?", item.FoodID).First(&models.Food{}).Error; err != nil {
		return nil, errors.New("food id not found")
	}
	if err := DB.Model(&item).Save(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}
