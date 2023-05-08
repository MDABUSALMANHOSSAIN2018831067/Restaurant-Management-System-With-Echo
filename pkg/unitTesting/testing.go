package unittesting

import (
	"fmt"
	"restaurant-management/pkg/domain"
	"restaurant-management/pkg/models"
	"testing"
)

type UserTesting struct {
	userService domain.UserServiceInterface
}

func SetUserTestingService(userService *domain.UserServiceInterface) *UserController {
	return &UserTesting{
		userService: *userService,
	}
}

var u = models.User{
	ID: 1,
}

func (service *UserTesting) UserAPITesting(t testing.T) {
	user := &models.User{}
	expected := user

	result, err := service.userService.LoginService(user.Email)
	if err != nil {
		fmt.Println(err)
	}

	if result != expected {
		t.Errorf("\"sayHello('%s')\" FAILED, expected -> %v, got -> %v", user, expected, result)
	} else {
		t.Logf("\"sayHello('%s')\" SUCCEDED, expected -> %v, got -> %v", user, expected, result)
	}
}
