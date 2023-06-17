package model

import "gorm.io/gorm"

type RecordBase struct {
	gorm.Model
	UserId  uint `gorm:"column:user_id;type:int"`
	PrizeId uint `gorm:"column:prize_id;type:int"`
}

func (ub *RecordBase) TableName() string {
	return "r_record"
}
