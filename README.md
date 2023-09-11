# CubeX AI Go语言SDK

本SDK是基于Go语言构建的，通过一行代码便可访问GPT大模型，利用其强大的自然语言处理技术，使您能够快速获取卓越的AI功能。

👉 CubeX AI: [CubeX AI官方网站](https://www.airb3.com.cn/?ref=github)

👉 API文档：[API文档链接](https://apifox.com/apidoc/shared-c2de4a48-bf44-4a6c-aacc-554885ac180e)

👉 官方公众号：CubeX AI

<img src="./images/so.png" />

### 前置准备
在使用之前，您需要在官方网站上申请API密钥（API key）和密钥串（API Secret）。

### 1. 快速开始

#### 1.1. 安装SDK

```shell
go get github.com/cubexai/cubexai-sdk-go
```

#### 1.2. 在CubeX AI官网上获取API密钥对
![Alt text](./images/apikey.png)

#### 1.3. 初始化Client对象并配置SDK密钥对

**注意事项：**

1. `completion`接口是异步的。调用后，接口会同步返回`aid`参数，此参数用于查询本次请求AI响应消息的ID。
2. 需要调用方自行轮询消息查询接口以获取AI回复的消息内容。服务端对消息的返回是流式分段增量返回。
3. 消息查询接口响应中的`status`字段有3种枚举值：
    * `RUN`: AI正在生成并回复内容，需继续轮询。
    * `LENGTH`: 请求的消息长度超过模型的MaxTokens参数规定长度，需回复【继续】指令。
    * `END`: 请求已完全处理且消息返回完毕，应停止轮询。
4. 当`status`字段状态为`LENGTH`或`END`时，服务端保留消息的有效期仅为1分钟。请及时获取消息内容并进行缓存或存储。

```
package main

import (
	"github.com/cubexai/cubexai-sdk-go/utils"
	"fmt"
)

func main() {
	appid := "6WAz************************0eL"   // 官网获得的APIKEY
	secret := "Nuf*************************h9Y"  // 官网获得的APISecret

	// 以下调用方式为快速跑通Demo，可以根据实际情况自行修改

	QueryMessage(appid, secret, "b8c24e81-cdcb-4aae-a7cf-abcdefg")
	SendMessage(appid, secret, "你好啊")
}

// 接收消息
func QueryMessage(appid, secret, aid string) {

	// 使用内置结构体构造请求参数
	params := utils.CubeXAIMessageRequestBody{
		Aid: aid,
	}

	// API接口URL
	api := "https://chat.airb3.cn/api/v1/openapi/chat/query"

	// 传递ak实例化一个客户端
	// 内置客户端实例提供了计算签名的方法，不用手动计算签名
	client := utils.NewHttpClient(appid, secret)

	// 使用初始化过的实例调用DoGet方法
	result, _err := client.DoGet(api, params)

	// 返回的是[]byte，需要配合结构体解构使用
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
		ModelType:    "gpt-35",
		ModelVersion: "",
	}

	api := "https://chat.airb3.cn/api/v1/openapi/chat/completions"

	client := utils.NewHttpClient(appid, secret)

	result, _err := client.DoPost(api, params)
	fmt.Println(string(result), _err)
}
```

#### 1.4. 参数结构说明

###### 1.4.1. 请求参数结构

以下内容仅作简要描述，详细信息请参见 [API文档](https://apifox.com/apidoc/shared-c2de4a48-bf44-4a6c-aacc-554885ac180e)。

* 请求头

```
type CubeXAIRequestHeader struct {
	XAppId     string `json:"x-appid"`
	XTimestamp string `json:"x-timestamp"`
	XSignature string `json:"x-signature"`
	XNonce     string `json:"x-nonce"`
}
```

* Completion接口
    * ModelId 模型ID，需要在官网已有的模型中查询
    * Messages 消息内容数组，可以携带多条历史记录，最新的消息放到最后。
        * -Messages.Role 可选值有`system`,`assistant`,`user`
        * -Messages.Content 消息内容正文
	* ModelType:  模型类型，默认gpt-3.5-turbo,
	* ModelVersion 模型版本，可为空
```
type CubeXAICompletionRequestBody struct {
	ModelId  string                                 `json:"mid"`
	Messages []CubeXAICompletionRequestBodyMessages `json:"messages"`
	ModelType:    "gpt-35",
	ModelVersion: "",
}

type CubeXAICompletionRequestBodyMessages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CubeXAICompletionResponseBody struct {
	Code    uint64                            `json:"code"`
	Message string                            `json:"message"`
	Data    CubeXAICompletionResponseBodyData `json:"data"`
}

type CubeXAICompletionResponseBodyData struct {
	Aid string `json:"aid"`
}
```

* QueryMessage接口
    * Aid 消息ID
```
type CubeXAIMessageRequestBody struct {
	Aid string `json:"aid"`
}

type CubeXAIMessageResponseBody struct {
	Code    uint64                         `json:"code"`
	Message string                         `json:"message"`
	Data    CubeXAIMessageResponseBodyData `json:"data"`
}

type CubeXAIMessageResponseBodyData struct {
	Role     string `json:"role"`
	Content  string `json:"content"`
	Balance  uint64 `json:"balance"`
	Quantity uint64 `json:"quantity"`
	Status   string `json:"status"`
}

```

###### 1.4.2 响应参数结构

* Code参数值为`10000`，且响应码为`200`时，代表请求成功，其他均为异常
    

```
type CubeXAIMessageResponse struct {
	Code    uint64                     `json:"code"`
	Message string                     `json:"message"`
	Data    interface{}                `json:"data"`
}
```

### 1.5. 签名认证

**温馨提示：** SDK已内置签名认证方法，可直接调用`doGet`或`doPost`方法发起请求，无需手动计算签名。

我们采用基于HMAC的请求签名机制以确保API请求的安全性。每个API请求必须包括以下HTTP头部字段：`X-APPID`，`X-TIMESTAMP`，`X-NONCE`和`X-SIGNATURE`。

* `X-APPID`：后端颁发的应用程序标识符。
* `X-TIMESTAMP`：请求发起的Unix时间戳。
* `X-NONCE`：保证每次签名唯一性的随机字符串。
* `X-SIGNATURE`：基于HMAC生成的请求签名。

签名过程如下：

1. 将请求体参数（如有）按字典序排序，并使用等号（=）拼接键值对，然后用和号（&）连接所有键值对，生成待签名字符串。
2. 使用后端颁发的`SECRET`和待签名字符串，通过HMAC算法生成签名字符串。
3. 服务器会校验`X-TIMESTAMP`，如果低于或超过服务器时间戳`30秒`，将导致验签失败。

