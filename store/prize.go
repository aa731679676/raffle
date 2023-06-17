package store

import (
	"crypto/rand"
	"errors"
	"math/big"
	"sort"
	"sync"
	"time"

	"cn.a2490/config"
	"cn.a2490/db"
	"cn.a2490/model"
)

type prizeStore struct {
	prizeList []model.PrizeBase
	mu        sync.Mutex
}

var store *prizeStore

func InitPrizeStore() {
	var prizeList []model.PrizeBase
	r := db.DB.Find(&prizeList)
	if r.Error != nil {
		panic("init prize store fail...")
	}
	if len(prizeList) <= 0 {
		panic("please set prize...")
	}
	prizes := processPrize(prizeList)
	// 根据概率排序
	sort.Slice(prizes, func(i, j int) bool {
		return (prizes)[i].Percent < (prizes)[j].Percent
	})
	store = &prizeStore{
		prizeList: prizes,
	}
}

func processPrize(prizeList []model.PrizeBase) []model.PrizeBase {
	var percentAll float32 = 0.0
	var prizes []model.PrizeBase = *new([]model.PrizeBase)
	// 核算总概率
	for _, prize := range prizeList {
		if prize.Remain > 0 {
			percentAll += prize.Percent
		}
	}
	// 重新计算概率
	var startPer float32 = 0.0
	for index, prize := range prizeList {
		if prize.Remain > 0 {
			percent := (prize.Percent / percentAll * 10000)
			prize.Percent = (startPer + percent)
			prizeList[index].Percent = (startPer + percent)
			startPer = prize.Percent
			prizes = append(prizes, prize)
		}
	}
	return prizes
}

func GetPrize(userId uint) (uint, error) {
	if len(store.prizeList) <= 0 {
		return 0, errors.New("no prize")
	}

	nowTime := time.Now()
	startTime := config.Config.Raffle.StartTime
	// 判断是否开始
	if startTime.After(nowTime) {
		return 0, errors.New("not start")
	}

	endTime := config.Config.Raffle.StartTime
	// 判断是否结束
	if nowTime.After(endTime) {
		return 0, errors.New("had end")
	}

	store.mu.Lock()
	defer store.mu.Unlock()

	// 生成 [0, 10000) 的随机数
	randRes, _ := rand.Int(rand.Reader, big.NewInt(10000))
	res := randRes.Int64()
	var prizeId uint = 0
	isGet := false
	for index, prize := range store.prizeList {
		if prize.Remain > 0 {
			if float32(res) < prize.Percent {
				isGet = true
				prizeId = prize.ID
				store.prizeList[index].Remain -= 1
				if prize.Remain == 0 {
					// 该商品被抽取完毕
					processPrize(store.prizeList)
				}
				// 存入数据库
				go saveRecord(userId, prizeId)
				break
			}
		}
	}
	if isGet {
		return prizeId, nil
	}
	return 0, errors.New("you are too late")
}

func saveRecord(userId uint, prizeId uint) {
	db.DB.Exec("UPDATE r_prize SET remain = remain-1 WHERE id = ?", prizeId)
	db.DB.Create(&model.RecordBase{UserId: userId, PrizeId: prizeId})
}
