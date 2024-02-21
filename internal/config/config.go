package config

import (
	"github.com/spf13/viper"
	"time"
)

// 全局变量，提供给内部的其他包使用
var (
	ServerSetting   *ServerSettingS
	AppSetting      *AppSettingS
	DatabaseSetting *DatabaseSettingS
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	DBType       string
	Url          string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// 初始化一个配置类，让viper读取指定的配置文件
func configPath() (*viper.Viper, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("internal/config/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return vp, nil
}

func readSection(vp *viper.Viper, k string, v interface{}) error {
	err := vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}

// 初始化配置，把所有的数据读取后放入global的全局变量中
func InitConfig() {
	vp, err := configPath()
	if err != nil {
		panic("配置文件读取错误")
	}
	err = readSection(vp, "Server", &ServerSetting)
	if err != nil {
		panic("Server类读取错误，检查server类映射是否正确")
	}
	err = readSection(vp, "App", &AppSetting)
	if err != nil {
		panic("App类读取错误，检查App类映射是否正确")
	}
	err = readSection(vp, "Database", &DatabaseSetting)
	if err != nil {
		panic("Database类读取错误，检查Database类映射是否正确")
	}

	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second
}
