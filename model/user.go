package model

import (
	"gin-my/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=20" label:"用户名"`
	Password string `gorm:"type:varchar(70);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
}

func CheckUser(username string) int {
	var user User
	db.Select("id,username").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return errmsg.USER_EXIST
	}
	return errmsg.SUCCSE
}

func CreateUser(user *User) int {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errmsg.INTERNAL_SERVER_ERROR
	}
	user.Password = string(hashPassword)
	err = db.Create(&user).Error
	if err != nil {
		return errmsg.INTERNAL_SERVER_ERROR
	}
	return errmsg.SUCCSE
}

func CheckLogin(username, password string) (User, int) {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return user, errmsg.USER_NOT_EXIST
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errmsg.ERROR_PASSWORD
	}
	return user, errmsg.SUCCSE
}

func GetUser(username string) User {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return user
	}
	return user
}
