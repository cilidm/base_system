package dao

import (
	"oss_manage/app/system/db"
	"oss_manage/app/system/model"
)

type SysConfDao interface {
	Insert(conf model.SysConf) error
	Update(id uint, conf model.SysConf) error
	FindBySysType(sysType model.SysType) (conf model.SysConf, err error)
}

func NewSysConfDaoImpl() SysConfDao {
	conf := new(SysConfDaoImpl)
	return conf
}

type SysConfDaoImpl struct {
}

func (s *SysConfDaoImpl) Insert(conf model.SysConf) error {
	err := db.GetDataBase().Instance().Where("type = ?", conf.Type).FirstOrCreate(&conf).Error
	return err
}

func (s *SysConfDaoImpl) Update(id uint, conf model.SysConf) error {
	err := db.GetDataBase().Instance().Model(model.SysConf{}).Where("id = ?", id).Updates(&conf).Error
	return err
}

func (s *SysConfDaoImpl) FindBySysType(sysType model.SysType) (conf model.SysConf, err error) {
	err = db.GetDataBase().Instance().Model(model.SysConf{}).Where("status = 1 AND type = ?", sysType).First(&conf).Error
	return conf, err
}
