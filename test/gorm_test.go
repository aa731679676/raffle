package test

import (
	"fmt"
	"testing"

	"cn.a2490/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGorm(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../db/raffle.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connect database success!")
	// Migrate the schema
	db.AutoMigrate(&model.PrizeBase{})
	db.AutoMigrate(&model.RecordBase{})
	db.AutoMigrate(&model.RemarkBase{})

	// Create
	// db.Create(&model.UserBase{Name: "zhangsan", Phone: "22222222222"})

	// Read
	// var ub *model.UserBase = &model.UserBase{}
	// db.First(ub, 1)

	// fmt.Println(ub.Name)
}
