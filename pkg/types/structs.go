package types

import (
	jwt "github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type SignedDetails struct {
	Email    string
	UserType string
	jwt.StandardClaims
}

type User struct {
	Email    string `json:"email" `
	Password string `json:"password" `
	UserType string `json:"user_type" `
}

func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required),
		validation.Field(&user.Password, validation.Required),
		validation.Field(&user.UserType))

}

type Token struct {
	UserToken        string
	UserRefreshtoken string
}

// package types

// import (
// 	jwt "github.com/dgrijalva/jwt-go"
// )

// type SignedUserDetails struct {
// 	Id    int
// 	Email string
// 	Type  string
// 	jwt.StandardClaims
// }

// type User struct {
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required"`
// 	Type     string `json:"type" validate:"required"`
// }

// type Token struct {
// 	User_Token        string
// 	User_Refreshtoken string
// }
