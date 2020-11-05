package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oss_manage/app/system/api/request"
	"oss_manage/app/system/api/response"
	"oss_manage/app/system/model"
	"oss_manage/app/system/pkg/e"
	"oss_manage/app/system/service"
)

func Profile(c *gin.Context) {
	pro := service.GetProfile(c)
	if pro.Avatar == "" {
		pro.Avatar = e.DefaultAvatar
	}
	c.HTML(http.StatusOK, "profile.html", pro)
}

func AvatarEdit(c *gin.Context) {
	var f request.AvatarForm
	if err := c.ShouldBind(&f); err != nil {
		response.ErrorResp(c).SetMsg("上传失败"+err.Error()).SetType(model.OperEdit).Log(e.AvatarEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	err := service.UpdateAvatarService(f.Avatar, f.ID, c)
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.AvatarEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetMsg(f.Avatar).SetType(model.OperEdit).Log(e.AvatarEdit, c.Request.PostForm).WriteJsonExit()
}

func ProfileEdit(c *gin.Context) {
	var pro request.ProfileForm
	if err := c.ShouldBind(&pro); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.ProfileEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	if err := service.ProfileEditService(pro, c); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.ProfileEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetType(model.OperEdit).Log(e.ProfileEdit, c.Request.PostForm).WriteJsonExit()
}

func PwdEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "pwd.html", gin.H{})
}

func PwdEditHandler(c *gin.Context) {
	var pwd request.PasswordForm
	if err := c.ShouldBind(&pwd); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.PwdEditHandler, c.Request.PostForm).WriteJsonExit()
		return
	}
	if err := service.PwdEditHandlerService(pwd, c); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.PwdEditHandler, c.Request.PostForm).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetType(model.OperEdit).Log(e.PwdEditHandler, c.Request.PostForm).WriteJsonExit()
}
