package user

import "gorm.io/gorm"

// Entity 模型

// User 用户实体
type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);index" json:"email"`
	Username string `gorm:"type:varchar(50);uniqueIndex" json:"username"`
	Nickname string `gorm:"type:varchar(50)" json:"nickname"`
	Password string `gorm:"type:varchar(255);" json:"-"`
}
