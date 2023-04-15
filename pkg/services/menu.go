package services

import (
	"errors"
	"restaurant-management/pkg/domain"
	"restaurant-management/pkg/models"
)

type MenuService struct {
	repo domain.MenuRepoInterface
}

func MenuServiceInstance(menuRepo domain.MenuRepoInterface) domain.MenuServiceInterface {
	return &MenuService{
		repo: menuRepo,
	}
}
func (service *MenuService) CreateMenuService(menu *models.Menu) error {
	err := service.repo.CreateMenu(menu)
	return err
}

func (service *MenuService) GetMenuService(ID uint) ([]models.Menu, error) {
	foods, err := service.repo.GetMenus(ID)
	return foods, err
}

func (service *MenuService) UpdateManuService(menu *models.Menu) (*models.Menu, error) {
	food, err := service.repo.UpdateManu(menu)
	if err != nil {
		return nil, errors.New("food update was unsuccesful")
	}
	return food, nil
}

func (service *MenuService) DeleteMenuService(ID uint) error {
	if err := service.repo.DeleteMenu(ID); err != nil {
		return errors.New("food deletion was unsuccessful")
	}
	return nil
}
