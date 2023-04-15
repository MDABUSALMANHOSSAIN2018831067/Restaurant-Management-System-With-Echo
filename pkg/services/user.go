package services

import (
	"errors"
	"restaurant-management/pkg/domain"
	"restaurant-management/pkg/models"
)

type UserService struct {
	repo domain.UserRepoInterface
}

func UserServiceInstance(userRepo domain.UserRepoInterface) domain.UserServiceInterface {
	return &UserService{
		repo: userRepo,
	}
}
func (service *UserService) RegistrationService(user *models.User) error {
	err := service.repo.Registration(user)
	return err
}

func (service *UserService) LoginService(email string) (*models.User, error) {
	foods, err := service.repo.Login(email)
	return foods, err

}

func (service *UserService) GetUserService(ID uint) ([]models.User, error) {
	food, err := service.repo.GetUsers(ID)
	if err != nil {
		return nil, err
	}
	return food, nil
}

func (service *UserService) DeleteUserService(ID uint) error {
	if err := service.repo.DeleteUser(ID); err != nil {
		return errors.New("user deletion was unsuccessful")
	}
	return nil
}

func (service *UserService) UpdateUserService(user *models.User) (*models.User, error) {
	user, err := service.repo.UpdateUser(user)
	if err != nil {
		return nil, errors.New("food update was unsuccesful")
	}
	return user, nil
}
