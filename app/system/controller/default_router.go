package controller

import (
	"oss_manage/app/system/middleware"
	"oss_manage/app/system/router"
)

func init() {
	r := router.New("system", "/system", middleware.AuthMiddleware)
	r.POST("upload/def_upload", DefaultUpload)
}
