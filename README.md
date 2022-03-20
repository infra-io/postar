# 📝 Postar

[![build](_icons/build.svg)](_icons/build.svg)
[![coverage](_icons/coverage.svg)](_icons/coverage.svg)
[![license](_icons/license.svg)](https://opensource.org/licenses/MIT)

**Postar** 是一个简单易用且低耦合的邮件服务，可以为您的应用程序提供邮件功能。

[Read me in English](./README.en.md)

### 🥇 功能特性

* 支持发送文本邮件和 HTML 邮件
* 支持同步、异步邮件发送，同步模式可配置超时
* 支持 http/http2/grpc/vex/udp 等网络协议
* 支持 signal 通知的平滑下线

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

### 🚀 安装方式

* 使用 Docker 的方式（推荐）

Docker 版本安装请看：[码云](https://gitee.com/avino-plan/postar-docker) / [GitHub](https://github.com/avino-plan/postar-docker) 。

访问 Docker Hub 上的主页：[https://hub.docker.com/r/fishgoddess/postar](https://hub.docker.com/r/fishgoddess/postar) 。

* 使用源码包的方式

Postar 的二进制执行包可以通过源码进行编译得到，一共有两种方式：

1. 在源码根目录执行 `make build` 会生成 target 目录，所有的二进制包都在里面。

2. 在 `cmd/postar` 目录下使用 `go build` 构建服务（或 `go run` 启动服务），参考 `go` 命令。

_注意：默认的配置文件路径是 `./postar.ini`，默认的日志输出是 `./log/postar.log`。_

> 想知道怎么使用？查看 [_examples](_examples)。

> 客户端：[码云](https://gitee.com/avino-plan/postar-client) 或 [Github](https://github.com/avino-plan/postar-client)。

### 👥 贡献者

如果您觉得 **postar** 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 Postar 使用的技术

| 项目         | 作者            | 描述                     | 链接                                                                                           |
|------------|---------------|------------------------|----------------------------------------------------------------------------------------------|
| logit      | FishGoddess   | 一个高性能、功能强大且极易上手的日志库    | [码云](https://gitee.com/go-logit/logit) / [GitHub](https://github.com/go-logit/logit)         |
| errors     | FishGoddess   | 一个用于优雅地处理 Go 中错误的库     | [码云](https://gitee.com/FishGoddess/errors) / [GitHub](https://github.com/FishGoddess/errors) |
| ants       | panjf2000     | 一个高性能且低损耗的 goroutine 池 | [GitHub](https://github.com/panjf2000/ants)                                                  |
| gomail     | alexcesaro    | 一个用于在 Go 中发送邮件的库       | [GitHub](https://github.com/go-gomail/gomail/tree/v2)                                        |
| ini        | unknwon       | 一个操作 ini 配置的库          | [GitHub](https://github.com/go-ini/ini)                                                      |
| httprouter | julienschmidt | 一个高性能的 http 路由库        | [GitHub](https://github.com/julienschmidt/httprouter)                                        |
