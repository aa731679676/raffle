package db

import (
	"log"
	"os"
	"time"

	"cn.a2490/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dbType := config.Config.Db.Type

	var dia gorm.Dialector = nil

	switch dbType {
	case config.SqliteType:
		dia = sqlite.Open(config.Config.Db.Dsn)
	case config.MysqlType:
		dia = mysql.Open(config.Config.Db.Dsn)
	}
	if dia == nil {
		panic("Database Init Error")
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL的阈值
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	var err error
	DB, err = gorm.Open(dia, &gorm.Config{
		PrepareStmt: true,
		Logger:      newLogger,
	})
	if err != nil {
		panic("Failed To Connect Database")
	}
}
