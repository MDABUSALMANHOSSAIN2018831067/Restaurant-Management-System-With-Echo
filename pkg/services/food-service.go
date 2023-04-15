package services

import (
	"errors"
	"restaurant-management/pkg/domain"
	"restaurant-management/pkg/models"
)

type FoodService struct {
	repo domain.FoodRepoInterface
}

func FoodServiceInstance(foodRepo domain.FoodRepoInterface) domain.FoodServiceInterface {
	return &FoodService{
		repo: foodRepo,
	}
}
func (service *FoodService) CreateFoodService(food *models.Food) error {
	err := service.repo.CreateFood(food)
	return err
}

func (service *FoodService) GetFoodService(ID uint) ([]models.Food, error) {
	foods, err := service.repo.GetFoods(ID)
	return foods, err
}

func (service *FoodService) UpdateFoodService(food *models.Food) (*models.Food, error) {
	food, err := service.repo.UpdateFood(food)
	if err != nil {
		return nil, errors.New("food update was unsuccesful")
	}
	return food, nil
}

func (service *FoodService) DeleteFoodService(ID uint) error {
	if err := service.repo.DeleteFood(ID); err != nil {
		return errors.New("food deletion was unsuccessful")
	}
	return nil
}
