package main

import (
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
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD [ENGINE]
Tanslation engine supports "caiyun" and "volcengine" (default) 火山
- https://fanyi.caiyunapp.com/
- https://translate.volcengine.com/translate
example: simpleDict hello
simpleDict hello volcengine
simpleDict hello caiyun
		`)
		os.Exit(1)
	}
	args := make([]string, 3)
	copy(args, os.Args)
	word, engine := args[1], args[2]

	switch engine {
	case "volcengine", "":
		{
			queryWithVolcengine(word)
		}
	case "caiyun":
		{
			queryWithCaiyun(word)
		}
	}
}
