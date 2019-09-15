package template

import "github.com/jinzhu/gorm"

type ModuleItem struct {
	gorm.Model
	ModuleCode string `gorm:"column:module_code"`
	Field string `gorm:"column:field"`
	IsRepeat int `gorm:"column:is_repeat"`//是否允许重复：0代表不允许重复，1代表允许重复
}
