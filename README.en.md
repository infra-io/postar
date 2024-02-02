# 📝 Postar

[![license](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![coverage](_icons/coverage.svg)](_icons/coverage.svg)
![build](https://github.com/infra-io/postar/actions/workflows/check.yml/badge.svg)

**Postar** is an easy-to-use and low-coupling email service, which can provide email functions for your applications.

[阅读中文版的 Read me](./README.md)

### 🥇 Features

* Plain and Html form email supports
* Synchronous/Asynchronous mode supports, and timeout is available in synchronous mode
* Support http/http2/grpc/vex protocol
* Gracefully shutdown with signal mechanism

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

### 🚀 Installation

1. Use Docker (recommend)

Installation manual: [Gitee](https://gitee.com/infra-io/postar-docker)
/ [GitHub](https://github.com/infra-io/postar-docker) .

Docker Hub: [https://hub.docker.com/r/fishgoddess/postar](https://hub.docker.com/r/fishgoddess/postar).

2. Use source code

Postar has two ways to get binary:

1. Invoking `make build` in the root of source code will generate target directory, which contains all binary files.

2. Building by `go build` (or running by `go run`) in `cmd/postar`, see `go`.

_Notice: Default config file is `./postar.ini`, default log output directory is `./log/postar.log`._

> Want to know how to use? See [_examples](_examples).

> Client: [Gitee](https://gitee.com/infra-io/postar-client) or [GitHub](https://github.com/infra-io/postar-client).

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.                                       |
