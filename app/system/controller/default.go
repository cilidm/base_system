package controller

import (
	"fmt"
	files "github.com/cilidm/toolbox/file"
	"github.com/gin-gonic/gin"
	"oss_manage/app/system/api/response"
	"oss_manage/app/system/model"
	"oss_manage/app/system/pkg/config"
	"oss_manage/app/system/pkg/e"
	"path/filepath"
	"time"
)

func DefaultUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperAdd).Log(e.DefaultUpload, c.Request.PostForm).WriteJsonExit()
		return
	}
	if file.Size > e.DefUploadSize {
		response.ErrorResp(c).SetMsg("文件大小超限").SetType(model.OperAdd).Log(e.DefaultUpload, file).WriteJsonExit()
		return
	}
	fmt.Println(file.Filename, file.Size, file.Header)
	day := time.Now().Format(e.TimeFormatDay)
	savePath := filepath.Join(config.Setting.ImgSavePath, day) // 按年月日归档保存
	err = files.IsNotExistMkDir(savePath)
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperAdd).Log(e.DefaultUpload, file).WriteJsonExit()
		return
	}
	if err := c.SaveUploadedFile(file, filepath.Join(savePath, file.Filename)); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperAdd).Log(e.DefaultUpload, file).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetMsg(filepath.Join(filepath.Join(config.Setting.ImgUrlPath, day), file.Filename)).SetType(model.OperAdd).Log(e.DefaultUpload, file).WriteJsonExit()
}
