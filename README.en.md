# 📝 postar

[![build](_icons/build.svg)](_icons/build.svg)
[![coverage](_icons/coverage.svg)](_icons/coverage.svg)
[![license](_icons/license.svg)](https://opensource.org/licenses/MIT)

**Postar** is an easy-to-use and low-coupling email service, which can provide email functions for your applications.

[阅读中文版的 Read me](./README.md)

### 🥇 Features

* Plain and Html form email supports
* Asynchronous send supports
* HTTP api supports

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

### 🚀 Installation

1. Use Docker (recommend)

Installation manual: [Gitee](https://gitee.com/avino-plan/postar-docker) / [GitHub](https://github.com/avino-plan/postar-docker) .

Docker Hub: [https://hub.docker.com/r/fishgoddess/postar](https://hub.docker.com/r/fishgoddess/postar).

2. Use source code

Postar has two ways to get binary:

1. Invoking `./build.sh` in the root of source code will generate target directory, which contains all binary files.

2. Build by `go` command, see `go build`.

_Notice: Default config file is `/opt/postar/conf/postar.ini`, default log output directory is `/opt/postar/log/`, and you need them to start service._

> Want to know how to use? See [_examples](_examples).

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects postar used

| Project | Author | Description | link |
| -----------|--------|-------------|------------------|
| logit | FishGoddess | A high-performance and easy-to-use logging foundation | [Gitee](https://gitee.com/FishGoddess/logit) / [GitHub](https://github.com/FishGoddess/logit) |
