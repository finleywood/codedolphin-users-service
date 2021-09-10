package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Tier      uint8  `json:"tier"`
	Password  string `json:"-"`
}

type UserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (uldto *UserLoginDTO) VerifyPassword(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(uldto.Password))
	return err == nil
}

func (udto *UserDTO) HashPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(udto.Password), bcrypt.DefaultCost)
	udto.Password = string(hashedPassword)
}

func (udto *UserDTO) ToUser() *User {
	var user User
	user.FirstName = udto.FirstName
	user.LastName = udto.LastName
	user.Email = udto.Email
	user.Tier = 0
	user.Password = udto.Password
	return &user
}
