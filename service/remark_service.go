package service

import (
	"cn.a2490/dao"
	"cn.a2490/model"
)

type RemarkService struct {
	*dao.RemarkDao
}

func NewRemarkService(remarkDao *dao.RemarkDao) *RemarkService {
	return &RemarkService{remarkDao}
}

func (service *RemarkService) List() *[]model.RemarkBase {
	var remarks []model.RemarkBase
	service.Model().Find(&remarks).Order("sort ASC")
	return &remarks
}
