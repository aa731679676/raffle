package model

import (
	"gorm.io/gorm"
)

type RemarkBase struct {
	gorm.Model
	Content string `gorm:"column:content;type:varchar(200)"`
	Sort    int    `gorm:"column:sort;type:int"`
}

func (ub *RemarkBase) TableName() string {
	return "r_remark"
}
