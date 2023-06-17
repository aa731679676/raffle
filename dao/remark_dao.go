package dao

import (
	"cn.a2490/db"
	"cn.a2490/model"
	"gorm.io/gorm"
)

type RemarkDao struct {
}

func NewRemarkDao() *RemarkDao {
	return &RemarkDao{}
}

func (remarkDao *RemarkDao) Model() *gorm.DB {
	return db.DB.Model(&model.RemarkBase{})
}
