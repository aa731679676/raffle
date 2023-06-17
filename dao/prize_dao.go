package dao

import (
	"cn.a2490/db"
	"cn.a2490/model"
	"gorm.io/gorm"
)

type PrizeDao struct {
}

func NewPrizeDao() *PrizeDao {
	return &PrizeDao{}
}

func (prizeDao *PrizeDao) Model() *gorm.DB {
	return db.DB.Model(&model.PrizeBase{})
}
