package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/logger/accesslog"
	"net/http"
	"reggie/internal/middleware"
	"reggie/internal/router/admin"
)

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
		emp.POST("", admin.SaveEmp)
		// 添加修改雇员接口
		emp.PUT("", admin.UpdateEmp)
		// 查询雇员接口
		emp.GET("/:id", admin.GetByIdEmp)
		// 禁用员工账号
		emp.POST("/status/*status", admin.StartOrStopEmp)
		emp.GET("/page", admin.PageEmp)
		// 这是个测试方法，之后会测试我们的jwt是否拦截
		emp.GET("/test", func(c context.Context, ctx *app.RequestContext) {
			ctx.String(http.StatusOK, "Fds")
		})
	}
	category := adm.Group("/category")
	{
		// 新增菜品路由
		category.POST("", admin.SaveCategory)
		// 菜品分类分页
		category.GET("/page", admin.PageCat)
		// 添加按照id删除
		category.DELETE("", admin.DeleteCat)
		// 添加修改菜品分类
		category.PUT("", admin.UpdateCat)
		// 启用禁用分类
		category.POST("/status/*status", admin.StartOrStopCat)
		// 根据类型查询分类
		category.GET("/list", admin.ListCat)
	}
	com := adm.Group("/common")
	{
		com.POST("/upload", admin.UploadImg)
	}

}
