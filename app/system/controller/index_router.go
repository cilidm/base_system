package controller

import (
	"oss_manage/app/system/middleware"
	"oss_manage/app/system/router"
)

func init() {
	r := router.New("system", "/system", middleware.AuthMiddleware)
	r.GET("/", Index) // 有问题
	r.GET("index", Index)
	r.GET("main", FrameIndex)
}
