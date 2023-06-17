package main

import (
	"cn.a2490/auth"
	"cn.a2490/config"
	"cn.a2490/db"
	"cn.a2490/router"
	"cn.a2490/store"
)

// @title Raffle
// @version 1.0
// @description a toy of golang
func main() {
	config.InitConfig()
	db.InitDB()
	store.InitPrizeStore()
	stopCleanTokenFunc := auth.InitTokenStore()
	router.InitServer(stopCleanTokenFunc)
}
