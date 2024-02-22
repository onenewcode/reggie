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
