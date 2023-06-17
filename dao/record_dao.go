package dao

import (
	"cn.a2490/db"
	"cn.a2490/model"
	"gorm.io/gorm"
)

type RecordDao struct {
}

func NewRecordDao() *RecordDao {
	return &RecordDao{}
}

func (recordDao *RecordDao) Model() *gorm.DB {
	return db.DB.Model(&model.RecordBase{})
}
