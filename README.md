# 📧 Postar

[![license](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![coverage](_icons/coverage.svg)](_icons/coverage.svg)
![Test And Build](https://github.com/infra-io/postar/actions/workflows/check.yml/badge.svg)

**Postar** 是一个简单易用且低耦合的邮件服务，可以为您的应用程序提供邮件功能。

[Read me in English](./README.en.md)

### 🥇 功能特性

* 支持发送文本邮件和 HTML 邮件
* 多业务空间隔离机制，支持账号和模板独立配置
* 支持 http/grpc 等网络协议
* 支持 signal 通知的平滑下线

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

### 🚀 安装方式

* 使用二进制包的方式（推荐）

我们提供了编译好的二进制包文件，到 [码云](https://gitee.com/infra-io/postar/releases) 或 [GitHub](https://github.com/infra-io/postar/releases) 上进行下载。

* 使用源码包的方式

在源码根目录执行 `make build` 会生成 target 目录，所有的二进制包都在里面。

_注意事项 1：默认的配置文件是 [postar.toml](./config/postar.toml) 和 [postar_admin.toml](./config/postar_admin.toml)，默认的日志文件存储在 `log` 目录下。_

_注意事项 2：服务依赖 mysql 组件存储邮件相关配置，需要先在 mysql 上创建对应的库表，具体 sql 可以查看 [postar.sql](./postar.sql)，后续库表变更也会体现在该文件中。_

_注意事项 3：客户端使用请查看 [码云](https://gitee.com/infra-io/postar-client) 或 [GitHub](https://github.com/infra-io/postar-client)。_

### 👥 贡献者

如果您觉得 **postar** 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。
