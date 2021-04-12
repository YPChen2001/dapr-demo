# Dapr demo 快速入门示例程序

在开始前，应该已成功安装了[dapr](https://docs.dapr.io/getting-started/install-dapr-cli/)
此代码基于go语言编写，并依赖dapr-go-sdk，go的版本应该在V1.15以及以上

- clone https://github.com/YPChen2001/dapr-demo.git
- 运行svc1：dapr run --app-id svc1 --app-port 50051 --dapr-http-port 3500 go run main.go
- 运行svc2: dapr run --app-id=svc2  go run main.go
- 使用IDE打开此项目，了解通过dapr是如何服务间调用以及状态管理的