package dao

import (
	"oss_manage/app/system/db"
	"oss_manage/app/system/model"
)

type AdminOnlineDao interface {
	Insert(online model.AdminOnline) error
	Delete(sessionID string) error
}

func NewAdminOnlineDaoImpl() AdminOnlineDao {
	online := new(AdminOnlineDaoImpl)
	return online
}

type AdminOnlineDaoImpl struct {
}

func (a *AdminOnlineDaoImpl) Insert(online model.AdminOnline) error {
	err := db.GetDataBase().Instance().Create(&online).Error
	return err
}

func (a *AdminOnlineDaoImpl) Delete(sessionID string) error {
	err := db.GetDataBase().Instance().Where("session_id = ?", sessionID).Delete(model.AdminOnline{}).Error
	return err
}
