# 📝 Postar

[![license](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![coverage](_icons/coverage.svg)](_icons/coverage.svg)
![build](https://github.com/infra-io/postar/actions/workflows/check.yml/badge.svg)

**Postar** is an easy-to-use and low-coupling email service, which can provide email functions for your applications.

[阅读中文版的 Read me](./README.md)

### 🥇 Features

* Plain and HTML form email supports
* Synchronous/Asynchronous mode supports, and timeout is available in synchronous mode
* Support http/grpc protocol
* Gracefully shutdown with signal mechanism

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

### 🚀 Installation

1. Use Docker (recommend)

See more information in [docker hub](https://hub.docker.com/r/fishgoddess/postar).

2. Use source code

Execute `make build` in the root of source code will generate target directory, which contains all binary files.

_Notice 1: Default config file are `postar.toml` and `postar-admin.toml`, default log file are `postar.log` and `postar-admin.log`._

_Notice 2: How to use client in [Gitee](https://gitee.com/infra-io/postar-client) or [GitHub](https://github.com/infra-io/postar-client)._

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.
