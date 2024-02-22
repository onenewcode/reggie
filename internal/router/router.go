package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"reggie/internal/middleware"
)

func InitRouter(r *server.Hertz) {
	myJwt := middleware.InitJwt()
	emp := r.Group("/admin/employee")
	emp.POST("/login", myJwt.LoginHandler)
}
