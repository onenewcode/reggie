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
