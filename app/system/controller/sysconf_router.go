package controller

import (
	"oss_manage/app/system/middleware"
	"oss_manage/app/system/router"
)

func init() {
	r := router.New("system", "/system", middleware.AuthMiddleware)

	r.GET("site/list", SiteList)
	r.POST("site/edit", SiteEdit)

	r.GET("mail/list", MailList)
	r.POST("mail/edit", MailEdit)
	r.POST("mail/test", MailTest)
}
