package main

import (
	"fmt"
	"net/http"
	_ "oss_manage/app/system/controller"
	"oss_manage/app/system/db"
	"oss_manage/app/system/pkg/config"
	"oss_manage/app/system/router"
)

func initDB() {
	db.InitDB()
}

func main() {
	initDB()
	r := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    config.Setting.ReadTimeout,
		WriteTimeout:   config.Setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
