# 第一章 构建项目框架
## 初始化项目结构
首先介绍开始时我们的项目的总体结构，结构和注释如下，大家可以跟着创建相同的结构，也可以按照自己的思维进行一定的更改。
```shell
.
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
├── main.go                 # 主程序入口，启动 HTTP 服务器
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
首先我们在根目录下创建main.go文件,然后添加以下内容。
>main.go
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
	EmpDao   = EmployeeDao{}
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
InitDB()函数负责初始化数据库链接，并把初始化后的链接放在一个DBEngine的全局变量中。在InitDB()函数中我们暂时只用到配置文件中Database.Url后续会添加更多配置可选项。而且会提供一个叫EmpDao的全局变量，以供用户操作employee表经行增删改查操作。

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
在这个文件中暂时只实现一个Login()函数，如果查询不到用户就返回404，状态码设置位0，查找到后就设置状态码1，data设置为查找到的用户。其中本文件的进行查询时，用到了internal/db/db.go文件提供的全局变量。

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
>main.go
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

# 添加JWT和swagger文档

## jwt
JWT（JSON Web Tokens）是一种开放标准（RFC 7519），用于在网络应用环境间安全地传输声明信息。它的主要作用包括：

- 身份验证（Authentication）： JWT通常用于用户登录认证，当用户成功登录后，服务器会生成一个包含用户身份信息和其他必要数据的JWT令牌，并将其发送给客户端。客户端在后续请求中将此令牌作为HTTP Header的一部分发送回服务器。服务器通过验证这个令牌来确认用户的身份和权限，无需再查询数据库或存储session信息。

- 授权（Authorization）： JWT可以携带访问控制所需的权限声明（claims），这些声明定义了持有者能访问哪些资源或者执行哪些操作。服务器根据令牌中的声明信息进行权限校验，决定是否允许客户端进行特定的操作。

- 无状态性（Statelessness）： 因为所有的必要信息都编码在JWT中，服务器不再需要维护客户端的状态信息，从而简化了服务端的设计并有利于实现水平扩展。

- 安全性（Security）： JWT通过数字签名（Signature）确保了数据的完整性和不可篡改性。使用密钥对令牌进行签名，服务器可以通过验证签名来确认令牌未被篡改且来自可信的源。

- 跨域支持（CORS）： JWT可以在多个域名、子域之间共享，有助于实现单点登录（SSO）和微服务架构下的认证与授权需求。

- 减少服务器负载（Reduced Server Load）： 由于JWT包含了足够的信息以供鉴权和授权，服务器在处理JWT时不需要额外查询数据库，从而降低了数据库查询次数，提高了系统的响应速度。

在我们的项目中jwt主要用于身份验证。
### 创建EmployeeLoginVO
这里我们为了让前端便于处理我们的数据和接受token，我们需要创建一个EmployeeLoginVO结构体。我们要在models文件夹下创建vo文件夹，然后新建common.go文件。
>internal/models/vo/common.go
```go
package vo

type EmployeeLoginVO struct {
	Id int64 `json:"id,omitempty"`

	UserName string `json:"user_name,omitempty"`

	Name string `json:"name,omitempty"`

	Token string `json:"token,omitempty"`
}
```
### 创建用于通知的静态变量。
首先我们要在models文件夹下创建constant文件夹，并且在该文件夹下创建messages_c文件夹和status_c文件夹，并且在这两个文件夹都创建一个名叫common.go的文件。然后添加以下内容。
>internal/models/constant/messages_c/common.go
```go
package message_c

/**
 * 信息提示常量类
 */
const (
	PASSWORD_ERROR                 = "密码错误"
	ACCOUNT_NOT_FOUND              = "账号不存在"
	ACCOUNT_LOCKED                 = "账号被锁定"
	UNKNOWN_ERROR                  = "未知错误"
	USER_NOT_LOGIN                 = "用户未登录"
	CATEGORY_BE_RELATED_BY_SETMEAL = "当前分类关联了套餐,不能删除"
	CATEGORY_BE_RELATED_BY_DISH    = "当前分类关联了菜品,不能删除"
	SHOPPING_CART_IS_NULL          = "购物车数据为空，不能下单"
	ADDRESS_BOOK_IS_NULL           = "用户地址为空，不能下单"
	LOGIN_FAILED                   = "登录失败"
	UPLOAD_FAILED                  = "文件上传失败"
	SETMEAL_ENABLE_FAILED          = "套餐内包含未启售菜品，无法启售"
	PASSWORD_EDIT_FAILED           = "密码修改失败"
	DISH_ON_SALE                   = "起售中的菜品不能删除"
	SETMEAL_ON_SALE                = "起售中的套餐不能删除"
	DISH_BE_RELATED_BY_SETMEAL     = "当前菜品关联了套餐,不能删除"
	ORDER_STATUS_ERROR             = "订单状态错误"
	ORDER_NOT_FOUND                = "订单不存在"
)


```
>internal/models/constant/status_c/common.go
```go
package status_c

/**
 * 状态常量，启用或者禁用
 */
const (
	//启用
	ENABLE = 1

	//禁用
	DISABLE = 0
)

```
我们这样定义全局的错误报错可以让我们比较方便的修改报错信息。
### 构建JWT中间件
首先我们要在internal文件夹下创建middleware文件夹，然后创建jwt.go文件，接着我们便开始完善我们的jwt文件。
>internal/middleware/jwt.go
```go
package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"log"
	"net/http"
	"reggie/internal/db"
	"reggie/internal/models/common"
	"reggie/internal/models/constant/message_c"
	"reggie/internal/models/constant/status_c"
	"reggie/internal/models/model"
	"reggie/internal/models/vo"
	"time"
)

var (
	// 设置我们存储的信息在jwt中的哪一个字段
	identityKey string = "reggie"
	// 设置从哪里获取jwt的信息，格式如下
	// - "header:<name>"
	// - "query:<name>"
	// - "cookie:<name>"
	// - "param:<name>"
	// - "form:<name>"
	jwtToken = "header: token"
)

// 设置标识处理函数
// 这里我们把通过定义identityKey获取负载的数据
func jwtIdentityHandler(ctx context.Context, c *app.RequestContext) interface{} {
	claims := jwt.ExtractClaims(ctx, c)
	return claims[identityKey]
}

// 生成jwt负载的函数，指定了Authenticator方法生成的数据如何存储和怎么样存储c.Get("JWT_PAYLOAD")访问
func jwtPayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*vo.EmployeeLoginVO); ok {
		return jwt.MapClaims{
			identityKey: v,
		}
	}
	return jwt.MapClaims{}
}

func jwtLoginResponse(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
	var elv, _ = c.Get(identityKey)
	rely := elv.(*vo.EmployeeLoginVO)
	rely.Token = token
	c.JSON(http.StatusOK, common.Result{1, "", rely})
}

// 返回值会被存在Claim数组中
func jwtAuthenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var empl model.Employee
	if err := c.BindAndValidate(&empl); err != nil {
		return nil, common.Result{0, jwt.ErrMissingLoginValues.Error(), nil}
	}
	emp := db.EmpDao.GetByUserName(empl.Username)
	var errorR common.Result
	log.Println(emp)
	if emp.Username != empl.Username {
		// 账号不存在
		errorR = common.Result{0, message_c.ACCOUNT_NOT_FOUND, nil}
		return nil, errorR
	}

	//密码比对
	if empl.Password != emp.Password {
		//密码错误
		errorR = common.Result{0, message_c.PASSWORD_ERROR, nil}
		return nil, errorR
	}

	if empl.Status == status_c.DISABLE {
		//账号被锁定
		errorR = common.Result{0, message_c.ACCOUNT_LOCKED, nil}
		return nil, errorR
	}

	elv := vo.EmployeeLoginVO{
		Id:       emp.ID,
		UserName: emp.Username,
		Name:     emp.Name,
		Token:    "",
	}
	// 这里我们把对象值存入c中，方便在返回函数中进行包装
	c.Set(identityKey, &elv)
	return &elv, nil

}
func InitJwtAdmin() *jwt.HertzJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm: "test zone",
		// 用于签名的密钥
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		// 用于在JWT中存储用户唯一标识身份的键值
		IdentityKey: identityKey,
		// 用于生成JWT载荷部分的声明
		PayloadFunc: jwtPayloadFunc,
		// 作用在登录成功后的每次请求中，用于设置从 token 提取用户信息的函数
		IdentityHandler: jwtIdentityHandler,
		// 用于设置登录时认证用户信息的函数
		Authenticator: jwtAuthenticator,
		LoginResponse: jwtLoginResponse,
		// 设置从哪里获取jwt的信息
		TokenLookup: jwtToken,
		// 不设置jwt表名前缀
		WithoutDefaultTokenHeadName: true,
		//  当用户未通过身份验证或授权时，调用此函数返回错误信息
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			// 不通过，响应401状态码
			c.String(http.StatusNotFound, message)
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	return authMiddleware
}

```
这里我们主要提供了一个对外暴露的InitJwtAdmin()函数，它可以让路由获取整个初始化后的中间件。接下来将讲述这了中间件的具体化初始步骤。
接下来将从上倒下介绍必要最重要的初始化参数。

首先介绍四个比较简单的阐述
- Realm: "test zone", 	 用于设置所属领域名称，默认为 hertz jwt
- Key:        []byte("secret key"), 用于设置签名密钥（必要配置）
- Timeout:    time.Hour, 用于设置 token 过期时间，默认为一小时
- MaxRefresh: time.Hour, 用于设置最大 token 刷新时间，允许客户端在 TokenTime + MaxRefresh 内刷新 token 的有效时间，追加一个 Timeout 的时长
- IdentityKey: 用于设置检索身份的键，

接下来我们介绍一对一般搭配的参数，`PayloadFunc`和I`identityHandler`，其中`PayloadFunc`的作用是指明我们的数据如何存储于存储在jwt的具体Claims中，我们实现的自定义函数`jwtPayloadFunc`直接把数据转换成`EmployeeVO`存储在jwt中名叫reggie的自定义字段中。`identityHandler`是用于设置获取身份信息的函数，它指明如何解析请求获取的jwt token的有效信息，我们在自定义的` jwtPayloadFunc`,把数据存在reggie字段中，现在通过该字段获取数据即可。具体可以看`jwtIdentityHandler`的实现。

`Authenticator`用于设置登录时认证用户信息的函数。其中我们自己实现的`jwtAuthenticator`函数作为Authenticator的参数，它在用户登陆时获取用户的用户名然后查询数据库，查询出来后就对各种属性进行依次判断，判断它是否合法，不合法就返回相应的错误信息，合法就对`Employee`包装为`EmployeeVo`类，然后把包装存储在`*app.RequestContext`,同时把返回值设置为包装后的对象。

**注意** `jwtAuthenticator`函数的返回值将会被加密为我们的jwt令牌。

`LoginResponse`参数，他表示用户登陆成功后的返回值，在我们实现的自定义`jwtLoginResponse`中我们先获取在`jwtAuthenticator`存储的`EmployeeVO`对象，然后把生成的token值赋予它，之后在进行包装之后直接返回。
`Unauthorized`参数是用户认证失败调用的函数，如果`Authenticator`的data返回值为nil就会调用这个函数。
以上就是我们中间件实现流程的大致介绍。

### 添加路由
因为我们引入入了中间件，所以我们登陆这一功能不在依赖于employee_service.gow文件。所以我么要修改router.go文件，修改结果如下。
>internal/router/router.go
```go
package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"net/http"
	"reggie/internal/middleware"
)

func InitRouter(r *server.Hertz) {
	myJwt := middleware.InitJwtAdmin()
	adm := r.Group("/admin")

	emp := adm.Group("/employee")
	emp.POST("/login", myJwt.LoginHandler)
	// 注意我们要把登陆放到中间件的前面，因为一旦启用中间件，接下来的请求都需要经过jwt的校验
	adm.Use(myJwt.MiddlewareFunc())
	{
		// 这里必须新生成一个emp，因为新生成的才含有我们的中间件
		emp := adm.Group("/employee")
		// 这是个测试方法，之后会测试我们的jwt是否拦截
		emp.GET("/test", func(c context.Context, ctx *app.RequestContext) {
			ctx.String(http.StatusOK, "Fds")
		})
	}

}
```
上面的代码，我们引入中间件，并让中间件校验`/admin`下的文件，我们还写了一个测试方法，可以通过访问 http://localhost:8080/admin/employee/test 测试我们是否有校验过我们的jwt令牌。

### 测试jwt
在我们启动项目后，可以通过postman或者apifox工具对我们的功能进行测试。这里以apifox为例。

在打开工具后我们访问 http://localhost:8080/admin/employee/login 如果返回结果如下就证明我们的jwt加密成功。
```shell
{
    "code": 1,
    "msg": "",
    "data": {
        "id": 1,
        "user_name": "admin",
        "name": "管理员",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDg1ODEwMDgsIm9yaWdfaWF0IjoxNzA4NTc3NDA4LCJyZWdnaWUiOnsiaWQiOjEsInVzZXJfbmFtZSI6ImFkbWluIiwibmFtZSI6IueuoeeQhuWRmCJ9fQ.haQWlngPcb4Di3HGlJl4J4UTboWE9ROnXiqnnYHjrag"
    }
}
```
我们也可以通过在线的jwt解密网站判断加密的信息是否正确。
这是上面jwt令牌的解析结果
```shell
{
  "exp": 1708581008,
  "orig_iat": 1708577408,
  "reggie": {
    "id": 1,
    "user_name": "admin",
    "name": "管理员"
  }
}
```
大家也可以自行访问 该网址 https://www.bejson.com/jwt/ 经行判断。

接下来要测试我们的jwt 拦截是否成功 http://localhost:8080/admin/employee/test 拦截成功会出现以下结果。
![im](images/img_1.png)

接下来我们在我们的测试接口中，添加我们的jwt 令牌，我们在header中添加token字段，并把我们访问成功生成的jwt token值粘贴在里面。
![im](images/img_2.png)

接着我们再次访问。访问成功便会有以下结果
![im](images/img_3.png)

这样我们整个jwt的功能就算大致完成了。

## swagger
wagger（现在被称为OpenAPI Specification，OAS）是一个用于描述RESTful API的标准格式和工具集。它允许开发人员使用YAML或JSON格式定义API的端点、输入参数、输出数据结构、HTTP方法以及认证要求等详细信息。

通过Swagger提供的工具和服务，可以实现以下功能：
1. API文档： Swagger自动生成交互式的API文档，让开发者能直观地了解API的各项接口及其用法，无需阅读冗长的文字说明。
2. 设计与开发一致性： 在API设计阶段即可使用Swagger来定义API规范，确保在开发过程中遵循一致的设计标准。
3. 模拟服务器： Swagger工具支持基于API定义生成模拟服务器，使得客户端开发者可以在后端服务尚未完成时就开始集成测试。
4. 代码生成： 根据Swagger定义，可以为多种编程语言生成客户端库或者服务器端框架代码，大大提高了开发效率。
5. 验证与测试： 开发者可以直接在Swagger UI上尝试调用API，并验证其响应结果是否符合预期，简化了API的调试过程。
6. 自动化集成： 可以将Swagger文件与CI/CD流程相结合，自动进行API的构建、部署和版本管理。

总之，Swagger/OpenAPI Specification是API生命周期管理中非常重要的一个工具，从设计到开发、测试再到最终的文档发布，都能提供有力的支持。
### 创建swagger
首先我们要安装swagger。
因为从 Go 1.17 开始，在 go mod 模式下通过 go get 下载对应库文件将无法自动编译并安装到 $GOPATH/bin 的路径， 所以不再推荐用 go get 来安装可执行文件的方式。可以使用 go install来代替。
在项目的根目录输入以下内容来安装swagger。
>go install github.com/swaggo/swag/cmd/swag@latest

然后在命令行运行
>swag init

这这是便会在根目录下生成一个名为docs的文件夹，这个文件夹下便记录了我们的所有注释的api信息。

我们需要在middleware文件夹下创建swagger.go文件，我们在该文件添加以下内容。
>internal/middleware/swagger.go
```go
package middleware

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

func InitSwagger(r *route.RouterGroup) {
	url := swagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition

	r.GET("/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
}

```

其中InitSwagger()函数负责初始化swagger相关的内容，这个函数的主要目的是把swagger文档的内容映射在一个路径上。

接着我们要修改路由。在router文件夹下的router.go文件添加以下内容。
>internal/router/router.go
```go
func InitRouter(r *server.Hertz) {
	swa := r.Group("/swagger")
	{
		middleware.InitSwagger(swa)
	}
	myJwt := middleware.InitJwtAdmin()

	adm := r.Group("/admin")
	emp := adm.Group("/employee")
	emp.POST("/login", myJwt.LoginHandler)
	// 注意我们要把登陆放到中间件的前面，因为一旦启用中间件，接下来的请求都需要经过jwt的校验
	adm.Use(myJwt.MiddlewareFunc())
	{
		// 这里必须新生成一个emp，因为新生成的才含有我们的中间件
		emp := adm.Group("/employee")
		// 这是个测试方法，之后会测试我们的jwt是否拦截
		emp.GET("/test", func(c context.Context, ctx *app.RequestContext) {
			ctx.String(http.StatusOK, "Fds")
		})
	}

}
```
然后我们需要在main.go中导入自己的docs文件路径。比如本项目叫reggic就导入以下路径
>_ "reggie/docs"

这样我们的swagger就配置完成了。
### 测试
首先我在根文件夹下的main.go文件添加注释。
>main.go
```go
package main

import (
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	_ "reggie/docs"
	"reggie/internal/config"
	"reggie/internal/db"
	"reggie/internal/router"
)

func init() {
	config.InitConfig()
	db.InitDB()
}

// @title regiee
// @version 0.1
// @description sky-take-out

// @contact.name onenewcode
// @contact.url https://github.com/onenewcode

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
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
然后我们在根目录运行
>swag init  

然后运行项目，访问 http://localhost:8080/swagger/index.html 如果出现以下界面就说明设置swagger配置成功，
![swagger](images/img_4.png)