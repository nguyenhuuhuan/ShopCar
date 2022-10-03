package models

import "net/mail"

type User struct {
	Base
	Username string `json:"username"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name"`
	Dob      string `json:"dob"`
	Provider string `json:"provider" binding:"oneof=FACEBOOK GOOGLE NORMAL"`
	Status   string `json:"status" binding:"oneof=ACTIVE INACTIVE"`
}

func (u *User) validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
