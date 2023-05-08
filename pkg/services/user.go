package services

import (
	"errors"
	"fmt"
	"restaurant-management/pkg/domain"
	"restaurant-management/pkg/models"
	"restaurant-management/pkg/redis"
	"restaurant-management/pkg/types"
	"strconv"
)

type UserService struct {
	repo domain.UserRepoInterface
}

func UserServiceInstance(userRepo domain.UserRepoInterface) domain.UserServiceInterface {
	return &UserService{
		repo: userRepo,
	}
}
func (service *UserService) RegistrationService(user *types.Registration) error {
	err := service.repo.Registration(user)
	return err
}

func (service *UserService) LoginService(email string) (*models.User, error) {
	users, err := service.repo.Login(email)
	return users, err

}

func (service *UserService) GetUserService(ID uint) ([]models.User, error) {
	useID := strconv.FormatUint(uint64(ID), 10)
	store := redis.NewRedisStore()
	getData, err := store.Get(useID)
	fmt.Println(err)
	if getData == nil {
		user, err := service.repo.GetUsers(ID)
		if err != nil {
			return nil, err
		}
		err = store.Set(useID, user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	// food, err := service.repo.GetUsers(ID)
	// if err != nil {
	// 	return nil, err
	// }
	return *getData, nil
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
