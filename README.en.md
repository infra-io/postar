# 📝 postar

[![build](_icons/build.svg)](_icons/build.svg)
[![coverage](_icons/coverage.svg)](_icons/coverage.svg)
[![license](_icons/license.svg)](https://opensource.org/licenses/MIT)

**Postar** is an easy-to-use and low-coupling email service, which can provide email functions for your applications.

[阅读中文版的 Read me](./README.md)

### 🥇 Features

* Plain and Html form email supports
* Synchronous/Asynchronous mode supports, and timeout is available in synchronous mode
* Support http/http2/grpc/vex/udp protocol
* Gracefully shutdown with signal mechanism

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

### 🚀 Installation

1. Use Docker (recommend)

Installation manual: [Gitee](https://gitee.com/avino-plan/postar-docker) / [GitHub](https://github.com/avino-plan/postar-docker) .

Docker Hub: [https://hub.docker.com/r/fishgoddess/postar](https://hub.docker.com/r/fishgoddess/postar).

2. Use source code

Postar has two ways to get binary:

1. Invoking `./build.sh` in the root of source code will generate target directory, which contains all binary files.

2. Building by `go build` (or running by `go run`) in `cmd/postar`, see `go`.

_Notice: Default config file is `./postar.ini`, default log output directory is `./log/service.log`._

> Want to know how to use? See [_examples](_examples).

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects postar used

| Project | Author      | Description                                           | link                                                                                            |
|---------|-------------|-------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| logit   | FishGoddess | A high-performance and easy-to-use logging foundation | [Gitee](https://gitee.com/go-logit/logit) / [GitHub](https://github.com/go-logit/logit)   |
| errors  | FishGoddess | A lib for handling error gracefully in Go             | [Gitee](https://gitee.com/FishGoddess/errors) / [GitHub](https://github.com/FishGoddess/errors) |
| ants | panjf2000 | A high-performance and low-cost goroutine pool   | [GitHub](https://github.com/panjf2000/ants) |
| gomail | alexcesaro | The best way to send emails in Go  | [GitHub](https://github.com/go-gomail/gomail/tree/v2) |
| ini | unknwon | Provides INI file read and write functionality in Go  | [GitHub](https://github.com/go-ini/ini) |
| httprouter | julienschmidt | A high performance HTTP request router  | [GitHub](https://github.com/julienschmidt/httprouter) |
