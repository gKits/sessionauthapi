package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `json:"Password"`
}

func (u *User) EncryptPasswd() error {
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPasswd)

	return nil
}
