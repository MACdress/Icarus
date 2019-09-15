package template

import "github.com/jinzhu/gorm"

type Module struct {
	 gorm.Model
	 ModuleCode string `gorm:"column:module_code"`
	 ModuleName string `gorm:"column:module_name"`
	 CreateUser string `gorm:"column:create_user"`
}