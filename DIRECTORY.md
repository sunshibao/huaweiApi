```
.
├── DIRECTORY.md # 本文件
├── Makefile # 构建文件
├── README.md # 帮助文档
├── builds # 构建目录
│   ├── debug # 本地测试目录
│   │   └── service #本地测试生成物，可以在Makefile中调整目录地址
│   ├── protos # protobuf相关脚本
│   │   └── protos.sh # 生成protobuf文档脚本
│   └── release # 生成Linux执行二进制文件的目录
│       └── service # 发行用二进制
├── cmd # 入口命令目录
│   └── main.go
├── config # 配置文件目录
│   ├── dev.json # 本地测试用配置文件
│   └── example.json # 示例配置文件
├── docs # 帮助文档
│   └── images # 帮助文档图片
│       ├── create_project_1.jpg
│       ├── create_project_2.jpg
│       ├── create_project_3.jpg
│       ├── create_project_4.jpg
│       ├── create_project_5.jpg
│       ├── create_project_6.jpg
│       └── use_for_project_1.jpg
├── go.mod
├── go.sum
└── pkg # 内容包
    ├── application # 项目框架相关
    │   ├── command.go # 客户端命令代码文件
    │   ├── router.go # RESTful路由文件
    │   ├── rpc_register.go # RPC 注册运行文件
    │   └── service.go # 进程启用服务
    ├── config 配置格式结构
    │   ├── duration.go # 时间范围配置格式
    │   └── structure.go # 配置结构体
    ├── constants # 常量包
    │   └── request.go # RESTful请求时的一些常量
    ├── databases # 数据库连接
    │   └── connection.go
    ├── models # 模块存放地址
    ├── restful # RESTful控制器逻辑，对外接口的控制器都存放这个目录中，可以再创建目录
    │   ├── status # 运行状态测试方法
    │   │   └── ping.go
    │   └── swaggerdocs # swagger生成文档与在gin中执行的文档
    │       ├── docs.go
    │       ├── swagger.json
    │       └── swagger.yaml
    ├── rpc # RPC控制器逻辑
    │   ├── protos # protobuf存放目录
    │   │   ├── state.pb.go
    │   │   ├── state.pb.micro.go # go-micro根据protobuf文件生成的go代码
    │   │   └── state.proto # protobuf接口定义文件
    │   └── state # protobuf实现的目录
    │       ├── ping.go
    │       └── state.go
    └── utils # 工具目录
        ├── config 配置文件监控器
        │   └── load_conf.go
        ├── h # RESTful回复内容用的工具类
        │   ├── h.go
        │   └── h_test.go
        ├── idcreator # ID生成器，实现了使用sonyflake
        │   ├── creator.go
        │   └── creator_test.go
        ├── log # 日志相关工具
        │   ├── init.go
        │   └── req_logger.go
        └── version # 版本号相关数据
            └── version.go
```