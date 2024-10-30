# 玉衡天枢项目模板

![Go Version](https://img.shields.io/github/go-mod/go-version/alioth-center/alioth-app-template)
![Go Report Card](https://goreportcard.com/badge/github.com/alioth-center/alioth-app-template)
![GitHub Actions](https://img.shields.io/github/actions/workflow/status/alioth-center/alioth-app-template/build-docker.yml?branch=main)
![License](https://img.shields.io/github/license/alioth-center/alioth-app-template)

**中文** | [English](./readme.en.md)

---- 


## 特性

1. 基于 gin 框架的 Web 服务模板
2. 已配置好项目架构，包括配置文件、日志、数据库等
3. 包含一个 example 示例接口和实现
4. 配置好了 Dockerfile 和 Github Actions 自动构建

## 使用

1. fork 本项目
2. 修改项目名称和模块路径，编辑配置文件
3. 开始编写你的业务逻辑
4. 运行此项目
5. 调试并测试你的接口
6. 提交你的代码到**你的**仓库
7. 开启 Github Actions 自动构建
8. 等待构建完成，查看构建结果
9. 使用你的镜像并部署你的服务

## 目录说明

```text
.
├── app
│   ├── api
│   ├── entity
│   ├── global
│   ├── router
│   └── service
├── config
├── go.mod
├── go.sum
├── main.go
└── readme.md
```

- `app` 业务逻辑代码
  - `api` 接口定义
  - `entity` 运行实体，如请求、响应等
  - `model` 数据模型(可选)
  - `dao` 数据访问层(可选)
  - `global` 全局变量
  - `router` 路由配置
  - `service` 业务逻辑
- `config` 配置文件
- `main.go` 项目入口

## 默认接口说明

模板内有一个 `example/ping` 接口，用于测试服务是否正常运行。

这是这个接口的请求，接口会原样返回请求的内容。

```shell
curl -X POST http://localhost:10000/example/ping \
  --header "Content-Type: application/json" \
  --data-raw '{"hello": "world"}'
```

## 贡献者

![Contributors](https://contrib.rocks/image?repo=alioth-center/alioth-app-template&max=1000)