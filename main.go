package main

import (
	"cubexai-sdk-go/utils"
	"fmt"
)

func main() {
	appid := "6WAz************************0eL"
	secret := "Nuf*************************h9Y"

	QueryMessage(appid, secret, "b8c24e81-cdcb-4aae-a7cf-abcdefg")
	SendMessage(appid, secret, "你好啊")
}

// 接收消息
func QueryMessage(appid, secret, aid string) {

	params := utils.CubeXAIMessageRequest{
		Aid: aid,
	}

	api := "https://chat.airb3.cn/api/v1/openapi/chat/query"

	client := utils.NewHttpClient(appid, secret)

	result, _err := client.DoGet(api, params)
	fmt.Println(string(result), _err)
}

// 发送消息
func SendMessage(appid, secret, content string) {

	params := &utils.CubeXAICompletionRequestBody{
		ModelId: "7d0fce18-0f36-4e01-9952-abcdefghijk",
		Messages: []utils.CubeXAICompletionRequestBodyMessages{
			{
				Role:    "user",
				Content: content,
			},
		},
	}

	api := "https://chat.airb3.cn/api/v1/openapi/chat/completions"

	client := utils.NewHttpClient(appid, secret)

	result, _err := client.DoPost(api, params)
	fmt.Println(string(result), _err)
}