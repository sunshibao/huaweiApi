# 配置中心服务框架
此服务主要用于提供配置信息。
# 使用方法

## 1、首先把仓库clone到本地
项目gitlab地址：`ssh://git@dev-gitlab.wanxingrowth.com:8022/wanxin-go-micro/shopping-mall-config-service.git`

## 2、然后执行下面命令
`go mod download`

## 3、编译方法

### 测试运行

```shell script
make run_service_local
```

## 生成RPC（Protobuf）接口

```shell script
make build_service_protos
```

## 生成Linux程序

```shell script
make build_service_cross
```

注释：
1.检查是否安装proto插件，如果没安装，请执行：
`GO111MODULE="on" GOPROXY=https://mirrors.aliyun.com/goproxy/,direct go get github.com/micro/protoc-gen-micro/v2`
2.检查是否安装swag插件，如果没安装，请执行：
`GO111MODULE="on" GOPROXY=https://mirrors.aliyun.com/goproxy/,direct go get github.com/swaggo/swag/cmd/swag@v1.6.5```