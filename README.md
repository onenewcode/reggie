# 第一章 构建项目框架
## 初始化项目结构
首先介绍开始时我们的项目的总体结构，结构和注释如下，大家可以跟着创建相同的结构，也可以按照自己的思维进行一定的更改。
```shell
.
├── cmd
│   └── server
│       └── main.go         # 主程序入口，启动 HTTP 服务器
├── internal
│   ├── config
│   │   ├── config.go       # 配置文件读取与解析
│   │   └── config.yaml     # 添加配置文件
│   ├── db
│   │   ├── employee.go     # 存储dao操作 
│   │   └──  db.go          # 数据库连接管理
│   ├── models              # 存储各种通用类
│   │   ├── common          # 通用类文件夹
│   │   │   └── common.go   # 存储返回值
│   │   └──  model          # 实体类文件
│   │       └── employee.go   
│   ├── router              # 存放路由
│   │   ├── api             # 存放不同业务
│   │   │   └── employee.go   
│   │   └──  router.go      # 负责注册路由
│   └── util                # 存放工具    
├── nginx-1.20.2            # 项目前端服务，双击nginx.exe启动不能放在含有中文目录的地方
├── pkg
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
同时我们在internal文件夹下创建config文件夹，同时创建config.go文件。然后添加以下内容。
>internal>config>config.go

```go
package config

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

然后继续在config.go添加以下内容。
>internal/config/config.go
```go
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
## 创建各种实体类
这里我们按需创建，要实现什么就创建什么实体类。
首先在internal文件夹下创建models文件夹，然后在models文件夹下创建model和common文件夹，在common文件夹下创建common.go文件。在该文件夹下添加以下内容。
>internal/models/common/common.go
```go
package common

type Result struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

```
Result类主要是为了包装返回值，让我们的返回值有一个通用返回值，好让前端处理，Data的类型作为一个空接口可以接受任意类型。

然后在model文件夹下创建employee.go文件，然后在文件里添加以下内容。
>internal/models/model/employee.go
```go
package model

import (
	"time"
)

const TableNameEmployee = "employee"

// Employee 员工信息
type Employee struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`        // 主键
	Name       string    `gorm:"column:name;not null;comment:姓名" json:"name"`                         // 姓名
	Username   string    `gorm:"column:username;not null;comment:用户名" json:"username"`                // 用户名
	Password   string    `gorm:"column:password;not null;comment:密码" json:"password"`                 // 密码
	Phone      string    `gorm:"column:phone;not null;comment:手机号" json:"phone"`                      // 手机号
	Sex        string    `gorm:"column:sex;not null;comment:性别" json:"sex"`                           // 性别
	IDNumber   string    `gorm:"column:id_number;not null;comment:身份证号" json:"id_number"`             // 身份证号
	Status     int32     `gorm:"column:status;not null;default:1;comment:状态 0:禁用，1:启用" json:"status"` // 状态 0:禁用，1:启用
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`                  // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                  // 更新时间
	CreateUser int64     `gorm:"column:create_user;comment:创建人" json:"create_user"`                   // 创建人
	UpdateUser int64     `gorm:"column:update_user;comment:修改人" json:"update_user"`                   // 修改人
}

// TableName Employee's table name
func (*Employee) TableName() string {
	return TableNameEmployee
}

```
这是我们的实体类，每个字段后面都有一个tag字段，它是为了在解析时，解析到相应的字段，其中grom标签grom在解析时字段的对应，json指明了在转化json不同字段的对应关系。而且TableName()函数必不可少，它是为了指定grom在解析的时候表明。

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
	"reggie/internal/router"
)

func init() {
	config.InitConfig()
}

func main() {
	h := server.New(
		server.WithHostPorts(config.ServerSetting.HttpPort),
		server.WithReadTimeout(config.ServerSetting.ReadTimeout),
		server.WithWriteTimeout(config.ServerSetting.WriteTimeout),
	)
	router.InitRouter(h)
	h.Use(recovery.Recovery()) // 可确保即使在处理请求过程中发生未预期的错误或异常，服务也能维持运行状态
	h.Spin()                   //可以实现优雅的推出
}


```
值得注意的是，init()函数，它再整个类初始时调用，他负责初始化整个项目所需的内容。其次就是main函数，它负责运行起来整个web项目。
## 链接数据库
首先我们在internal文件夹下创建db文件夹然后在db文件夹下创建db.go，db.go文件主要负责初始化数据库，并提供全局的数据库链接，以供其它的模块使用。
>internal/db/db.go
```go

package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reggie/internal/config"
)

var (
	DBEngine *gorm.DB
)

func InitDB() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.DatabaseSetting.Url, // DSN data source name
		DefaultStringSize:         256,                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("数据库链接失败")
	}
	DBEngine = db
}

```
InitDB()函数负责初始化数据库链接，并把初始化后的链接放在一个DBEngine的全局变量中。在InitDB()函数中我们暂时只用到配置文件中Database.Url后续会添加更多配置可选项。

接下来在db文件夹下创建employee_dao.go文件，这个文件主要负责employee表的各种sql操作。
>internal/db/employee_dao.go
> 
```go
package db

import "reggie/internal/models/model"

type EmployeeDao struct {
}

func (*EmployeeDao) GetByUserName(username string) *model.Employee {
	var emp model.Employee
	DBEngine.Where("username=?", username).First(&emp)
	return &emp
}
```
在这里我们定义了EmployeeDao类，所有的sql方法都由他实现，如果我们需要调用dao层操作，只需要生成一个类即可，调用所有的方法。

## 创建路由
首先我们在internal文件夹下创建api文件夹和router.go文件。接着在api文件夹下创建employee_service.go文件，然后在文件夹下添加以下内容。
>internal/router/api/employee_service.go
```go
package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"reggie/internal/db"
	"reggie/internal/models/common"
	"reggie/internal/models/model"
)

var (
	empDao = db.EmployeeDao{}
)

func Login(ctx context.Context, c *app.RequestContext) {
	var empL model.Employee
	// 参数绑定转化为结构体
	err := c.Bind(&empL)
	if err != nil {
		log.Println("Employee 参数绑定失败")
	}
	//password := c.Query("password")
	emp := empDao.GetByUserName(empL.Username)
	if emp == nil {
		c.JSON(http.StatusNotFound, common.Result{0, "未知用户", nil})
	}
	c.JSON(http.StatusOK, common.Result{1, "", emp})

}
```
在这里我们先初始化一个私有的全局变量，接下来所有要调用sql的操作都通过这个私有的全局变量。在这个文件中暂时只实现一个Login()函数，如果查询不到用户就返回404，状态码设置位0，查找到后就设置状态码1，data设置为查找到的用户。

接下来我们就在router文件夹下的router.go文件夹下添加以下内容。
>internal/router/router.go
```go
package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"reggie/internal/router/api"
)

func InitRouter(r *server.Hertz) {
	// 为每个静态资源目录创建一个 http.FileServer
	emp := r.Group("/admin/employee")
	emp.POST("/login", api.Login)
}

```
router.go文件主要是为了注册不同的路由，并提供了一个InitRouter()函数，为了方便主函数调用，初始化路由。

## 更改主函数
在这里我们要为主函数添加初始化路由和初始化数据库。
>internal/cmd/server/main.go
```go
package main

import (
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"reggie/internal/config"
	"reggie/internal/db"
	"reggie/internal/router"
)

func init() {
	config.InitConfig()
	db.InitDB()
}

func main() {
	h := server.New(
		server.WithHostPorts(config.ServerSetting.HttpPort),
		server.WithReadTimeout(config.ServerSetting.ReadTimeout),
		server.WithWriteTimeout(config.ServerSetting.WriteTimeout),
	)
	router.InitRouter(h)
	h.Use(recovery.Recovery()) // 可确保即使在处理请求过程中发生未预期的错误或异常，服务也能维持运行状态
	h.Spin()                   //可以实现优雅的推出
}

```
这样我们的整体框架就完成了。
## 测试
我们双击nginx-1.20.2文件夹下的nginx.exe,启动前端程序，然后启动服务端程序。在浏览器访问 http://localhost/#/login 。
![img.png](images/img1.png)
点击登陆，出现以下界面就表示部署成功。
![img.png](images/img2.png)