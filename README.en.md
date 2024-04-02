# 📝 Postar

[![license](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![coverage](_icons/coverage.svg)](_icons/coverage.svg)
![build](https://github.com/infra-io/postar/actions/workflows/check.yml/badge.svg)

**Postar** is an easy-to-use and low-coupling email service, which can provide email functions for your applications.

[阅读中文版的 Read me](./README.md)

### 🥇 Features

* Plain and HTML form email supports
* Multi spaces supports, accounts and templates configuration supports
* Support http/grpc protocol
* Gracefully shutdown with signal mechanism

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

### 🚀 Installation

1. Use binary (recommend)

We provide built binaries, go check on [码云](https://gitee.com/infra-io/postar/releases) or [GitHub](https://github.com/infra-io/postar/releases) and download.

2. Use source code

Execute `make build` in the root of source code will generate target directory, which contains all binary files.

_Notice 1: Default config file are [`postar.toml`](./config/postar.toml) and [`postar_admin.toml`](./config/postar_admin.toml), default log files are stored in `log`._

_Notice 2: Postar needs mysql to store some configurations about email, so you should create tables in [`postar.sql`](./postar.sql) which will record all changes._

_Notice 3: How to use client in [Gitee](https://gitee.com/infra-io/postar-client) or [GitHub](https://github.com/infra-io/postar-client)._

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.
