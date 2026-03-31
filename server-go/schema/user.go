package schema

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50)" json:"name"`
	Email    string `gorm:"type:varchar(100);index" json:"email"`
	Username string `gorm:"type:varchar(50);uniqueIndex" json:"username"`
	Password string `gorm:"type:varchar(255);" json:"-"`
}
