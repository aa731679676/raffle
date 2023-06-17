package model

import "gorm.io/gorm"

type PrizeBase struct {
	gorm.Model
	Name    string  `gorm:"column:name;type:varchar(50)"`
	Total   int     `gorm:"column:total;type:int"`
	Remain  int     `gorm:"column:remain;type:int"`
	Percent float32 `gorm:"column:percent;type:decimal(4,2)"`
}

func (ub *PrizeBase) TableName() string {
	return "r_prize"
}
