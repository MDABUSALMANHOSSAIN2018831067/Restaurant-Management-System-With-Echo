package models

import (
	"errors"
	"regexp"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT" json:"_id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" `
	Password  string `json:"password,omitempty" `
	Email     string `json:"email,omitempty" `
	Phone     string `json:"phone,omitempty"`
	UserType  string `json:"user_type"  gorm:"default:'user'"`
	//UserID    string `json:"user_id"`
}

func NameValidate(name string) validation.RuleFunc {
	return func(value interface{}) error {
		name := value.(string)
		if _, err := strconv.Atoi(name); err == nil || len(name) < 4 || len(name) > 200 {
			return errors.New("please enter valid name")
		}
		return nil
	}
}

func EmailValidate(email string) validation.RuleFunc {
	return func(value interface{}) error {
		email := value.(string)
		pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		re := regexp.MustCompile(pattern)
		err := re.MatchString(email)
		if !err {
			return errors.New("invalid email address")
		}
		return nil
	}
}

// var UserRule = []validation.Rule{
// 	validation.Required,
// 	validation.Length(2, 20),
// }

func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.FirstName, validation.Required,
			validation.By(NameValidate(user.FirstName))),
		validation.Field(&user.LastName, validation.Required,
			validation.By(NameValidate(user.LastName))),
		validation.Field(&user.Email, validation.Required, validation.By(EmailValidate(user.Email))),
		validation.Field(&user.Password, validation.Required),
		validation.Field(&user.Phone, validation.Required),
		validation.Field(&user.UserType),
	)
}
