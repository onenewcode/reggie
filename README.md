# 第一章
## 项目结构
首先介绍开始时我们的项目的总体结构，结构和注释如下，大家可以跟着创建相同的结构，也可以按照自己的思维进行一定的更改。
```shell
.
├── cmd
│   └── server
│       └── main.go         # 主程序入口，启动 HTTP 服务器
├── internal
│   ├── config
│   │   ├── config.go       # 配置文件读取与解析
│   │   ├── config.yaml     # 添加配置文件
│   ├── db
│   │   └──  db.go          # 数据库连接管理
│   ├── global              # 设置全局变量 
│   │   ├── db.go           # 设置全局数据库参数
│   │   └── setting.go      # 设置配置文件的具体参数
│   ├── service
│   └── util
├── pkg
├── router
├── public                  # 静态资源目录（HTML, CSS, JS等）
├── templates               # HTML模板文件
├── go.mod                  # Go模块定义文件
├── go.sum                  # 依赖包的校验和信息
└── README.md               # 项目文档与说明
```
## 创建配置文件
我们的配置文件采取yaml格式，读取配置文件的工具我们采用viper，这里本人使用的viper版本是v1.16.0。
我们在首先在根目录创建internal文件夹，在该文件夹下创建config文件夹，然后创建config.yaml文件。然后填入以下内容。
>internal/config/config.yaml
```yaml
Server:
  RunMode: debug
  HttpPort: :8080
  ReadTimeout: 60
  WriteTimeout: 60
App:
  DefaultPageSize: 10
  MaxPageSize: 100
  LogSavePath: storage/logs
  LogFileName: tools
  LogFileExt: .log
Database:
  DBType: mysql
  # 记得改成自己的数据库链接，第一个root是用户名，第二个root是密码
  Url: root:root@tcp(121.37.143.160:3306)/reggie?charset=utf8&parseTime=True&loc=Local
  TablePrefix: #设置表前缀
  Charset: utf8
  ParseTime: True
  MaxIdleConns: 10
  MaxOpenConns: 30
```
同时我们在internal文件夹下创建global文件夹，同时创建settiing.go文件。然后添加以下内容。
>internal>global>setting.go

```go
package global

import "time"

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

```
在上面的代码中var中包含的代码是我们设定的全局变量，主要是作为一个单例，方面其他包调用配置信息初始化自己的应用。

注意这里我们的结构体的定义要和我们yaml中定义的字段相同，但是可以忽略大小写。这是因为后面我们要用到viper进行读取配置文件的操作是进行字段映射，保持同名字段可以节约我们很多时间。

在config文件夹下创建config.go文件，然后添加以下内容。
>internal/config/config.go
```go
package config

import (
	"github.com/spf13/viper"
	"reggie/internal/global"
	"time"
)

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
	err = readSection(vp, "Server", &global.ServerSetting)
	if err != nil {
		panic("Server类读取错误，检查server类映射是否正确")
	}
	err = readSection(vp, "App", &global.AppSetting)
	if err != nil {
		panic("App类读取错误，检查App类映射是否正确")
	}
	err = readSection(vp, "Database", &global.DatabaseSetting)
	if err != nil {
		panic("Database类读取错误，检查Database类映射是否正确")
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
}
```
config.go中只有三个函数，其中configPath()函数表示读取指定文件目录下的配置文件,它返回一个viper.Viper对象的指针，现在我们写死，之后有机会了更改。InitConfig()负责初始化配置，把所有的数据读取后放入global的全局变量中，其中主要是调用了readSection()方法，readSection()是对viper中UnmarshalKey()方法的再封装，我们通过key和结构体来读取配置文件。
![img.png](images/img.png)
就如上图圈的，把他们的值作为key值来映射结构体。

## 创建服务器启动类
首先我们在根目录下创建cmd文件夹，然后在该文件夹下新建server文件夹，然后在server文件夹下创建main.go文件,然后添加以下内容。
>cmd/server/main.go
```go
package main

import (
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"reggie/internal/config"
	"reggie/internal/db"
	"reggie/internal/global"
)

func init() {
	config.InitConfig()
}

func main() {
	h := server.New(
		server.WithHostPorts(global.ServerSetting.HttpPort),
		server.WithReadTimeout(global.ServerSetting.ReadTimeout),
		server.WithWriteTimeout(global.ServerSetting.WriteTimeout),
	)
	h.Use(recovery.Recovery()) // 可确保即使在处理请求过程中发生未预期的错误或异常，服务也能维持运行状态
	h.Spin()                   //可以实现优雅的推出
}

```
值得注意的是，init()函数，它再整个类初始时调用，他负责初始化整个项目所需的内容。其次就是main函数，它负责运行起来整个web项目。


