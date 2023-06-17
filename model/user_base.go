package model

import (
	"gorm.io/gorm"
)

type UserBase struct {
	gorm.Model
	Name  string `gorm:"column:name;type:varchar(30)"`
	Phone string `gorm:"column:phone;type:varchar(20)"`
}

func (ub *UserBase) TableName() string {
	return "r_user"
}
