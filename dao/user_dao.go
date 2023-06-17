package dao

import (
	"cn.a2490/db"
	"cn.a2490/model"
	"gorm.io/gorm"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (userDao *UserDao) Model() *gorm.DB {
	return db.DB.Model(&model.UserBase{})
}
