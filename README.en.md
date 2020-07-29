# 📝 postar

[![License](_icon/license.svg)](https://opensource.org/licenses/MIT)

**Postar** is an easy-to-use and low-coupling email service, which can provide email functions for your applications.

[阅读中文版的 Read me](./README.md)

### 🥇 Features

* Plain and Html form email supports
* Asynchronous send supports
* HTTP api supports
* JsonRPC and gRPC Supports

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

### 🚀 Installation

> Use docker (recommend)

Installation manual: [docker](_examples/install/docker_installation_manual.md) .

Access the homepage of Docker Hub: [https://hub.docker.com/r/fishgoddess/postar](https://hub.docker.com/r/fishgoddess/postar) .

> Use binary

**postar only provides three versions in amd64: windows, linux and darwin.**

Installation manual: 
[windows](_examples/install/windows_installation_manual.md) .
/
[linux_and_mac](_examples/install/linux_and_mac_installation_manual.md).

### 📖 Guides

* Coming soon

### 🔥 Benchmarks

> Test cases：[test/benchmark_test.go](./test/benchmark_test.go)

| server | times/s (large is better) |  ns/op (small is better) | B/op | allocs/op |
| -----------|--------|-------------|-------------|-------------|
| http | &nbsp; 3165 | 386013 ns/op | 14838 B/op | 89 allocs/op |
| **jsonrpc** | **17462** | **&nbsp; 69567 ns/op** | **&nbsp; &nbsp; 712 B/op** | **15 allocs/op** |
| grpc | 10000 | 132845 ns/op | &nbsp; 5248 B/op | 98 allocs/op |

> Environment：I7-6700HQ CPU @ 2.6 GHZ, 16 GB RAM

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects postar used

| Project | Author | Description | link |
| -----------|--------|-------------|------------------|
| logit | FishGoddess | A high-performance and easy-to-use logging foundation | [GitHub](https://github.com/FishGoddess/logit) / [Gitee](https://gitee.com/FishGoddess/logit) |
