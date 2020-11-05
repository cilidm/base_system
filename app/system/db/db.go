package db

import (
	"fmt"
	"github.com/cilidm/toolbox/logging"
	"github.com/gchaincl/dotsql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"oss_manage/app/system/model"
	"oss_manage/app/system/pkg/config"
	"sync"
)

type DataBase interface {
	Open()
	Close() (err error)
	Instance() *gorm.DB
}

type DataBaseImpl struct {
	db   *gorm.DB
	once sync.Once
}

func (d *DataBaseImpl) Open() {
	d.once.Do(func() {
		var err error
		d.db, err = gorm.Open(config.Setting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Setting.DBUser,
			config.Setting.DBPwd,
			config.Setting.DBHost,
			config.Setting.DBTableName))
		if err != nil {
			logging.Fatal("数据库连接失败:【" + err.Error() + "】")
		}

		d.db.SingularTable(true)
		if config.Setting.RunMode != "release" {
			d.db.LogMode(true)
		}
		//d.db.DB().SetMaxIdleConns(10)
		//d.db.DB().SetMaxOpenConns(100)
		d.initModel()       // 校验数据表
		d.initMyCallbacks() // 注册回调函数
	})
}

func (d *DataBaseImpl) Close() (err error) {
	err = d.db.Close()
	return
}

func (d *DataBaseImpl) Instance() *gorm.DB {
	if d.db == nil {
		logging.Fatal("无法连接数据库")
	}
	return d.db
}

var ModelGroup []interface{}

func (d *DataBaseImpl) initModel() {
	d.checkTableData(&model.Admin{})
	d.checkTableData(&model.AdminOnline{})
	d.checkTableData(&model.Auth{})
	d.checkTableData(&model.LoginInfo{})
	d.checkTableData(&model.OperLog{})
	d.checkTableData(&model.Role{})
	d.checkTableData(&model.RoleAuth{})
	d.checkTableData(&model.SysConf{})

	if len(ModelGroup) > 0 {
		for _, m := range ModelGroup {
			d.checkTableData(m)
		}
	}
}

func (d *DataBaseImpl) checkTableData(tb interface{}) {
	if d.db.HasTable(tb) == false {
		d.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(tb)
		var sqlName string
		if _, ok := tb.(*model.Admin); ok {
			sqlName = "create-admin"
		} else if _, ok := tb.(*model.Auth); ok {
			sqlName = "create-auth"
		} else if _, ok := tb.(*model.Role); ok {
			sqlName = "create-role"
		} else if _, ok := tb.(*model.RoleAuth); ok {
			sqlName = "create-role-auth"
		}
		if sqlName != "" {
			d.initData(sqlName)
		}
	}
}

func (d *DataBaseImpl) initData(sqlName string) {
	dot, err := dotsql.LoadFromFile("data/base_system.sql")
	if err != nil {
		logging.Fatal("无法加载初始数据，请检查data文件夹下是否存在数据信息")
	}
	_, err = dot.Exec(d.db.DB(), sqlName)
	if err != nil {
		logging.Warn("执行"+sqlName+"失败", err)
		return
	}
}

func (d *DataBaseImpl) initMyCallbacks() {
	d.db.Callback().Create().Replace("gorm:update_time_stamp", model.ForBeforeCreate)
	d.db.Callback().Update().Replace("gorm:update_time_stamp", model.ForBeforeUpdate)
}
