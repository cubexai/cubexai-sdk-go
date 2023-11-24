package main

import (
	"fmt"

	"github.com/cubexai/cubexai-sdk-go/utils"
)

func main() {
	appid := "xxx"
	secret := "xxx"

	// QueryMessage(appid, secret, "fe58156e-3425-4fb0-a92a-d1c8fb47f94f")
	SendMessage(appid, secret, "你好")
}

// 接收消息
func QueryMessage(appid, secret, aid string) {

	params := utils.CubeXAIMessageRequestBody{
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
		ModelId: "f7a47e8d-0013-4cbc-ab2c-172b40395a87",
		Messages: []utils.CubeXAICompletionRequestBodyMessages{
			{
				Role:    "user",
				Content: content,
			},
		},
		ModelType:    "gpt-4",
		ModelVersion: "4-8k",
	}

	api := "https://chat.airb3.cn/api/v1/openapi/chat/completions"

	client := utils.NewHttpClient(appid, secret)

	result, _err := client.DoPost(api, params)
	fmt.Println(string(result), _err)
}

func EmbeddingText(text string) {
	api := "https://www.cubexai.cn/api/v1/openapi/embedding/text"
	params := utils.EmbeddingParams{
		Input: text,
	}

	client := utils.NewHttpClient(appid, secret)

	result, _err := client.DoPost(api, params)

	fmt.Println(result, _err)
}
