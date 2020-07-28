// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/07/29 00:02:13

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/avino-plan/postar/src/core"
	"github.com/avino-plan/postar/src/models"
)

// Benchmark the http server.
//
func BenchmarkHttpServer(b *testing.B) {

	sendTask := models.NewSendTaskWithDefaultOptions()
	sendTask.Email = &core.Email{
		To:          "fishinlove@163.com",
		Subject:     "jsonrpc 测试 postar 运行情况",
		ContentType: "text/html; charset=utf-8",
		Body:        "<h1>哈喽！来自 <span style=\"color: #123456;\">postar<span> 的问候！</h1>",
	}
	sendTask.Options.Sync = true
	bodyBytes, err := json.Marshal(sendTask)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := http.Post("http://localhost:5779/send", "application/json; charset=utf-8", bytes.NewBuffer(bodyBytes))
		if err != nil {
			b.Fatal(err)
		}
		resp.Body.Close()
	}
}
