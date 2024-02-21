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
