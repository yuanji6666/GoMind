package user

import (
	"github.com/yuanji6666/gopherAI/common/mysql"
	"github.com/yuanji6666/gopherAI/utils"
)

// Repository 数据访问层

// InsertUser 插入新用户
func InsertUser(userEntity *User) (*User, error) {
	err := mysql.DB.Create(userEntity).Error
	return userEntity, err
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(username string) (*User, error) {
	var user User
	err := mysql.DB.Where("username=?", username).First(&user).Error
	return &user, err
}

// IsExistUser 判断用户是否存在
func IsExistUser(username string) (bool, *User) {
	user, err := GetUserByUsername(username)
	if err == nil {
		return true, user
	} else {
		return false, nil
	}
}

// RegisterUser 注册新用户
func RegisterUser(username, email, password string) (user *User, ok bool) {
	user, err := InsertUser(&User{
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
