package service

import (
	"cn.a2490/dao"
	"cn.a2490/model"
	"cn.a2490/store"
)

type RecordService struct {
	recordDao *dao.RecordDao
	prizeDao  *dao.PrizeDao
}

func NewRecordService(recordDao *dao.RecordDao, prizeDao *dao.PrizeDao) *RecordService {
	return &RecordService{recordDao, prizeDao}
}

func (service *RecordService) DoDraw(userId uint) (*model.PrizeBase, error) {
	// 查询是否已经进行了抽奖
	recordBase := model.RecordBase{}
	service.recordDao.Model().First(&recordBase, "user_id = ?", userId)
	var prizeId uint
	if recordBase.ID != 0 {
		prizeId = recordBase.PrizeId
	} else {
		var err error
		prizeId, err = store.GetPrize(userId)
		if err != nil {
			return nil, err
		}
	}
	prize := model.PrizeBase{}
	service.prizeDao.Model().First(&prize, "id = ?", prizeId)
	return &prize, nil
}
