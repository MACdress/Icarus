package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UName string `gorm:"column:u_name"`
	UCode string `gorm:"column:u_code"`
	UPwd  string `gorm:"column:u_pwd"`
}
