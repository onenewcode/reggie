# 总体项目结构
```shell
.
├── cmd
│   └── server
│       └── main.go          # 主程序入口，启动 HTTP 服务器
├── internal
│   ├── config
│   │   ├── config.go       # 配置文件读取与解析
│   │   ├── dev.toml        # 开发环境配置文件
│   │   ├── prod.toml       # 生产环境配置文件
│   │   └── test.toml       # 测试环境配置文件
│   ├── db
│   │   ├── db.go           # 数据库连接管理
│   │   ├── model.go        # 数据库模型定义
│   │   └── repository.go   # 数据库操作层（Repository）
│   ├── middleware
│   │   ├── auth.go         # 认证中间件
│   │   └── logging.go      # 日志记录中间件
│   ├── service
│   │   ├── user_service.go # 用户服务逻辑实现
│   │   └── ...             # 其他服务逻辑实现
│   └── util
│       ├── errors.go        # 自定义错误类型与处理
│       ├── logger.go        # 日志工具包
│       └── ...
├── pkg
│   └── rpc                 # 如果有 gRPC 或其他远程过程调用相关的代码
│       ├── proto
│       │   ├── user.proto  # gRPC 定义文件
│       │   └── ...
│       ├── client.go       # gRPC 客户端封装
│       └── server.go       # gRPC 服务端实现
├── router
│   ├── api_v1.go           # API v1 路由定义
│   └── web.go              # 前端路由或其他公共路由
├── public                  # 静态资源目录（HTML, CSS, JS等）
├── templates               # HTML模板文件
├── go.mod                  # Go模块定义文件
├── go.sum                  # 依赖包的校验和信息
└── README.md               # 项目文档与说明
```