package controller

import (
	"oss_manage/app/system/middleware"
	"oss_manage/app/system/router"
)

func init() {
	r := router.New("system", "/system", middleware.AuthMiddleware)
	r.GET("user/edit", Profile)
	r.POST("user/edit", ProfileEdit)
	r.POST("user/avatar", AvatarEdit)

	r.GET("user/pwd", PwdEdit)
	r.POST("user/pwd", PwdEditHandler)
}
