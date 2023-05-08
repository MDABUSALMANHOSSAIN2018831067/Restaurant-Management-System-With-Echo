package repositories

import (
	"errors"
	"fmt"
	"restaurant-management/pkg/domain"
	"restaurant-management/pkg/models"
	"gorm.io/gorm"
)

type FoodRepo struct {
	DB *gorm.DB
}

func FoodDBInterface(db *gorm.DB) domain.FoodRepoInterface {
	return &FoodRepo{
		DB: db,
	}

}

func (foodRepo *FoodRepo) CreateFood(food *models.Food) error {
	if err := foodRepo.DB.Where("name = ? AND price=?", food.Name, food.Price).First(&food).Error; err == nil {
		return errors.New("food name already exists")
	}
	if err := foodRepo.DB.Create(&food).Error; err != nil {
		return err
	}
	return nil
}

func (foodRepo *FoodRepo) GetFoods(ID uint) ([]models.Food, error) {
	fmt.Println(ID)
	var Food []models.Food
	if ID != 0 {
		if err := foodRepo.DB.Where("id = ?", ID).First(&Food).Error; err != nil {
			return nil, errors.New("food id not found")
		}
	} else {
		foodRepo.DB.Find(&Food)
	}
	return Food, nil
}

//	func (foodRepo *FoodRepo) UpdateFood(food *models.Food) error {
//		fmt.Println("kkkkkkk", food.Name)
//		if err := foodRepo.DB.Save(&food).Error; err != nil {
//			return err
//		}
//		return nil
//	}
func (foodRepo *FoodRepo) UpdateFood(food *models.Food) (*models.Food, error) {
	if err := foodRepo.DB.Model(&food).Save(&food).Error; err != nil {
		return nil, err
	}
	fmt.Println("food", food)
	return food, nil
}

func (foodRepo *FoodRepo) DeleteFood(ID uint) error {
	var Food models.Food
	if err := foodRepo.DB.
		Where("id = ?", ID).Delete(&Food).
		Error; err != nil {
		return err
	}
	return nil
}
