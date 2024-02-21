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
