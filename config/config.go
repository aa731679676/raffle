package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type SqlType string

const (
	SqliteType SqlType = "sqllite"
	MysqlType  SqlType = "mysql"
)

type DbConfig struct {
	Type SqlType `yaml:"type"`
	Dsn  string  `yaml:"dsn"`
}

type TokenConfig struct {
	TokenKey     string `yaml:"tokenKey"`
	Timeout      int    `yaml:"timeout"`      // token有效期，单位s
	CleanTimeout int    `yaml:"cleanTimeout"` // 当token存储过长时的有效期，单位s
	CleanRegular int    `yaml:"cleanTimeout"` // 执行清理的间隔
	IsConcurrent bool   `yaml:"isConcurrent"` // 是否允许同一账号并发登录
	IsShare      bool   `yaml:"isShare"`      // 在多人登录同一账号时，是否共用一个token
}

type RaffleConfig struct {
	Title     string    `yaml:"title"`
	StartTime time.Time `yaml:"startTime"`
	EndTime   time.Time `yaml:"endTime"`
}

type ConfigCommon struct {
	Port   string        `yaml:"port"`
	Db     *DbConfig     `yaml:"db"`
	Token  *TokenConfig  `yaml:"token"`
	Raffle *RaffleConfig `yaml:"raffle"`
}

var Config *ConfigCommon = &ConfigCommon{
	Port: "8080",
	Db: &DbConfig{
		Type: SqliteType,
		Dsn:  "db/raffle.db",
	},
	Token: &TokenConfig{
		TokenKey:     "raffleKey",
		Timeout:      60 * 60 * 60 * 24 * 30,
		CleanTimeout: 60 * 60 * 60 * 24 * 3,
		CleanRegular: 60 * 60 * 60 * 2,
		IsConcurrent: true,
		IsShare:      true,
	},
	Raffle: &RaffleConfig{
		Title:     "抽奖活动",
		StartTime: time.Now(),
		EndTime:   time.Now().AddDate(0, 0, 10),
	},
}

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Read Config File Error: %s", err)
	}
	if err := viper.Unmarshal(Config); err != nil {
		panic(err)
	}
}
