package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/logger/accesslog"
	"net/http"
	"reggie/internal/middleware"
	"reggie/internal/router/admin"
	"reggie/internal/router/user"
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
	dish := adm.Group("/dish")
	{
		// 添加菜品
		dish.POST("", admin.SaveDish)
		// 菜品分页查询
		dish.GET("/page", admin.PageDish)
		// 菜品批量删除
		dish.DELETE("", admin.DeleteDish)
		// 按照id查询菜品
		dish.GET("/:id", admin.GetByIdDish)
		// 添加修改菜品
		dish.PUT("", admin.UpdateDish)
		// 启用禁用分类
		dish.POST("/status/*status", admin.StartOrStopDish)
		// 根据类型查询分类
		dish.GET("/list", admin.ListDish)
	}
	// 套餐接口
	meal := adm.Group("/setmeal")
	{
		meal.POST("", admin.SaveSetMealWithDish)
		meal.GET("/page", admin.PageSetMeal)
		meal.GET("/:id", admin.GetByIDDishMeal)
		meal.DELETE("", admin.DeleteBatchMeal)
		meal.PUT("", admin.UpdateMeal)
		meal.POST("/status/:status", admin.StartOrStopMeal)
	}
	shop := adm.Group("/shop")
	{
		shop.PUT("/:status", admin.SetStatusShop)
		shop.GET("/status", admin.GetStatusShop)
	}

	users := r.Group("/user")
	user_jwt := middleware.InitJwtUser()
	us := users.Group("/user")
	us.POST("/login", user_jwt.LoginHandler)
	users.Use(user_jwt.MiddlewareFunc())
	u_category := users.Group("/category")
	{
		u_category.GET("/list", user.ListCategory)
	}
	u_shop := users.Group("/shop")
	{
		u_shop.GET("/status", admin.GetStatusShop)
	}
	u_dish := users.Group("/dish")
	{
		u_dish.GET("/list", user.ListDish)
	}
	u_shoppingCart := users.Group("/shoppingCart")
	{
		u_shoppingCart.POST("/add", user.AddShoppingCart)
		u_shoppingCart.GET("/list", user.ListShoppingCart)
		u_shoppingCart.DELETE("/clean", user.CleanShoppingCart)
	}
	u_address := users.Group("/addressBook")
	{
		u_address.GET("/list", user.ListAddress)
		u_address.POST("", user.SaveAddress)
		u_address.GET("/:id", user.GetByIdAddress)
		u_address.PUT("", user.UpdateAddress)
		u_address.PUT("/default", user.SetDefaultAddress)
		u_address.DELETE("", user.DeleteByIdAddress)
		u_address.GET("/default", user.GetDefaultAddress)
	}
}
