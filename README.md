# 📝 postar

[![build](_icons/build.svg)](_icons/build.svg)
[![coverage](_icons/coverage.svg)](_icons/coverage.svg)
[![license](_icons/license.svg)](https://opensource.org/licenses/MIT)

**postar** 是一个简单易用且低耦合的邮件服务，可以为您的应用程序提供邮件功能。

[Read me in English](./README.en.md)

### 🥇 功能特性

* 支持发送文本邮件和 HTML 邮件
* 支持异步邮件发送
* 支持 http 请求调用接口

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

### 🚀 安装方式

> 使用 docker 的方式（推荐）

Docker 版本安装请看：[docker](_examples/install/docker_installation_manual.md) 。

访问 Docker Hub 上的主页：[https://hub.docker.com/r/fishgoddess/postar](https://hub.docker.com/r/fishgoddess/postar) 。

> 使用安装包的方式

**postar 仅提供 windows、linux 和 darwin 三种系统下 amd64 架构的二进制包。**

Windows 版本安装请看：[windows](_examples/install/windows_installation_manual.md) 。

Linux 和 Mac 版本安装请看：[linux_and_mac](_examples/install/linux_and_mac_installation_manual.md) 。

### 📖 使用手册

* 敬请期待

### 🔥 性能测试

> 测试文件：[benchmark_test.go](_examples/test/benchmark_test.go)

| 服务器类型 | 1 秒内运行次数 (越大越好) |  每个操作消耗时间 (越小越好) | B/op (越小越好) | allocs/op (越小越好) |
| -----------|--------|-------------|-------------|-------------|
| http | &nbsp; 3165 | 386013 ns/op | 14838 B/op | 89 allocs/op |
| **jsonrpc** | **17462** | **&nbsp; 69567 ns/op** | **&nbsp; &nbsp; 712 B/op** | **15 allocs/op** |
| grpc | 10000 | 132845 ns/op | &nbsp; 5248 B/op | 98 allocs/op |

> 测试环境：I7-6700HQ CPU @ 2.6 GHZ，16 GB RAM

### 👥 贡献者

如果您觉得 **postar** 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 postar 使用的技术

| 项目 | 作者 | 描述 | 链接 |
| -----------|--------|-------------|-------------------|
| logit | FishGoddess | 一个高性能、功能强大且极易上手的日志库 | [GitHub](https://github.com/FishGoddess/logit) / [码云](https://gitee.com/FishGoddess/logit) |

