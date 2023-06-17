package service

import (
	"errors"
	"strconv"

	"cn.a2490/auth"
	"cn.a2490/dao"
	"cn.a2490/model"
)

type UserService struct {
	*dao.UserDao
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{userDao}
}

func (userService *UserService) CreateUser(phone string, name string) error {
	var count int64
	userService.Model().Where("phone = ?", phone).Count(&count)
	if count > 0 {
		return errors.New("this phone has exists")
	}
	userBase := &model.UserBase{Name: name, Phone: phone}
	userService.Model().Save(userBase)
	if userBase.ID != 0 {
		return nil
	}
	return errors.New("create user error")
}

func (userService *UserService) GetToken(phone string) (string, error) {
	userBase := model.UserBase{}
	userService.Model().First(&userBase, "phone = ?", phone)
	if userBase.ID == 0 {
		return "", errors.New("this phone is invalidity")
	}
	token := auth.DoLogin(strconv.Itoa(int(userBase.ID)))
	return token, nil
}
