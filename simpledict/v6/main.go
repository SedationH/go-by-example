package main

import (
	"context"
	"fmt"
	"os"
)

/*
	流程梳理
		1. 获取用户输入
		2. 选择翻译引擎
		3. 进行请求 获得数据 cURL https://curlconverter.com/#go
		4. 反序列化 数据 json2go
		5. 打印结果
*/
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
		`)
		os.Exit(1)
	}
	word := os.Args[1]
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		queryWithVolcengine(word)
		cancel()
	}()

	go func() {
		queryWithCaiyun(word)
		cancel()
	}()

	<-ctx.Done()
}
