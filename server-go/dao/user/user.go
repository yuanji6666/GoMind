package user

import (
	"github.com/yuanji6666/gopherAI/common/mysql"
	"github.com/yuanji6666/gopherAI/schema"
	"github.com/yuanji6666/gopherAI/utils"
)

func InsertUser(user *schema.User) (*schema.User, error) {
	err := mysql.DB.Create(user).Error
	return user, err
}

func GetUserByUsername(username string) (*schema.User, error) {
	var user schema.User
	err := mysql.DB.Where("username=?", username).First(&user).Error
	return &user, err
}

func IsExistUser(username string) (bool, *schema.User) {
	user, err := GetUserByUsername(username)
	if err == nil {
		return true, user
	} else {
		return false, nil
	}
}

func Register(username, email, password string) (user *schema.User, ok bool) {
	user, err := InsertUser(&schema.User{
		Username: username,
		Name:     username,
		Email:    email,
		Password: utils.MD5(password),
	})
	if err != nil {
		return nil, false
	}
	return user, true
}
