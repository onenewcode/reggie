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
│   │   ├── admin            # 存放不同业务
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
IDNumber   string    `gorm:"column:id_number;not null;comment:身份证号" json:"idNumber"`              // 身份证号
Status     int32     `gorm:"column:status;not null;default:1;comment:状态 0:禁用，1:启用" json:"status"` // 状态 0:禁用，1:启用
CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`                  // 创建时间
UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                  // 更新时间
CreateUser int64     `gorm:"column:create_user;comment:创建人" json:"create_user"`                   // 创建人
UpdateUser int64     `gorm:"column:update_user;comment:修改人" json:"update_user"`                   // 修改人
}

// TableName Employee's table name
func (*Employee) TableName() string {
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
首先我们在internal文件夹下创建admin文件夹和router.go文件。接着在admin文件夹下创建employee_router.go文件，然后在文件夹下添加以下内容。
>internal/router/admin/employee_router.go
```go
package admin

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
	"reggie/internal/router/admin"
)

func InitRouter(r *server.Hertz) {
	// 为每个静态资源目录创建一个 http.FileServer
	emp := r.Group("/admin/employee")
	emp.POST("/login", admin.Login)
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

# 第二章 JWT和swagger文档

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

const (
	IdentityKey = "reggie"
	// 设置我们存储的信息在jwt中的哪一个字段
	// 设置从哪里获取jwt的信息，格式如下
	// - "header:<name>"
	// - "query:<name>"
	// - "cookie:<name>"
	// - "param:<name>"
	// - "form:<name>"
	JwtToken = "header: token"

)
// 设置标识处理函数
// 这里我们把通过定义identityKey获取负载的数据
func jwtIdentityHandler(ctx context.Context, c *app.RequestContext) interface{} {
	claims := jwt.ExtractClaims(ctx, c)
	return claims[IdentityKey]
}

// 生成jwt负载的函数，指定了Authenticator方法生成的数据如何存储和怎么样存储c.Get("JWT_PAYLOAD")访问
func jwtPayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*vo.EmployeeLoginVO); ok {
		return jwt.MapClaims{
			IdentityKey: v,
		}
	}
	return jwt.MapClaims{}
}

func jwtLoginResponse(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
	var elv, _ = c.Get(IdentityKey)
	rely := elv.(*vo.EmployeeLoginVO)
	rely.Token = token
	c.JSON(http.StatusOK, common.Result{1, "", rely})
}

// 返回值会被存在Claim数组中
func jwtAuthenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var empl model.Employee
	if err := c.BindAndValidate(&empl); err != nil {
		log.Println(jwt.ErrMissingLoginValues.Error())
		return nil, common.Result{0, jwt.ErrMissingLoginValues.Error(), nil}
	}
	emp := db.EmpDao.GetByUserName(empl.Username)
	var errorR common.Result
	if emp.Username != empl.Username {
		log.Println(message_c.ACCOUNT_NOT_FOUND)
		// 账号不存在
		errorR = common.Result{0, message_c.ACCOUNT_NOT_FOUND, nil}
		return nil, errorR
	}

	//密码比对
	if empl.Password != emp.Password {
		log.Println(message_c.PASSWORD_ERROR)
		//密码错误
		errorR = common.Result{0, message_c.PASSWORD_ERROR, nil}
		return nil, errorR
	}

	if emp.Status == status_c.DISABLE {
		log.Println(message_c.ACCOUNT_LOCKED)
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
	c.Set(IdentityKey, &elv)
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
		IdentityKey: IdentityKey,
		// 用于生成JWT载荷部分的声明
		PayloadFunc: jwtPayloadFunc,
		// 作用在登录成功后的每次请求中，用于设置从 token 提取用户信息的函数
		IdentityHandler: jwtIdentityHandler,
		// 用于设置登录时认证用户信息的函数
		Authenticator: jwtAuthenticator,
		// 登陆回复
		LoginResponse: jwtLoginResponse,
		LogoutResponse: func(ctx context.Context, c *app.RequestContext, code int) {
			c.JSON(code, common.Result{1, "", nil})
		},
		// 设置从哪里获取jwt的信息
		TokenLookup: JwtToken,
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
因为我们引入入了中间件，所以我们登陆这一功能不在依赖于employee_router.gow文件。所以我么要修改router.go文件，修改结果如下。
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
		emp.POST("/logout", myJwt.LogoutHandler)
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
package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"net/http"
	"reggie/internal/middleware"
)

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
		emp.POST("/logout", myJwt.LogoutHandler)
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
其中swagger的注释可以参考以下链接 https://github.com/swaggo/swag/blob/master/README_zh-CN.md#%E5%A3%B0%E6%98%8E%E5%BC%8F%E6%B3%A8%E9%87%8A%E6%A0%BC%E5%BC%8F
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

# 第三章 完善管理员功能
- 新增员工
- 员工分页查询
- 启用禁用员工账号
- 编辑员工
- 导入分类模块功能代码

## 新增员工

**注意事项：**

1. 账号必须是唯一的
2. 手机号为合法的11位手机号码
3. 身份证号为合法的18位身份证号码
4. 密码默认为123456

###  代码开发
首先我们要在 status_c和messages_s文件夹的common.go文件添加一些新的静态变量。
>internal/models/constant/status_c/common.go
```go
//设置新用户的默认密码
	DEFAULT_PASSWORD = "123456"
```
>internal/models/constant/messages_s/common.go
```go
ALREADY_EXISTS                 = "已存在"
```
#### 添加路由
我们要在路由中添加以下内容
>internal/router/router.go
```go
func InitRouter(r *server.Hertz) {
	// 添加日志
	r.Use(accesslog.New(accesslog.WithFormat("[${time}] ${status} - ${latency} ${method} ${path} ${queryParams}")))
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
		// 启动jwt
		emp.POST("/logout", myJwt.LogoutHandler)
		// 添加雇员接口
		// 这是个测试方法，之后会测试我们的jwt是否拦截
		emp.GET("/test", func(c context.Context, ctx *app.RequestContext) {
			ctx.String(http.StatusOK, "Fds")
		})

	}

}
```
#### router层

**employee_router.go中创建新增员工方法**
>internal/router/admin/employee.go
```go
// 存储用户
// @Summary 存储用户
// @Accept application/json
// @Produce application/json
// @router /admin/employee [post]
func SaveEmp(ctx context.Context, c *app.RequestContext) {
var empL model.Employee
// 参数绑定转化为结构体
err := c.Bind(&empL)
if err != nil {
log.Println("Employee 参数绑定失败")
}
log.Printf("新增用户:{%s}", empL.Username)
flag := service.SavEmp(&empL)
if flag == true {
c.JSON(http.StatusOK, common.Result{1, "", nil})
}
c.JSON(http.StatusBadRequest, common.Result{1, message_c.ALREADY_EXISTS, nil})
}
```

**注:** Result类定义了后端统一返回结果格式。其具体格式请参考
>internal/models/common/common.go
```go
package common

import "encoding/json"

type Result struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r Result) Error() string {
	jsonBytes, _ := json.Marshal(r)

	// 将JSON字节转为字符串并打印
	return string(jsonBytes)
}

```


####  Service层

**在EmployeeServiceImpl中实现新增员工方法**
>internal/router/service/employee_service.go

```go
// 添加成功返回true，添加失败返回flase
func SavEmp(emp *model.Employee) bool {
//设置账号的状态，默认正常状态 1表示正常 0表示锁定
emp.Status = status_c.ENABLE

//设置密码，默认密码123456
emp.Password = status_c.DEFAULT_PASSWORD

//设置当前记录的创建时间和修改时间
emp.CreateTime, emp.UpdateTime = time.Now(), time.Now()

//设置当前记录创建人id和修改人id
emp.CreateUser, emp.UpdateUser = 1, 1 //目前是假数据，之后会继续完善
// 判断是否用户是否重名
if db.EmpDao.GetByUserName(emp.Username).Username == emp.Username {
return false
}
db.EmpDao.Insert(emp)
return true
}

```




#### dao层

**在employee_dao.go中声明Insert方法**
>internal/db/employee/employee.go

```go
func (*EmployeeDao) Insert(emp *model.Employee) {
    DBEngine.Create(emp)
}
```

### 功能测试

代码已经发开发完毕，对新增员工功能进行测试。

**功能测试实现方式：**

- 通过接口文档测试
- 通前后端联调测试

接下来我们使用上述两种方式分别测试。



####  接口文档测试
这里我们用apifox的测试工具经行测试。
我们的访问地址为 http://localhost:8080/admin/employee 然后在请求中添加以下的json内容。
json数据：
```json
{
  "id": 0,
  "idNumber": "111222333444555666",
  "name": "xiaozhi",
  "phone": "13812344321",
  "sex": "1",
  "username": "小智"
}
```
具体的成果如下图
![api测试](images/img_5.png)
然后我们运行项目进行测试。返回结果为以下：
- 响应码：404 报错
- 请求体: auth header is empty

这是因为我们jwt进行了校验，需要用户登陆才能添加用户。所以我们要先调用登陆api，然后把获取到的登陆的token粘贴到，我们添加用户的请求头。
![api测试](images/img_6.png)

然后我们在添加令牌之后重新测试添加用户这一个api。

**查看employee表：**
![api测试](images/img_7.png)

测试成功。



####  前后端联调测试

启动nginx,访问 http://localhost

登录-->员工管理-->添加员工

![api测试](images/img_8.png)
保存后，查看employee表
![api测试](images/img_9.png)
测试成功。



**注意:** 由于开发阶段前端和后端是并行开发的，后端完成某个功能后，此时前端对应的功能可能还没有开发完成，
导致无法进行前后端联调测试。所以在开发阶段，后端测试主要以接口文档测试为主。




### 代码完善

目前，程序存在的问题主要有一个：

- 新增员工时，创建人id和修改人id设置为固定值



**描述**：新增员工时，创建人id和修改人id设置为固定值

**分析：**

```go
func SavEmp(emp *model.Employee) bool {
/////
//设置当前记录创建人id和修改人id
emp.CreateUser, emp.UpdateUser = 1, 1 //目前是假数据，之后会继续完善
/////
}
```

**解决：**

通过某种方式动态获取当前登录员工的id。
![d](images/img_10.png)

员工登录成功后会生成JWT令牌并响应给前端，jwt令牌中存贮这一些用户的信息：后续请求中，前端会携带JWT令牌，通过JWT令牌可以解析出当前登录员工id：
我们的中间件会把jwt 令牌信息存储在app.RequestContext中，其中可以通过`c.Get("JWT_PAYLOAD")`获取，数据的存储格式为map[string]interface{}嵌套结构组成的。

所以我们要修改employee_router.go中的SaveEmp函数，让它获取用户的id数据，注入一个`Employee`类中的，`UpdateUser`和`CreateUser`属性，并传递给service层。
>internal/router/admin/employee_router.go
```go
func SaveEmp(ctx context.Context, c *app.RequestContext) {
	var empL model.Employee
	// 参数绑定转化为结构体
	err := c.Bind(&empL)
	if err != nil {
		log.Println("Employee 参数绑定失败")
		c.JSON(http.StatusBadRequest, common.Result{1, message_c.UNKNOWN_ERROR, nil})
	} else {
		// 获取jwt_payload的信息,并把信息赋予empL
		{
			jwt_payload, _ := c.Get("JWT_PAYLOAD")
			// 类型转换,我们的数据在claims中是以map[string]interface{}嵌套结构组成的。
			claims := jwt_payload.(jwt.MapClaims)
			origin_emp := claims[middleware.IdentityKey].(map[string]interface{})
			emp_id := origin_emp["id"].(float64)
			empL.CreateUser, empL.UpdateUser = int64(emp_id), int64(emp_id)
		}
		log.Printf("新增用户:{%s}", empL.Username)
		flag := service.SavEmp(&empL)
		if flag == true {
			c.JSON(http.StatusOK, common.Result{1, "", nil})
		}
		c.JSON(http.StatusBadRequest, common.Result{1, message_c.ALREADY_EXISTS, nil})
	}
}
```
同时我们还要删除service层给`UpdateUser`和`CreateUser`属性赋值的代码。
```go
func SavEmp(emp *model.Employee) bool {
	//设置账号的状态，默认正常状态 1表示正常 0表示锁定
	emp.Status = status_c.ENABLE

	//设置密码，默认密码123456
	emp.Password = status_c.DEFAULT_PASSWORD

	//设置当前记录的创建时间和修改时间
	emp.CreateTime, emp.UpdateTime = time.Now(), time.Now()

	//设置当前记录创建人id和修改人id
	//emp.CreateUser, emp.UpdateUser = 1, 1 //目前是假数据，之后会继续完善
	// 判断是否用户是否重名
	if db.EmpDao.GetByUserName(emp.Username).Username == emp.Username {
		return false
	}
	db.EmpDao.Insert(emp)
	return true
}
```



## 员工分页查询

### 需求分析和设计

####  产品原型

系统中的员工很多的时候，如果在一个页面中全部展示出来会显得比较乱，不便于查看，所以一般的系统中都会以分页的方式来展示列表数据。而在我们的分页查询页面中, 除了分页条件以外，还有一个查询条件 "员工姓名"。

**查询员工原型：**
![im](images/img_11.png)

**业务规则**：

- 根据页码展示员工信息
- 每页展示10条数据
- 分页查询时可以根据需要，输入员工姓名进行查询

###  代码开发

#### 设计DTO类

根据请求参数进行封装，要在dto的common.go文件下添加以下内容
>internal/models/dto/common.go
```go
type EmployeePageQueryDTO struct {
//员工姓名
Name *string `json:"name,omitempty" form:"name,omitempty"`
//页码
Page int `json:"page,omitempty" form:"page,omitempty"`
//每页显示记录数
PageSize int `json:"pageSize,omitempty" form:"pageSize,omitempty"`
}

```



####  封装PageResult

后面所有的分页查询，统一都封装为PageResult对象。我们要在common文件夹下的common.go添加以下内容。
>internal/models/common/common.go

```go

type PageResult struct {
Total   int64       `json:"total,omitempty"`   //总记录数
Records interface{} `json:"records,omitempty"` //当前页数据集合
}

```

员工信息分页查询后端返回的对象类型为: Result<PageResult>,我们可以查看common/common.go的Result结构体。
>internal/models/common/common.go
```go
type Result struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r Result) Error() string {
	jsonBytes, _ := json.Marshal(r)

	// 将JSON字节转为字符串并打印
	return string(jsonBytes)
}
```


#### 添加路由
我们要在根路由添加我们的分页路由。
>internal/router/router.go
```go
// 添加雇员接口
		emp.POST("", admin.SaveEmp)
```
#### router层

我们需要在router层添加分页查询方法。
>internal/router/admin/employee_router.go
```go
// 分页查询
// @Summary 分页查询
// @Accept application/json
// @Produce application/json
// @router /admin/employee/page [get]
func PageEmp(ctx context.Context, c *app.RequestContext) {
	var page dto.EmployeePageQueryDTO
	// 参数绑定转化为结构体
	err := c.Bind(&page)
	if err != nil {
		log.Println("Employee 参数绑定失败")
		c.JSON(http.StatusBadRequest, common.Result{1, message_c.UNKNOWN_ERROR, nil})
	} else {
		log.Println("员工分页查询，参数为：", page.Name)

		c.JSON(http.StatusOK, common.Result{1, "", service.PageQueryEmp(&page)})
	}
}

```

#### service
我们也需要在service添加一定的逻辑进行数据处理,并返回一个分页查询的结果。
>internal/router/admin/employee_service.go
```go
func PageQueryEmp(page *dto.EmployeePageQueryDTO) *common.PageResult {
	var pageResult = common.PageResult{}

	pageResult.Records, pageResult.Total = db.EmpDao.PageQuery(page)

	return &pageResult
}
```


####  dao层

在 EmployeeDao 中实现 PageQuery 方法：
值得注意的是我们还需要通过name判断是否进行模糊查询。

```go
func (*EmployeeDao) PageQuery(page *dto.EmployeePageQueryDTO) (*[]model.Employee, int64) {
var (
users []model.Employee
count int64
)
origin_sql := DBEngine
// 判断是否含有name，有name不为nil，就进行模糊查询。
if page.Name != nil {
origin_sql = origin_sql.Where("name LIKE ?", "%"+*page.Name+"%").Find(&users)
}
origin_sql.Model(&model.Employee{}).Count(&count)
origin_sql.Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Order("create_time desc").Find(&users)
return &users, count
}
```



### 测试

可以通过接口文档进行测试，也可以进行前后端联调测试。

接下来使用两种方式分别测试：



#### 接口文档测试
打开apifox，添加访问地址 http://localhost:8080/admin/employee/page
然后在params添加以下参数。
![](images/img_12.png)

然后运行程序，进行测试，测试结果如下
```json
{
    "code": 1,
    "msg": "",
    "data": {
        "total": 1,
        "records": [
            {
                "id": 4,
                "name": "xiaozhi",
                "username": "小智",
                "password": "123456",
                "phone": "13812344321",
                "sex": "1",
                "id_number": "",
                "status": 1,
                "create_time": "2024-02-23T16:08:30+08:00",
                "update_time": "2024-02-23T16:08:30+08:00",
                "create_user": 1,
                "update_user": 1
            }
        ]
    }
}
```
如果出现**auth header is empty**的报错，是因为我们的jwt进行了拦截，需要我们登陆获取token，然后粘贴到我们的分页请求中的请求头上。



####  前后端联调测试
登陆
**点击员工管理**
![](images/img_13.png)


**输入员工姓名为zhangsan**
![](images/img_14.png)

不难发现，**最后操作时间格式**不显示，在**代码完善**中解决。

因为不会改前端，暂且没想到较好的解决办法，先搁置改bug。




##  启用禁用员工账号

###  需求分析与设计
####  产品原型

在员工管理列表页面，可以对某个员工账号进行启用或者禁用操作。账号禁用的员工不能登录系统，启用后的员工可以正常登录。如果某个员工账号状态为正常，则按钮显示为 "禁用"，如果员工账号状态为已禁用，则按钮显示为"启用"。

**启禁用员工原型：**

![](images/img_15.png)
**业务规则：**

- 可以对状态为“启用” 的员工账号进行“禁用”操作
- 可以对状态为“禁用”的员工账号进行“启用”操作
- 状态为“禁用”的员工账号不能登录系统



#### 接口设计

![](images/img_16.png)

1). 路径参数携带状态值。

2). 同时，把id传递过去，明确对哪个用户进行操作。

3). 返回数据code状态是必须，其它是非必须。



### 代码开发

#### 修改代码
首先我们要把获取Jwt信息的方法封装成一个函数，因为我们的其他函数可能也需要获取jwt的信息。我们要在jwt.go文件添加以下内容。
>internal/middleware/jwt.go
```javascript
func GetJwtPayload(c *app.RequestContext) int64 {

	jwt_payload, _ := c.Get("JWT_PAYLOAD")
	// 类型转换,我们的数据在claims中是以map[string]interface{}嵌套结构组成的。
	claims := jwt_payload.(jwt.MapClaims)
	origin_emp := claims[IdentityKey].(map[string]interface{})
	emp_id := origin_emp["id"].(float64)
	return int64(emp_id)
}
```
然后也要同步修改employee_router.go的SaveEmp函数中的方法。
>internal/router/admin/employee_router.go
```go
func SaveEmp(ctx context.Context, c *app.RequestContext) {
	var empL model.Employee
	// 参数绑定转化为结构体
	err := c.Bind(&empL)
	if err != nil {
		log.Println("Employee 参数绑定失败")
		c.JSON(http.StatusBadRequest, common.Result{1, message_c.UNKNOWN_ERROR, nil})
	} else {
		// 获取jwt_payload的信息,并把信息赋予empL
		emp_id := middleware.GetJwtPayload(c)
		empL.CreateUser, empL.UpdateUser = emp_id, emp_id
		log.Printf("新增用户:{%s}", empL.Username)
		flag := service.SaveEmp(&empL)
		if flag == true {
			c.JSON(http.StatusOK, common.Result{1, "", nil})
		}
		c.JSON(http.StatusBadRequest, common.Result{1, message_c.ALREADY_EXISTS, nil})
	}
}
```
#### 添加路由
我们首先要在在根路由中添加我们的启用禁用员工账号的路由。
>internal/router/router.go
```go
	// 禁用员工账号
		emp.POST("/status/*status", admin.StartOrStopEmp)
```
####  router层

service模块中，根据接口设计中的请求参数形式对应的在 employee_router.go中创建启用禁用员工账号的方法：
>internal/router/admin/employee_router.go
```go
// 禁用员工账号
// @Summary 禁用员工账号
// @Accept application/json
// @Produce application/json
// @router /admin/employee/status [post]
func StartOrStopEmp(ctx context.Context, c *app.RequestContext) {
	status, id := c.Param("status"), c.Query("id")
	log.Printf("启用禁用员工账号：{%s},{%s}", status, id)
	status_r, _ := strconv.ParseInt(status, 10, 32)
	id_r, _ := strconv.ParseInt(id, 10, 64)
	service.StartOrStopEmp(int32(status_r), id_r, middleware.GetJwtPayload(c))
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}
```


#### service层

在 employee_service.go 中实现启用禁用员工账号的业务方法StartOrStopEmp：
>internal/router/service/employee_service.go
```go
func StartOrStopEmp(status int32, id int64, update_user int64) {
	emp := model.Employee{
		ID:         id,
		Status:     status,
		UpdateUser: update_user,
		UpdateTime: time.Now(),
	}
	db.EmpDao.UpdateStatus(&emp)
}
```



#### dao层

在 employee_dao.go中添加 Update 方法：
>internal/db/employee_dao.go
```go
func (*EmployeeDao) UpdateStatus(emp *model.Employee) {
log.Println(*emp)
DBEngine.Select("status", "update_time", "update_user").Updates(emp)
}

```


### 功能测试

#### 接口文档测试
打开apifox 访问我们的地址 http://localhost:8080/admin/employee/status/0
然后我们在Params中添加参数,同时在herder中添加我们的jwt令牌
![img](images/img_17.png)

**测试前:** 查询employee表中员工账号状态
![img](images/img_18.png)

**开始测试**
![img](images/img_19.png)


**测试完毕后**，再次查询员工账号状态
![img](images/img_20.png)

#### 前后端联调测试

**测试前：**
![img](images/img_21.png)

**点击启用:**
![img](images/img_22.png)


##  编辑员工

###  需求分析与设计

#### 产品原型

在员工管理列表页面点击 "编辑" 按钮，跳转到编辑页面，在编辑页面回显员工信息并进行修改，最后点击 "保存" 按钮完成编辑操作。

**员工列表原型：**

![img](images/img_23.png)

**修改页面原型**：

注：点击修改时，数据应该正常回显到修改页面。

![img](images/img_24.png)


####  接口设计

根据上述原型图分析，编辑员工功能涉及到两个接口：

- 根据id查询员工信息
- 编辑员工信息

**1). 根据id查询员工信息**

![img](images/img_25.png)


**2). 编辑员工信息**
![img](images/img_25.png)

**注:因为是修改功能，请求方式可设置为PUT。**



###  代码开发
####  回显员工信息功能
**1). 添加路由**
在我们的根路由其中添加关于按照用户id查询用户信息
>internal/router/router.go
```go
// 查询雇员接口
		emp.GET("/:id", admin.GetByIdEmp)
```
**2). router逻辑添加**
在 employee_router.go 中创建 GetByIdEmp 方法：
>internal/router/admin/employee_router.go
```go
// 根据id查找雇员
// @Summary 根据id查找雇员
// @Accept application/json
// @Produce application/json
// @router /admin/employee/status [get]
func GetByIdEmp(ctx context.Context, c *app.RequestContext) {
	// 首先从请求url获取id
	id := c.Param("id")
	log.Printf("查询员工账号：{%s}", id)
	id_r, _ := strconv.ParseInt(id, 10, 64)
	emp := service.GetByIdEmp(id_r)
	c.JSON(http.StatusOK, common.Result{1, "", emp})
}
```
**3). 添加service逻辑**
我们需要在我们的employee_service.go添加相应功能。
>internal/router/service/employee_service.go
```go
func GetByIdEmp(id int64) *model.Employee {
	return db.EmpDao.GetById(id)
}

```
**4). dao层**

在 employee_dao.go文件中实现 GetById 方法：
>internal/db/employee_dao.go
```go
func (*EmployeeDao) GetById(id int64) *model.Employee {
	var emp model.Employee
	DBEngine.Where("id=?", id).First(&emp)
	return &emp
}
```

#### 修改员工信息功能

**1). 添加路由**
在我们的根路由其中添加关于按照用户id查询用户信息
>internal/router/router.go
```go
	// 添加修改雇员接口
emp.PUT("", admin.UpdateEmp)
```
**2). router逻辑添加**
在 employee_router.go 中创建 GetByIdEmp 方法：
>internal/router/admin/employee_router.go
```go
// 更新雇员信息
// @Summary 根据id更新雇员信息
// @Accept application/json
// @Produce application/json
// @router /admin/employee [put]
func UpdateEmp(ctx context.Context, c *app.RequestContext) {
var emp model.Employee
c.BindAndValidate(&emp)
emp.UpdateUser,emp.UpdateTime=middleware.GetJwtPayload(c),time.Now()
log.Println("编辑员工信息：", emp)
service.UpdateEmp(&emp)
c.JSON(http.StatusOK, common.Result{1, "", nil})
}

```
**3). 添加service逻辑**
我们需要在我们的employee_service.go添加相应功能。
>internal/router/service/employee_service.go
```go
func UpdateEmp(emp *model.Employee) {
db.EmpDao.Update(emp)
}
```
**4). dao层**
在 employee_dao.go文件中实现 Update 方法：
>internal/db/employee_dao.go
```go
func (*EmployeeDao) Update(emp *model.Employee) {
DBEngine.Updates(emp)
}
```

### 功能测试

#### 接口文档测试

分别测试**根据id查询员工信息**和**编辑员工信息**两个接口

**1). 根据id查询员工信息**

查询employee表中的数据，以id=的记录为例 我们在api 工具访问以下地址，记得加上jwt令牌

![image](images/img_27.png)

开始测试
![image](images/img_28.png)
获取到了id=4的相关员工信息
![image](images/img_29.png)
**2). 编辑员工信息**

修改id=4的员工信息，**name**由**小智**改为**zhangsan**，**username**改为**zhangsan**。
![image](images/img_30.png)

查看employee表数据

![image](images/img_31.png)



#### 前后端联调测试

进入到员工列表查询
![image](images/img_33.png)

对员工姓名为rdesfdf的员工数据修改，点击修改，数据已回显
![image](images/img_32.png)

修改名称为xiaozhi，点击保存，就会数据回显
![image](images/img_34.png)

## 总结
目前为止项目的结构。
```shell
│  go.mod
│  go.sum
│  main.go #程序入口
│  README.md
│      
├─docs # swagger 文档
├─images # README文档中的图片
├─internal
│  ├─config
│  │      config.go
│  │      config.yaml
│  │      
│  ├─db # dao层
│  │      db.go
│  │      employee_dao.go
│  │      
│  ├─middleware # 中间件
│  │      jwt.go
│  │      swagger.go
│  │      
│  ├─models #存储各种类
│  │  ├─common
│  │  │      common.go
│  │  │      
│  │  ├─constant # 定义常量
│  │  │  ├─message_c
│  │  │  │      common.go
│  │  │  │      
│  │  │  └─status_c
│  │  │          common.go
│  │  │          
│  │  ├─dto # 传输类
│  │  │      common.go
│  │  │      
│  │  ├─model #实体类
│  │  │      
│  │  └─vo
│  │          common.go
│  │          
│  ├─router 
│  │  │  router.go
│  │  │  
│  │  ├─admin #router
│  │  │      employee_router.go
│  │  │      
│  │  └─service #service
│  │          employee_service.go
│  │          
│  └─utils
│          gen.go
│          
├─nginx-1.20.2 #前端
├─sql  #数据库文件
│      sky.sql
│      
└─templates
     
```

#  第四章 完善菜品分类功能

## 需求分析与设计

### 产品原型

后台系统中可以管理分类信息，分类包括两种类型，分别是 **菜品分类** 和 **套餐分类** 。

先来分析**菜品分类**相关功能。

**新增菜品分类:** 当我们在后台系统中添加菜品时需要选择一个菜品分类，在移动端也会按照菜品分类来展示对应的菜品。

**菜品分类分页查询:** 系统中的分类很多的时候，如果在一个页面中全部展示出来会显得比较乱，不便于查看，所以一般的系统中都会以分页的方式来展示列表数据。

**根据id删除菜品分类:** 在分类管理列表页面，可以对某个分类进行删除操作。需要注意的是当分类关联了菜品或者套餐时，此分类不允许删除。

**修改菜品分类:** 在分类管理列表页面点击修改按钮，弹出修改窗口，在修改窗口回显分类信息并进行修改，最后点击确定按钮完成修改操作。

**启用禁用菜品分类:** 在分类管理列表页面，可以对某个分类进行启用或者禁用操作。

**分类类型查询:** 当点击分类类型下拉框时，从数据库中查询所有的菜品分类数据进行展示。



**分类管理原型:**

![image](images/img_35.png)

**业务规则：**

- 分类名称必须是唯一的
- 分类按照类型可以分为菜品分类和套餐分类
- 新添加的分类状态默认为“禁用”



###  接口设计

根据上述原型图分析，菜品分类模块共涉及6个接口。

- 新增分类
- 分类分页查询
- 根据id删除分类
- 修改分类
- 启用禁用分类
- 根据类型查询分类

接下来，详细地分析每个接口。


**1). 新增分类**
![image](images/img_36.png)

**2). 分类分页查询**
![image](images/img_37.png)
**3). 根据id删除分类**
![image](images/img_38.png)
**4). 修改分类**
![image](images/img_39.png)
**5). 启用禁用分类**
![image](images/img_40.png)
**6). 根据类型查询分类**
![image](images/img_41.png)
### 表设计

**category表结构:**

| **字段名**  | **数据类型** | **说明**     | **备注**            |
| ----------- | ------------ | ------------ | ------------------- |
| id          | bigint       | 主键         | 自增                |
| name        | varchar(32)  | 分类名称     | 唯一                |
| type        | int          | 分类类型     | 1菜品分类 2套餐分类 |
| sort        | int          | 排序字段     | 用于分类数据的排序  |
| status      | int          | 状态         | 1启用 0禁用         |
| create_time | datetime     | 创建时间     |                     |
| update_time | datetime     | 最后修改时间 |                     |
| create_user | bigint       | 创建人id     |                     |
| update_user | bigint       | 最后修改人id |                     |

同时我们要在model文件夹下新建category.go文件，并且在该文件下添加以下内容。
>internal/models/model/category.go
```go
package model

import (
	"time"
)

const TableNameCategory = "category"

// Category 菜品及套餐分类
type Category struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	Type       int32     `gorm:"column:type;comment:类型   1 菜品分类 2 套餐分类" json:"type"`           // 类型   1 菜品分类 2 套餐分类
	Name       string    `gorm:"column:name;not null;comment:分类名称" json:"name"`                // 分类名称
	Sort       int32     `gorm:"column:sort;not null;comment:顺序" json:"sort"`                  // 顺序
	Status     int32     `gorm:"column:status;comment:分类状态 0:禁用，1:启用" json:"status"`           // 分类状态 0:禁用，1:启用
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`           // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`           // 更新时间
	CreateUser int64     `gorm:"column:create_user;comment:创建人" json:"create_user"`            // 创建人
	UpdateUser int64     `gorm:"column:update_user;comment:修改人" json:"update_user"`            // 修改人
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}

```
## 新增分类
接口详情
![image](images/img_36.png)

###  代码开发
**1). 添加路由**
在我们的根路由其中添加新增菜品的路由
>internal/router/router.go
```go
category := adm.Group("/category")
{
// 新增菜品路由
category.POST("", admin.SaveCategory)
}
```
**2). router逻辑添加**
在 category_router.go 中创建 SaveCategory 方法：
>internal/router/admin/category_router.go
```go
// 新增菜品
// @Summary 新增菜品
// @Accept application/json
// @Produce application/json
// @router /admin/category [post]
func SaveCategory(ctx context.Context, c *app.RequestContext) {
var category model.Category
c.Bind(&category)
// 赋予创建用户和更新用户的数据
category.CreateUser, category.UpdateUser = middleware.GetJwtPayload(c), middleware.GetJwtPayload(c)
// 赋予创建时间和更新时间数据
category.CreateTime, category.UpdateTime = time.Now(), time.Now()
log.Println("新增分类：", category)
service.SaveCategory(&category)
c.JSON(http.StatusOK, common.Result{1, "", nil})
}

```
**3). 添加service逻辑**
我们需要在我们的category_service.go添加相应功能。
>internal/router/service/category_service.go
```go
func SaveCategory(category *model.Category) {
db.CatDao.Save(category)
}
```
**4). dao层**

在 category_dao.go文件中实现 Save 方法：
>internal/db/category_dao.go
```go
func (*CategoryDao) Save(category *model.Category) {
DBEngine.Create(category)
}
```
### 功能测试

#### 接口文档测试

测试**添加菜品分类功能**


**添加菜品分类功能**
查看category表数据
![image](images/img_44.png)

我们再api工具在 访问 http://localhost:8080/admin/category 添加jwt令牌，同时在body添加以下内容。
```json
{
  "name": "十分感动",
  "type": 1,
  "sort": 3
}
```
![image](images/img_45.png)
运行程序进行测试。

查看employee表数据

![image](images/img_46.png)
#### 前后端联调测试
前后端联调启动有问题，具体解决方式可以参考，[这里的问题二](#代码存在的问题)



## 分类分页查询
接口详情
![image](images/img_37.png) 

###  代码开发
####  菜品分类分页
**1). 添加路由**
在我们的根路由其中添加新增菜品的路由
>internal/router/router.go
```go
// 菜品分类分页
category.GET("/page", admin.PageCat)
```
**2). router逻辑添加**
在 category_router.go 中创建 PageCat 方法：
>internal/router/admin/category_router.go
```go
// 菜品分类分页
// @Summary 新增菜品
// @Accept application/json
// @Produce application/json
// @router /admin/category/page [get]
func PageCat(ctx context.Context, c *app.RequestContext) {
var categoryPage dto.CategoryPageQueryDTO
c.Bind(&categoryPage)
log.Println("菜品分类查询", categoryPage)
cat := service.PageQueryDat(&categoryPage)
c.JSON(http.StatusOK, common.Result{1, "", cat})

}

```
**3). 添加service逻辑**
我们需要在我们的category_service.go添加相应功能。
>internal/router/service/category_service.go
```go
func PageQueryDat(categoryPage *dto.CategoryPageQueryDTO) *common.PageResult {
var pageResult = common.PageResult{}
pageResult.Records, pageResult.Total = db.CatDao.PageQuery(categoryPage)
return &pageResult
}

```
**4). dao层**

在 category_dao.go文件中实现 PageQuery 方法：
>internal/db/category_dao.go
```go
func (*CategoryDao) PageQuery(page *dto.CategoryPageQueryDTO) (*[]model.Category, int64) {
var (
cat   []model.Category
count int64
)
origin_sql := DBEngine
// 判断是否含有name，有name不为nil，就进行模糊查询。
if page.Name != nil {
origin_sql = origin_sql.Where("name LIKE ?", "%"+*page.Name+"%")
}
if page.Type != nil {
origin_sql = origin_sql.Where("type=?", page.Type)
}
origin_sql.Model(&model.Category{}).Count(&count)
origin_sql.Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Order("create_time desc").Find(&cat)
return &cat, count
}
```

### 功能测试

#### 接口文档测试

测试**添加菜品分类分页查询功能**

我们再api工具在 访问 http://localhost:8080/admin/category/page/?page=1&pageSize=10&type=2 添加jwt令牌。
![image](images/img_47.png)
运行程序进行测试。
调试结果
```shell
{
    "code": 1,
    "msg": "",
    "data": {
        "total": 2,
        "records": [
            {
                "id": 15,
                "type": 2,
                "name": "商务套餐",
                "sort": 13,
                "status": 1,
                "create_time": "2022-06-09T22:14:10+08:00",
                "update_time": "2022-06-10T11:04:48+08:00",
                "create_user": 1,
                "update_user": 1
            },
            {
                "id": 13,
                "type": 2,
                "name": "人气套餐",
                "sort": 12,
                "status": 1,
                "create_time": "2022-06-09T22:11:38+08:00",
                "update_time": "2022-06-10T11:04:40+08:00",
                "create_user": 1,
                "update_user": 1
            }
        ]
    }
}
```



#### 前后端联调测试
登陆，然后点击分类管理，效果如下。
![image](images/img_48.png)

## 根据id删除分类
![image](images/img_38.png)


###  代码开发
####  根据id删除
**1). 添加路由**
在我们的根路由其中添加新增菜品的路由
>internal/router/router.go
```go
// 添加按照id删除
		category.DELETE("", admin.DeleteCat)
```
**2). router逻辑添加**
在 category_router.go 中创建 PageCat 方法：
>internal/router/admin/category_router.go
```go
// 按照id删除菜品分类
// @Summary 新增菜品
// @Accept application/json
// @Produce application/json
// @router /admin/category [delete]
func DeleteCat(ctx context.Context, c *app.RequestContext) {
id := c.Query("id")
log.Printf("查询员工账号：{%s}", id)
id_r, _ := strconv.ParseInt(id, 10, 64)
if err := service.DeleteCat(&id_r); err != nil {
log.Println(err)
c.JSON(http.StatusOK, common.Result{0, "", nil})
} else {
c.JSON(http.StatusOK, common.Result{1, "", nil})
}
}

```
**3). 添加service逻辑**
我们需要在我们的category_service.go添加相应功能。
>internal/router/service/category_service.go
```go
func DeleteCat(id *int64)  *error {
err:=db.CatDao.Delete(id)
if err!=nil {
return err
}
return nil
}
```
**4). dao层**

在 category_dao.go文件中实现 PageQuery 方法：
>internal/db/category_dao.go
```go
func (*CategoryDao) Delete(id *int64) *error {
err := DBEngine.Delete(&model.Category{}, id).Error
if err != nil {
return &err
}
return nil
}
```

### 功能测试

#### 接口文档测试

测试**按照id删除菜品分类功能**
首先查看我们的的表格
![image](images/img_49.png)

我们再api工具在 访问 http://localhost:8080/admin/category?id=42 添加jwt令牌。
![image](images/img_50.png)

运行程序进行测试。
调试结果
![image](images/img_51.png)
我们可以看到id为42的信息已被删除。

#### 前后端联调测试
登陆，然后点击分类管理，然后点击删除。
![image](images/img_52.png)
可以看到刚才点击删除的条目已经被删除。
![image](images/img_53.png)

## 修改菜品分类信息
接口信息
![image](images/img_39.png)
###  代码开发

**1). 添加路由**
在我们的根路由其中添加新增菜品的路由
>internal/router/router.go
```go
		// 添加修改菜品分类
category.PUT("", admin.UpdateCat)
```
**2). router逻辑添加**
在 category_router.go 中创建 PageCat 方法：
>internal/router/admin/category_router.go
```go
// 修改菜品分类信息
// @Summary 修改菜品分类信息
// @Accept application/json
// @Produce application/json
// @router /admin/category [put]
func UpdateCat(ctx context.Context, c *app.RequestContext) {
var category model.Category
c.Bind(&category)
// 赋予创建时间和更新时间数据
category.CreateTime, category.UpdateTime = time.Now(), time.Now()
log.Println("修改菜品分类信息：", category)
service.UpdateCategory(&category)
c.JSON(http.StatusOK, common.Result{1, "", nil})
}

```
**3). 添加service逻辑**
我们需要在我们的category_service.go添加相应功能。
>internal/router/service/category_service.go
```go
func UpdateCategory(cat *model.Category) {
db.CatDao.Update(cat)
}
```
**4). dao层**

在 category_dao.go文件中实现 PageQuery 方法：
>internal/db/category_dao.go
```go
func (*CategoryDao) Update(category *model.Category) {
DBEngine.Updates(category)
}

```

### 功能测试

#### 接口文档测试

测试**修改菜品分类功能**
首先查看数据库
![image](images/img_54.png)
我们要把id为43的条目名字改为十分
我们再api工具在 访问 http://localhost:8080/admin/category 添加jwt令牌。
并在请求体添加json格式的数据。
```json
{
  "id": 43,
  "name": "十分",
  "sort": 3
}
```
![image](images/img_56.png)
运行程序进行测试。
调试结果
![image](images/img_55.png)
由上图看到我们的修改已经成功

#### 前后端联调测试
登陆，然后点击分类管理，点击修改，我们把名称修改为abc。
![image](images/img_57.png)
修改结果
![image](images/img_58.png)

## 启用禁用分类
接口信息
![image](images/img_41.png)
###  代码开发

**1). 添加路由**
在我们的根路由其中添加新增菜品的路由
>internal/router/router.go
```go
	// 根据类型查询分类
category.GET("/list", admin.ListCat)
```
**2). router逻辑添加**
在 category_router.go 中创建 PageCat 方法：
>internal/router/admin/category_router.go
```go
// 启用禁用菜品分类
// @Summary 启用禁用菜品分类
// @Accept application/json
// @Produce application/json
// @router /admin/category/list [get]
func ListCat(ctx context.Context, c *app.RequestContext) {
ty_pe := c.Query("type")
log.Printf("启用禁用员工账号：{%s}", ty_pe)
tp, _ := strconv.ParseInt(ty_pe, 10, 64)
service.ListCat(&tp)
c.JSON(http.StatusOK, common.Result{1, "", nil})
}
```
**3). 添加service逻辑**
我们需要在我们的category_service.go添加相应功能。
>internal/router/service/category_service.go
```go
func ListCat(tp *int64) *[]model.Category {
return db.CatDao.List(tp)
}
```
**4). dao层**

在 category_dao.go文件中实现 List 方法：
>internal/db/category_dao.go
```go
func (*CategoryDao) List(tp *int64) *[]model.Category {
var cat []model.Category
DBEngine.Where("type=?", tp).Find(&cat)
return &cat
}

```

### 功能测试

#### 接口文档测试

测试**添加菜品分类分页查询功能**

我们再api工具在 访问 http://localhost:8080/admin/category/page/?page=1&pageSize=10&type=2 添加jwt令牌。
![image](images/img_47.png)
运行程序进行测试。
调试结果
```shell
{
    "code": 1,
    "msg": "",
    "data": {
        "total": 2,
        "records": [
            {
                "id": 15,
                "type": 2,
                "name": "商务套餐",
                "sort": 13,
                "status": 1,
                "create_time": "2022-06-09T22:14:10+08:00",
                "update_time": "2022-06-10T11:04:48+08:00",
                "create_user": 1,
                "update_user": 1
            },
            {
                "id": 13,
                "type": 2,
                "name": "人气套餐",
                "sort": 12,
                "status": 1,
                "create_time": "2022-06-09T22:11:38+08:00",
                "update_time": "2022-06-10T11:04:40+08:00",
                "create_user": 1,
                "update_user": 1
            }
        ]
    }
}
```



#### 前后端联调测试
登陆，然后点击分类管理，效果如下。
![image](images/img_48.png)


## 根据类型查询分类
接口信息
![image](images/img_41.png)
###  代码开发

**1). 添加路由**
在我们的根路由其中添加新增菜品的路由
>internal/router/router.go
```go
		// 启用禁用分类
emp.POST("/status/*status", admin.StartOrStopCat)
```
**2). router逻辑添加**
在 category_router.go 中创建 PageCat 方法：
>internal/router/admin/category_router.go
```go
// 启用禁用菜品分类
// @Summary 启用禁用菜品分类
// @Accept application/json
// @Produce application/json
// @router /admin/categorystatus [post]
func StartOrStopCat(ctx context.Context, c *app.RequestContext) {
status, id := c.Param("status"), c.Query("id")
log.Printf("启用禁用员工账号：{%s},{%s}", status, id)
status_r, _ := strconv.ParseInt(status, 10, 32)
id_r, _ := strconv.ParseInt(id, 10, 64)
service.StartOrStopCat(int32(status_r), id_r, middleware.GetJwtPayload(c))
c.JSON(http.StatusOK, common.Result{1, "", nil})
}
```
**3). 添加service逻辑**
我们需要在我们的category_service.go添加相应功能。
>internal/router/service/category_service.go
```go
func StartOrStopCat(status int32, id int64, update_user int64) {
cat := model.Category{
ID:         id,
Status:     status,
UpdateUser: update_user,
UpdateTime: time.Now(),
}
db.CatDao.UpdateStatus(&cat)
}
```
**4). dao层**

在 category_dao.go文件中实现 PageQuery 方法：
>internal/db/category_dao.go
```go
func (*CategoryDao) UpdateStatus(cat *model.Category) {
DBEngine.Select("status", "update_time", "update_user").Updates(cat)
}
```

### 功能测试

#### 接口文档测试

测试**添加菜品分类分页查询功能**

我们再api工具在 访问 http://localhost:8080/admin/category/page/?page=1&pageSize=10&type=2 添加jwt令牌。
![image](images/img_47.png)
运行程序进行测试。
调试结果
```shell
{
    "code": 1,
    "msg": "",
    "data": {
        "total": 2,
        "records": [
            {
                "id": 15,
                "type": 2,
                "name": "商务套餐",
                "sort": 13,
                "status": 1,
                "create_time": "2022-06-09T22:14:10+08:00",
                "update_time": "2022-06-10T11:04:48+08:00",
                "create_user": 1,
                "update_user": 1
            },
            {
                "id": 13,
                "type": 2,
                "name": "人气套餐",
                "sort": 12,
                "status": 1,
                "create_time": "2022-06-09T22:11:38+08:00",
                "update_time": "2022-06-10T11:04:40+08:00",
                "create_user": 1,
                "update_user": 1
            }
        ]
    }
}
```

#### 前后端联调测试
登陆，然后点击分类管理，效果如下。
![image](images/img_48.png)

#  第五章 晚上菜品功能
##  新增菜品
###  需求分析与设计
####  产品原型
后台系统中可以管理菜品信息，通过 **新增功能**来添加一个新的菜品，在添加菜品时需要选择当前菜品所属的菜品分类，并且需要上传菜品图片。
**新增菜品原型：**
![image](images/img_59.png)

当填写完表单信息, 点击"保存"按钮后, 会提交该表单的数据到服务端, 在服务端中需要接受数据, 然后将数据保存至数据库中。



**业务规则：**

- 菜品名称必须是唯一的
- 菜品必须属于某个分类下，不能单独存在
- 新增菜品时可以根据情况选择菜品的口味
- 每个菜品必须对应一张图片



#### 接口设计

根据上述原型图先**粗粒度**设计接口，共包含3个接口。

**接口设计：**

- 根据类型查询分类（已完成）
- 文件上传
- 新增菜品



接下来**细粒度**分析每个接口，明确每个接口的请求方式、请求路径、传入参数和返回值。

**1. 根据类型查询分类**
![image](images/img_60.png)

**2. 文件上传**
![image](images/img_61.png)


**3. 新增菜品**
![image](images/img_62.png)


####  表设计

通过原型图进行分析：
![image](images/img_63.png)

新增菜品，其实就是将新增页面录入的菜品信息插入到dish表，如果添加了口味做法，还需要向dish_flavor表插入数据。所以在新增菜品时，涉及到两个表：

| 表名        | 说明       |
| ----------- | ---------- |
| dish        | 菜品表     |
| dish_flavor | 菜品口味表 |



**1). 菜品表:dish**

| **字段名**  | **数据类型**  | **说明**     | **备注**    |
| ----------- | ------------- | ------------ | ----------- |
| id          | bigint        | 主键         | 自增        |
| name        | varchar(32)   | 菜品名称     | 唯一        |
| category_id | bigint        | 分类id       | 逻辑外键    |
| price       | decimal(10,2) | 菜品价格     |             |
| image       | varchar(255)  | 图片路径     |             |
| description | varchar(255)  | 菜品描述     |             |
| status      | int           | 售卖状态     | 1起售 0停售 |
| create_time | datetime      | 创建时间     |             |
| update_time | datetime      | 最后修改时间 |             |
| create_user | bigint        | 创建人id     |             |
| update_user | bigint        | 最后修改人id |             |

**2). 菜品口味表:dish_flavor**

| **字段名** | **数据类型** | **说明** | **备注** |
| ---------- | ------------ | -------- | -------- |
| id         | bigint       | 主键     | 自增     |
| dish_id    | bigint       | 菜品id   | 逻辑外键 |
| name       | varchar(32)  | 口味名称 |          |
| value      | varchar(255) | 口味值   |          |



###  代码开发

####  文件上传实现

因为在新增菜品时，需要上传菜品对应的图片(文件)，包括后绪其它功能也会使用到文件上传，故要实现通用的文件上传接口。

文件上传，是指将本地图片、视频、音频等文件上传到服务器上，可以供其他用户浏览或下载的过程。文件上传在项目中应用非常广泛，我们经常发抖音、发朋友圈都用到了文件上传功能。

实现文件上传服务，需要有存储的支持，那么我们的解决方案将以下几种：

1. 直接将图片保存到服务的硬盘（
    1. 优点：开发便捷，成本低
    2. 缺点：扩容困难
2. 使用分布式文件系统进行存储
    1. 优点：容易实现扩容
    2. 缺点：开发复杂度稍大（有成熟的产品可以使用，比如：FastDFS,MinIO）
3. 使用第三方的存储服务（例如OSS）
    1. 优点：开发简单，拥有强大功能，免维护
    2. 缺点：付费

在本项目选用Minio服务进行文件存储。


**实现步骤：**

**1). 定义Minio相关配置,生成一个配置对象**
我们在internal文件夹下创建minio文件夹,然后文件夹创建,config.go文件,并生成一个
```go
package main

import (
	"github.com/minio/minio-go"
	"log"
)

var (
	MinioClient *minio.Client
)

const (
	endpoint        = "121.37.143.160:9000" //兼容对象存储服务endpoint,也可以设置自己的服务器地址
	accessKeyID     = "minioadmin"          // 对象存储的Access key
	secretAccessKey = "minioadmin"          /// 对象存储的Secret key
	ssl             = false                 //true代表使用HTTPS
)

func init() {
	// 初使化minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, ssl)
	if err != nil {
		log.Println(err)
	} else {
		MinioClient = minioClient
	}
}
func main() {
	if MinioClient != nil {
		log.Println("链接服务器成功")
	}
}

```
**2). 生成防腐层**



其中，AliOssUtil.java已在sky-common模块中定义

```go

```



####  新增菜品实现

**1). 设计DTO类**

在sky-pojo模块中

```java

```



**2). Controller层**

进入到sky-server模块

```java
package com.sky.controller.admin;

/**
 * 菜品管理
 */
@RestController
@RequestMapping("/admin/dish")
@Api(tags = "菜品相关接口")
@Slf4j
public class DishController {

    @Autowired
    private DishService dishService;

    /**
     * 新增菜品
     *
     * @param dishDTO
     * @return
     */
    @PostMapping
    @ApiOperation("新增菜品")
    public Result save(@RequestBody DishDTO dishDTO) {
        log.info("新增菜品：{}", dishDTO);
        dishService.saveWithFlavor(dishDTO);//后绪步骤开发
        return Result.success();
    }
}
```



**3). Service层接口**

```java
package com.sky.service;

import com.sky.dto.DishDTO;
import com.sky.entity.Dish;

public interface DishService {

    /**
     * 新增菜品和对应的口味
     *
     * @param dishDTO
     */
    public void saveWithFlavor(DishDTO dishDTO);

}
```



**4). Service层实现类**

```java
package com.sky.service.impl;


@Service
@Slf4j
public class DishServiceImpl implements DishService {

    @Autowired
    private DishMapper dishMapper;
    @Autowired
    private DishFlavorMapper dishFlavorMapper;

    /**
     * 新增菜品和对应的口味
     *
     * @param dishDTO
     */
    @Transactional
    public void saveWithFlavor(DishDTO dishDTO) {

        Dish dish = new Dish();
        BeanUtils.copyProperties(dishDTO, dish);

        //向菜品表插入1条数据
        dishMapper.insert(dish);//后绪步骤实现

        //获取insert语句生成的主键值
        Long dishId = dish.getId();

        List<DishFlavor> flavors = dishDTO.getFlavors();
        if (flavors != null && flavors.size() > 0) {
            flavors.forEach(dishFlavor -> {
                dishFlavor.setDishId(dishId);
            });
            //向口味表插入n条数据
            dishFlavorMapper.insertBatch(flavors);//后绪步骤实现
        }
    }

}
```



**5). Mapper层**

DishMapper.java中添加

```java
	/**
     * 插入菜品数据
     *
     * @param dish
     */
    @AutoFill(value = OperationType.INSERT)
    void insert(Dish dish);
```

在/resources/mapper中创建DishMapper.xml

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-mapper.dtd" >
<mapper namespace="com.sky.mapper.DishMapper">

    <insert id="insert" useGeneratedKeys="true" keyProperty="id">
        insert into dish (name, category_id, price, image, description, create_time, update_time, create_user,update_user, status)
        values (#{name}, #{categoryId}, #{price}, #{image}, #{description}, #{createTime}, #{updateTime}, #{createUser}, #{updateUser}, #{status})
    </insert>
</mapper>

```

DishFlavorMapper.java

```java
package com.sky.mapper;

import com.sky.entity.DishFlavor;
import java.util.List;

@Mapper
public interface DishFlavorMapper {
    /**
     * 批量插入口味数据
     * @param flavors
     */
    void insertBatch(List<DishFlavor> flavors);

}
```

在/resources/mapper中创建DishFlavorMapper.xml

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-mapper.dtd" >
<mapper namespace="com.sky.mapper.DishFlavorMapper">
    <insert id="insertBatch">
        insert into dish_flavor (dish_id, name, value) VALUES
        <foreach collection="flavors" item="df" separator=",">
            (#{df.dishId},#{df.name},#{df.value})
        </foreach>
    </insert>
</mapper>
```



### 2.3 功能测试

进入到菜品管理--->新建菜品

<img src="assets/image-20221121195440804.png" alt="image-20221121195440804" style="zoom:50%;" /> 

由于没有实现菜品查询功能，所以保存后，暂且在表中查看添加的数据。

dish表：

<img src="assets/image-20221121195737692.png" alt="image-20221121195737692" style="zoom:50%;" /> 

dish_flavor表：

<img src="assets/image-20221121195902555.png" alt="image-20221121195902555" style="zoom:50%;" /> 

测试成功。



### 2.4代码提交

<img src="assets/image-20221121200332933.png" alt="image-20221121200332933" style="zoom:50%;" />  

后续步骤和上述功能代码提交一致，不再赘述。



## 3. 菜品分页查询

### 3.1 需求分析和设计

#### 3.1.1 产品原型

系统中的菜品数据很多的时候，如果在一个页面中全部展示出来会显得比较乱，不便于查看，所以一般的系统中都会以分页的方式来展示列表数据。

**菜品分页原型：**

<img src="assets/image-20221121201552489.png" alt="image-20221121201552489" style="zoom: 67%;" /> 

在菜品列表展示时，除了菜品的基本信息(名称、售价、售卖状态、最后操作时间)外，还有两个字段略微特殊，第一个是图片字段 ，我们从数据库查询出来的仅仅是图片的名字，图片要想在表格中回显展示出来，就需要下载这个图片。第二个是菜品分类，这里展示的是分类名称，而不是分类ID，此时我们就需要根据菜品的分类ID，去分类表中查询分类信息，然后在页面展示。

**业务规则：**

- 根据页码展示菜品信息
- 每页展示10条数据
- 分页查询时可以根据需要输入菜品名称、菜品分类、菜品状态进行查询



#### 3.1.2 接口设计

根据上述原型图，设计出相应的接口。

<img src="assets/image-20221121202019258.png" alt="image-20221121202019258" style="zoom:50%;" /> <img src="assets/image-20221121202033284.png" alt="image-20221121202033284" style="zoom:50%;" />



### 3.2 代码开发

#### 3.2.1 设计DTO类

**根据菜品分页查询接口定义设计对应的DTO：**

在sky-pojo模块中，已定义

```java
package com.sky.dto;

import lombok.Data;
import java.io.Serializable;

@Data
public class DishPageQueryDTO implements Serializable {

    private int page;
    private int pageSize;
    private String name;
    private Integer categoryId; //分类id
    private Integer status; //状态 0表示禁用 1表示启用

}
```



#### 3.2.2 设计VO类

**根据菜品分页查询接口定义设计对应的VO：**

在sky-pojo模块中，已定义

```java
package com.sky.vo;

import com.sky.entity.DishFlavor;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import java.io.Serializable;
import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.List;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class DishVO implements Serializable {

    private Long id;
    //菜品名称
    private String name;
    //菜品分类id
    private Long categoryId;
    //菜品价格
    private BigDecimal price;
    //图片
    private String image;
    //描述信息
    private String description;
    //0 停售 1 起售
    private Integer status;
    //更新时间
    private LocalDateTime updateTime;
    //分类名称
    private String categoryName;
    //菜品关联的口味
    private List<DishFlavor> flavors = new ArrayList<>();
}
```



#### 3.2.3 Controller层

**根据接口定义创建DishController的page分页查询方法：**

```java
	/**
     * 菜品分页查询
     *
     * @param dishPageQueryDTO
     * @return
     */
    @GetMapping("/page")
    @ApiOperation("菜品分页查询")
    public Result<PageResult> page(DishPageQueryDTO dishPageQueryDTO) {
        log.info("菜品分页查询:{}", dishPageQueryDTO);
        PageResult pageResult = dishService.pageQuery(dishPageQueryDTO);//后绪步骤定义
        return Result.success(pageResult);
    }
```



#### 3.2.4 Service层接口

**在 DishService 中扩展分页查询方法：**

```java
	/**
     * 菜品分页查询
     *
     * @param dishPageQueryDTO
     * @return
     */
    PageResult pageQuery(DishPageQueryDTO dishPageQueryDTO);
```



#### 3.2.5 Service层实现类

**在 DishServiceImpl 中实现分页查询方法：**

```java
	/**
     * 菜品分页查询
     *
     * @param dishPageQueryDTO
     * @return
     */
    public PageResult pageQuery(DishPageQueryDTO dishPageQueryDTO) {
        PageHelper.startPage(dishPageQueryDTO.getPage(), dishPageQueryDTO.getPageSize());
        Page<DishVO> page = dishMapper.pageQuery(dishPageQueryDTO);//后绪步骤实现
        return new PageResult(page.getTotal(), page.getResult());
    }
```



#### 3.2.6 Mapper层

**在 DishMapper 接口中声明 pageQuery 方法：**

```java
	/**
     * 菜品分页查询
     *
     * @param dishPageQueryDTO
     * @return
     */
    Page<DishVO> pageQuery(DishPageQueryDTO dishPageQueryDTO);
```

**在 DishMapper.xml 中编写SQL：**

```xml
<select id="pageQuery" resultType="com.sky.vo.DishVO">
        select d.* , c.name as categoryName from dish d left outer join category c on d.category_id = c.id
        <where>
            <if test="name != null">
                and d.name like concat('%',#{name},'%')
            </if>
            <if test="categoryId != null">
                and d.category_id = #{categoryId}
            </if>
            <if test="status != null">
                and d.status = #{status}
            </if>
        </where>
        order by d.create_time desc
</select>
```



### 3.3 功能测试

#### 3.3.1 接口文档测试

**启动服务：**访问http://localhost:8080/doc.html，进入菜品分页查询接口

**注意：**使用admin用户登录重新获取token，防止token失效。

<img src="assets/image-20221121210252403.png" alt="image-20221121210252403" style="zoom:50%;" />  

**点击发送：**

<img src="assets/image-20221121210333489.png" alt="image-20221121210333489" style="zoom: 67%;" /> 



#### 3.3.2 前后端联调测试

启动nginx,访问 http://localhost

**点击菜品管理**

<img src="assets/image-20221121210500188.png" alt="image-20221121210500188" style="zoom:50%;" /> 

数据成功查出。



## 4. 删除菜品

### 4.1 需求分析和设计

#### 4.1.1 产品原型

在菜品列表页面，每个菜品后面对应的操作分别为**修改**、**删除**、**停售**，可通过删除功能完成对菜品及相关的数据进行删除。

**删除菜品原型：**

<img src="assets/image-20221121211236356.png" alt="image-20221121211236356" style="zoom:67%;" /> 



**业务规则：**

- 可以一次删除一个菜品，也可以批量删除菜品
- 起售中的菜品不能删除
- 被套餐关联的菜品不能删除
- 删除菜品后，关联的口味数据也需要删除掉



#### 4.1.2 接口设计

根据上述原型图，设计出相应的接口。

<img src="assets/image-20221121211801121.png" alt="image-20221121211801121" style="zoom:50%;" /> <img src="assets/image-20221121211814429.png" alt="image-20221121211814429" style="zoom:50%;" />

**注意：**删除一个菜品和批量删除菜品共用一个接口，故ids可包含多个菜品id,之间用逗号分隔。



#### 4.1.3 表设计

在进行删除菜品操作时，会涉及到以下三张表。

<img src="assets/image-20221121212436851.png" alt="image-20221121212436851" style="zoom:50%;" /> 

**注意事项：**

- 在dish表中删除菜品基本数据时，同时，也要把关联在dish_flavor表中的数据一块删除。
- setmeal_dish表为菜品和套餐关联的中间表。
- 若删除的菜品数据关联着某个套餐，此时，删除失败。
- 若要删除套餐关联的菜品数据，先解除两者关联，再对菜品进行删除。



### 4.2 代码开发

#### 4.1.2 Controller层

**根据删除菜品的接口定义在DishController中创建方法：**

```java
	/**
     * 菜品批量删除
     *
     * @param ids
     * @return
     */
    @DeleteMapping
    @ApiOperation("菜品批量删除")
    public Result delete(@RequestParam List<Long> ids) {
        log.info("菜品批量删除：{}", ids);
        dishService.deleteBatch(ids);//后绪步骤实现
        return Result.success();
    }
```



#### 4.2.2 Service层接口

**在DishService接口中声明deleteBatch方法：**

```java
	/**
     * 菜品批量删除
     *
     * @param ids
     */
    void deleteBatch(List<Long> ids);
```



#### 4.2.3 Service层实现类

**在DishServiceImpl中实现deleteBatch方法：**

```java
    @Autowired
    private SetmealDishMapper setmealDishMapper;
	/**
     * 菜品批量删除
     *
     * @param ids
     */
    @Transactional//事务
    public void deleteBatch(List<Long> ids) {
        //判断当前菜品是否能够删除---是否存在起售中的菜品？？
        for (Long id : ids) {
            Dish dish = dishMapper.getById(id);//后绪步骤实现
            if (dish.getStatus() == StatusConstant.ENABLE) {
                //当前菜品处于起售中，不能删除
                throw new DeletionNotAllowedException(MessageConstant.DISH_ON_SALE);
            }
        }

        //判断当前菜品是否能够删除---是否被套餐关联了？？
        List<Long> setmealIds = setmealDishMapper.getSetmealIdsByDishIds(ids);
        if (setmealIds != null && setmealIds.size() > 0) {
            //当前菜品被套餐关联了，不能删除
            throw new DeletionNotAllowedException(MessageConstant.DISH_BE_RELATED_BY_SETMEAL);
        }

        //删除菜品表中的菜品数据
        for (Long id : ids) {
            dishMapper.deleteById(id);//后绪步骤实现
            //删除菜品关联的口味数据
            dishFlavorMapper.deleteByDishId(id);//后绪步骤实现
        }
    }
```



#### 4.2.4 Mapper层

**在DishMapper中声明getById方法，并配置SQL：**

```java
	/**
     * 根据主键查询菜品
     *
     * @param id
     * @return
     */
    @Select("select * from dish where id = #{id}")
    Dish getById(Long id);
```

**创建SetmealDishMapper，声明getSetmealIdsByDishIds方法，并在xml文件中编写SQL：**

```java
package com.sky.mapper;

import com.sky.entity.SetmealDish;
import org.apache.ibatis.annotations.Delete;
import org.apache.ibatis.annotations.Mapper;
import java.util.List;

@Mapper
public interface SetmealDishMapper {
    /**
     * 根据菜品id查询对应的套餐id
     *
     * @param dishIds
     * @return
     */
    //select setmeal_id from setmeal_dish where dish_id in (1,2,3,4)
    List<Long> getSetmealIdsByDishIds(List<Long> dishIds);
}

```

SetmealDishMapper.xml

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-mapper.dtd" >
<mapper namespace="com.sky.mapper.SetmealDishMapper">
    <select id="getSetmealIdsByDishIds" resultType="java.lang.Long">
        select setmeal_id from setmeal_dish where dish_id in
        <foreach collection="dishIds" item="dishId" separator="," open="(" close=")">
            #{dishId}
        </foreach>
    </select>
</mapper>
```

**在DishMapper.java中声明deleteById方法并配置SQL：**

```java
	/**
     * 根据主键删除菜品数据
     *
     * @param id
     */
    @Delete("delete from dish where id = #{id}")
    void deleteById(Long id);
```

**在DishFlavorMapper中声明deleteByDishId方法并配置SQL：**

```java
    /**
     * 根据菜品id删除对应的口味数据
     * @param dishId
     */
    @Delete("delete from dish_flavor where dish_id = #{dishId}")
    void deleteByDishId(Long dishId);
```



### 4.3 功能测试

既可以通过Swagger接口文档进行测试，也可以通过前后端联调测试，接下来，我们直接使用**前后端联调测试**。

进入到菜品列表查询页面

<img src="assets/image-20221122125332084.png" alt="image-20221122125332084" style="zoom:50%;" /> 

对测试菜品进行删除操作

<img src="assets/image-20221122125625014.png" alt="image-20221122125625014" style="zoom:50%;" /> 

同时，进到dish表和dish_flavor两个表查看**测试菜品**的相关数据都已被成功删除。



再次，删除状态为启售的菜品

<img src="assets/image-20221122125841464.png" alt="image-20221122125841464" style="zoom:50%;" /> 

点击批量删除

<img src="assets/image-20221122130016566.png" alt="image-20221122130016566" style="zoom:50%;" /> 

删除失败，因为起售中的菜品不能删除。

### 4.4 代码提交

<img src="assets/image-20221122130426605.png" alt="image-20221122130426605" style="zoom:50%;" /> 

后续步骤和上述功能代码提交一致，不再赘述。



## 5. 修改菜品

### 5.1 需求分析和设计

#### 5.1.1 产品原型

在菜品管理列表页面点击修改按钮，跳转到修改菜品页面，在修改页面回显菜品相关信息并进行修改，最后点击保存按钮完成修改操作。

**修改菜品原型：**

<img src="assets/image-20221122130837173.png" alt="image-20221122130837173" style="zoom:50%;" /> 



#### 5.1.2 接口设计

通过对上述原型图进行分析，该页面共涉及4个接口。

**接口：**

- 根据id查询菜品
- 根据类型查询分类(已实现)
- 文件上传(已实现)
- 修改菜品

我们只需要实现**根据id查询菜品**和**修改菜品**两个接口，接下来，我们来重点分析这两个接口。

**1). 根据id查询菜品**

<img src="assets/image-20221122131733147.png" alt="image-20221122131733147" style="zoom:50%;" /><img src="assets/image-20221122131743509.png" alt="image-20221122131743509" style="zoom:50%;" />

**2). 修改菜品**

<img src="assets/image-20221122131837393.png" alt="image-20221122131837393" style="zoom:50%;" /> <img src="assets/image-20221122131847583.png" alt="image-20221122131847583" style="zoom:50%;" />



<img src="assets/image-20221122131914533.png" alt="image-20221122131914533" style="zoom:50%;" /> 

**注:因为是修改功能，请求方式可设置为PUT。**



### 5.2 代码开发

#### 5.2.1 根据id查询菜品实现

**1). Controller层**

**根据id查询菜品的接口定义在DishController中创建方法：**

```java
    /**
     * 根据id查询菜品
     *
     * @param id
     * @return
     */
    @GetMapping("/{id}")
    @ApiOperation("根据id查询菜品")
    public Result<DishVO> getById(@PathVariable Long id) {
        log.info("根据id查询菜品：{}", id);
        DishVO dishVO = dishService.getByIdWithFlavor(id);//后绪步骤实现
        return Result.success(dishVO);
    }
```



**2). Service层接口**

**在DishService接口中声明getByIdWithFlavor方法：**

```java
	/**
     * 根据id查询菜品和对应的口味数据
     *
     * @param id
     * @return
     */
    DishVO getByIdWithFlavor(Long id);
```



**3). Service层实现类**

**在DishServiceImpl中实现getByIdWithFlavor方法：**

```java
	/**
     * 根据id查询菜品和对应的口味数据
     *
     * @param id
     * @return
     */
    public DishVO getByIdWithFlavor(Long id) {
        //根据id查询菜品数据
        Dish dish = dishMapper.getById(id);

        //根据菜品id查询口味数据
        List<DishFlavor> dishFlavors = dishFlavorMapper.getByDishId(id);//后绪步骤实现

        //将查询到的数据封装到VO
        DishVO dishVO = new DishVO();
        BeanUtils.copyProperties(dish, dishVO);
        dishVO.setFlavors(dishFlavors);

        return dishVO;
    }
```



**4). Mapper层**

**在DishFlavorMapper中声明getByDishId方法，并配置SQL：**

```java
    /**
     * 根据菜品id查询对应的口味数据
     * @param dishId
     * @return
     */
    @Select("select * from dish_flavor where dish_id = #{dishId}")
    List<DishFlavor> getByDishId(Long dishId);
```



#### 5.2.1 修改菜品实现

**1). Controller层**

**根据修改菜品的接口定义在DishController中创建方法：**

```java
	/**
     * 修改菜品
     *
     * @param dishDTO
     * @return
     */
    @PutMapping
    @ApiOperation("修改菜品")
    public Result update(@RequestBody DishDTO dishDTO) {
        log.info("修改菜品：{}", dishDTO);
        dishService.updateWithFlavor(dishDTO);
        return Result.success();
    }
```



**2). Service层接口**

**在DishService接口中声明updateWithFlavor方法：**

```java
	/**
     * 根据id修改菜品基本信息和对应的口味信息
     *
     * @param dishDTO
     */
    void updateWithFlavor(DishDTO dishDTO);
```



**3). Service层实现类**

**在DishServiceImpl中实现updateWithFlavor方法：**

```java
	/**
     * 根据id修改菜品基本信息和对应的口味信息
     *
     * @param dishDTO
     */
    public void updateWithFlavor(DishDTO dishDTO) {
        Dish dish = new Dish();
        BeanUtils.copyProperties(dishDTO, dish);

        //修改菜品表基本信息
        dishMapper.update(dish);

        //删除原有的口味数据
        dishFlavorMapper.deleteByDishId(dishDTO.getId());

        //重新插入口味数据
        List<DishFlavor> flavors = dishDTO.getFlavors();
        if (flavors != null && flavors.size() > 0) {
            flavors.forEach(dishFlavor -> {
                dishFlavor.setDishId(dishDTO.getId());
            });
            //向口味表插入n条数据
            dishFlavorMapper.insertBatch(flavors);
        }
    }
```



**4). Mapper层**

**在DishMapper中，声明update方法：**

```java
	/**
     * 根据id动态修改菜品数据
     *
     * @param dish
     */
    @AutoFill(value = OperationType.UPDATE)
    void update(Dish dish);
```

**并在DishMapper.xml文件中编写SQL:**

```xml
<update id="update">
        update dish
        <set>
            <if test="name != null">name = #{name},</if>
            <if test="categoryId != null">category_id = #{categoryId},</if>
            <if test="price != null">price = #{price},</if>
            <if test="image != null">image = #{image},</if>
            <if test="description != null">description = #{description},</if>
            <if test="status != null">status = #{status},</if>
            <if test="updateTime != null">update_time = #{updateTime},</if>
            <if test="updateUser != null">update_user = #{updateUser},</if>
        </set>
        where id = #{id}
</update>
```



### 5.3 功能测试

本次测试直接通过**前后端联调测试** ，可使用Debug方式启动项目，观察运行中步骤。

进入菜品列表查询页面，对第一个菜品的价格进行修改

<img src="assets/image-20221122141233080.png" alt="image-20221122141233080" style="zoom:50%;" /> 

点击修改，回显成功

<img src="assets/image-20221122141348677.png" alt="image-20221122141348677" style="zoom:50%;" /> 

菜品价格修改后，点击保存

<img src="assets/image-20221122141456498.png" alt="image-20221122141456498" style="zoom:50%;" /> 

修改成功



### 5.4 代码提交

<img src="assets/image-20221122141554380.png" alt="image-20221122141554380" style="zoom:50%;" /> 

后续步骤和上述功能代码提交一致，不再赘述。


<a name="代码存在的问题"></a>
# 代码存在的问题
## 问题一 日期无法正常显示
**示例** 管理员工的界面。
![image](images/img_42.png)
### 解决方法
#### 添加vo视图
我们可以新建一个EmployeeVO视图,将日期格式化成字符串之后在返回给前端。
#### 更改前端
我们需要更改前端的的处理日期的逻辑
## 问题二 菜品种类和套餐种类无法添加
下图两个功能无法使用，具体原因是因为json返回的格式问题。
![image](images/img_43.png)

样例发出的json，其中type和sort都是字符串
```json
{
  "name": "十分感动",
  "type": "1",
  "sort": "3"
}
```
我们需要的json，其中type和sort都是数字
```json
{
  "name": "十分感动",
  "type": 1,
  "sort": 3
}
```
### 解决方法
#### 添加vo视图（不推荐）
我们可以新建一个视图,让视图能够接受字符串类型，之后我们再转换成int类型赋值给需要插入的类，但是不建议这么做，因为go的类型转换还是有点繁琐，建议还是改前端。
#### 更改前端
我们需要更改前端的发出json的逻辑
