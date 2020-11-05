package controller

import (
	"oss_manage/app/system/middleware"
	"oss_manage/app/system/router"
)

func init() {
	r := router.New("system", "/")
	r.GET("/", middleware.CheckDefaultPage)
	r.GET("login", middleware.CheckLoginPage, Login)
	r.POST("login", LoginHandler)
	r.GET("logout", Logout)
	r.POST("isLogin", nil)
	r.GET("not_found", NotFound)
}
