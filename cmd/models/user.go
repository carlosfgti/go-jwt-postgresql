package models

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func (user *User) Create(map[string]interface{}) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(&user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Token{UserId: user.ID})
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))
	user.Token = tokenString

	return user.Token
}
