// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/29 00:02:13

package main

import (
	"bytes"
	"net/http"
	"testing"
)

// Benchmark the http server.
// BenchmarkHttpServer-8               1795            668408 ns/op           16885 B/op        120 allocs/op
func BenchmarkHttpServer(b *testing.B) {

	body := []byte(`
{
  "email": {
    "to": "fishinlove@163.com",
    "subject": "测试 postar 运行情况",
    "contentType": "text/html; charset=utf-8",
    "body": "<h1>哈喽！来自 <span style=\"color: #123456;\">postar<span> 的问候！</h1>"
  },
  "options": {
    "sync": true
  }
}
    `)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := http.Post("http://localhost:5779/send", "application/json; charset=utf-8", bytes.NewBuffer(body))
		if err != nil {
			b.Fatal(err)
		}
		resp.Body.Close()
	}
}
