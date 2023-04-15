package repositories

import (
	"errors"
	"fmt"
	"restaurant-management/pkg/domain"
	"restaurant-management/pkg/models"

	"gorm.io/gorm"
)

type MenuRepo struct {
	DB *gorm.DB
}

func MenuDBInterface(db *gorm.DB) domain.MenuRepoInterface {
	return &MenuRepo{
		DB: db,
	}

}
func (menuRepo *MenuRepo) CreateMenu(menu *models.Menu) error {
	if err := menuRepo.DB.Where("id = ?", menu.FoodID).First(&models.Food{}).Error; err != nil {
		return errors.New("food id not found")
	}
	if err := menuRepo.DB.Where("category = ? AND food_id = ?", menu.Category, menu.FoodID).First(&menu).Error; err == nil {
		return errors.New("menu name already exists")
	}
	if err := menuRepo.DB.Create(&menu).Error; err != nil {
		return err
	}
	return nil
}

func (menuRepo *MenuRepo) GetMenus(ID uint) ([]models.Menu, error) {
	fmt.Println(ID)
	var menu []models.Menu
	query := menuRepo.DB.Preload("Food")
	if ID != 0 {
		if err := menuRepo.DB.Where("id = ?", ID).First(&menu).Error; err != nil {
			return nil, errors.New("menu id not found")
		}
	} else {
		if err := query.Find(&menu).Error; err != nil {
			return nil, err
		}
	}
	return menu, nil
}

func (menuRepo *MenuRepo) DeleteMenu(ID uint) error {
	var Food models.Food
	if err := menuRepo.DB.
		Where("id = ?", ID).Delete(&Food).
		Error; err != nil {
		return err
	}
	return nil
}

func (menuRepo *MenuRepo) UpdateManu(menu *models.Menu) (*models.Menu, error) {
	if err := menuRepo.DB.Where("id = ?", menu.FoodID).First(&models.Food{}).Error; err != nil {
		return nil, errors.New("food id not found")
	}
	if err := menuRepo.DB.Model(&menu).Save(&menu).Error; err != nil {
		return nil, err
	}
	return menu, nil
}
