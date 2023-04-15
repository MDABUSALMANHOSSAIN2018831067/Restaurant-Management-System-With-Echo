package controllers

import (
	"fmt"
	"log"
	"net/http"
	"restaurant-management/pkg/consts"
	"restaurant-management/pkg/domain"
	"restaurant-management/pkg/helpers"
	"restaurant-management/pkg/models"
	"restaurant-management/pkg/types"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService domain.UserServiceInterface
}

func SetUserService(userService *domain.UserServiceInterface) *UserController {
	return &UserController{
		userService: *userService,
	}
}

// Registration user
func (userController *UserController) Registration(e echo.Context) error {
	reqUser := &models.User{}
	if err := e.Bind(reqUser); err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidInput)
	}
	if err := reqUser.Validate(); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	password := HashPassword(reqUser.Password)
	reqUser.Password = password
	if err := userController.userService.RegistrationService(reqUser); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "User was registrated successfully")

}

// Login user
func (userController *UserController) Login(c echo.Context) error {
	var user = &types.User{}
	var model_user = &models.User{}
	var tokens = types.Token{}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}

	if err := user.Validate(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	model_user, err := userController.userService.LoginService(user.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	passwordIsValid, msg := VerifyPassword(user.Password, model_user.Password)
	if !passwordIsValid {
		return c.JSON(http.StatusInternalServerError, msg)
	}
	token, refreshToken, err := helpers.GenerateAllTokens(model_user.Email, model_user.UserType)
	tokens.UserToken = token
	tokens.UserRefreshtoken = refreshToken
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tokens)
}

// Get user
func (userController *UserController) GetUsers(e echo.Context) error {
	tempID := e.QueryParam("_id")
	ID, err := strconv.ParseInt(tempID, 0, 0)
	if err != nil && tempID != "" {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}

	user, err := userController.userService.GetUserService(uint(ID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, user)
}

// Delete user
func (userController *UserController) DeleteUser(e echo.Context) error {
	tempID := e.Param("_id")
	ID, err := strconv.ParseUint(tempID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}

	_, err = userController.userService.GetUserService(uint(ID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if err := userController.userService.DeleteUserService(uint(ID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, "User was deleted successfully")
}

// UpdateUser update an user
func (userController *UserController) UpdateUser(c echo.Context) error {
	var user = &models.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id := c.Param("_id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, consts.InvalidID)
	}
	oldres, err := userController.userService.GetUserService(uint(userID))
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	user.ID = uint(userID)

	checkedUser := UpdateUserField(user, oldres)

	res, err := userController.userService.UpdateUserService(checkedUser)
	if err != nil || res == nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "update was success")
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = "email or password is incorrect"
		check = false
	}
	return check, msg
}

func UpdateUserField(user *models.User, old_user []models.User) *models.User {
	if user.FirstName == "" {
		user.FirstName = old_user[0].FirstName
	}
	if user.LastName == "" {
		user.LastName = old_user[0].LastName
	}
	if user.UserType == "" {
		user.UserType = old_user[0].UserType
	}
	if user.Password == "" {
		user.Password = old_user[0].Password
	}
	if user.Phone == "" {
		user.Phone = old_user[0].Phone
	}
	return user
}
