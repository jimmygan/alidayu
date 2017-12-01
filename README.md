# 阿里大于API开发包

阿里大于短信golang SDK，对应版本阿里云接口版本 2017-05-25

# 使用方法

`go get github.com/jimmygan/alidayu`

# API列表

1. 短信发送API(SendSms)

2. 短信查询API(QuerySendDetails)

3. 短信消息API ```(TODO)```

# DEMO

```
package main

import (
	"fmt"
	"github.com/jimmygan/alidayu"
)

func main()  {
	fmt.Println("aliyun短信测试")
	// 初始化参数，请到阿里云申请密钥
	alidayu.InitAPI("xxxxx", "xxxxxx")
	params := make(map[string]interface{})
	params["date"] = "2017-12-01"

	// 发送短信
	result,err := alidayu.SendSms("测试","templateCode","13900000000","outId",params,"")
	fmt.Println(result, err)

	// bizId := result.BizId
	// 短信发送成功之后 返回bizId
	bizId := "123"

	// 查询发送记录
	result2,err1 := alidayu.QuerySendDetails("13900000000", "20171201",bizId,10, 1)
	fmt.Println(result2, err1)
}

```
