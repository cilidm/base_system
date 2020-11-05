package config

import (
	"github.com/go-ini/ini"
	"log"
	"sync"
	"time"
)

var (
	Cfg     *ini.File
	once    sync.Once
	Setting = &SettingConf{}
	server  = &Server{}
)

type Server struct {
	HTTPPort     int
	ReadTimeout  int
	WriteTimeout int
}

type SettingConf struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string
	UploadTmpDir string
	ImgSavePath  string
	ImgUrlPath   string
	DBType       string
	DBUser       string
	DBPwd        string
	DBHost       string
	DBTableName  string
	DBPath       string
	RedisAddr    string
	RedisPWD     string
	RedisDB      int
}

func init() {
	iniPath := "conf/app.ini"
	once.Do(func() {
		var err error
		Cfg, err = ini.Load(iniPath)
		if err != nil {
			log.Fatal("未发现配置文件")
		}
		err = Cfg.Section("runMode").MapTo(Setting)
		if err != nil {
			log.Fatal(err, "映射配置文件出错，请检查runMode配置")
		}
		err = Cfg.Section("app").MapTo(Setting)
		if err != nil {
			log.Fatal(err, "映射配置文件出错，请检查app配置")
		}
		err = Cfg.Section("database").MapTo(Setting)
		if err != nil {
			log.Fatal(err, "映射配置文件出错，请检查database配置")
		}
		err = Cfg.Section("redis").MapTo(Setting)
		if err != nil {
			log.Fatal(err, "映射配置文件出错，请检查redis配置")
		}

		err = Cfg.Section("server").MapTo(server)
		if err != nil {
			log.Fatal(err, "映射配置文件出错，请检查server配置")
		}
		Setting.HTTPPort = server.HTTPPort
		Setting.ReadTimeout = time.Duration(server.ReadTimeout) * time.Second
		Setting.WriteTimeout = time.Duration(server.WriteTimeout) * time.Second
	})
}
