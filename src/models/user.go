package models

import (
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

type User struct {
	Base
	Username    string `json:"username"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required" validate:"passwd"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Dob         string `json:"dob"`
	Provider    string `json:"provider" binding:"oneof=FACEBOOK GOOGLE NORMAL"`
	Status      string `json:"status" binding:"oneof=ACTIVE INACTIVE"`
}

func (u *User) validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (u User) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func (u *User) ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return err == nil
}
