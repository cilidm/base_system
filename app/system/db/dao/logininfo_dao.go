package dao

import (
	"oss_manage/app/system/db"
	"oss_manage/app/system/model"
)

type LoginInfoDao interface {
	Insert(info model.LoginInfo) error
}

func NewLoginInfoImpl() LoginInfoDao {
	info := new(LoginInfoDaoImpl)
	return info
}

type LoginInfoDaoImpl struct {
}

func (l *LoginInfoDaoImpl) Insert(info model.LoginInfo) error {
	err := db.GetDataBase().Instance().Create(&info).Error
	return err
}
