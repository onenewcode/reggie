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
