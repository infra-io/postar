## ✒ 未来版本的新特性 (Features in future versions)

### v0.2.0
* 支持邮件中携带附件
* 异步邮件发送功能
* 邮件发送信息监控平台
* 添加对 WebSocket 协议的支持
* 支持 ini 配置文件，后续可能考虑更换为 toml 格式
* 多发送者支持，每个发送者之间互不干扰，提高并发数
* dialer 对象池或连接池优化，减少资源占用和连接的耗时
* 简化配置文件的使用，主要是取消日志配置文件，集成进 postar 的主配置文件中

### v0.1.2-alpha
* 增加 JsonRPC 远程调用接口
* 增加 gRPC 远程调用接口
* 增加 UDP 远程调用接口

### v0.1.1-alpha
* 修复 wg.Add 调用位置导致的 Wait 失效问题，Linux 下尤其明显

### v0.1.0-alpha
* 基础的 HTTP API 提供邮件发送功能
* 支持 PLAIN 和 HTML 两种邮件格式