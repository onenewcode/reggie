package config

import (
	"github.com/spf13/viper"
	"reggie/internal/global"
	"time"
)

type Setting struct {
	vp *viper.Viper
}

// 初始化一个配置类，让viper读取指定的配置文件
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("internal/config/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
func InitConfig() {
	setting, err := NewSetting()
	if err != nil {
		panic("配置文件读取错误")
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		panic("Server类读取错误，检查server类映射是否正确")
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		panic("App类读取错误，检查App类映射是否正确")
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		panic("Database类读取错误，检查Database类映射是否正确")
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
}
