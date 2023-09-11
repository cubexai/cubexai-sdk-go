package utils

type CubeXAIRequestHeader struct {
	XAppId     string `json:"x-appid"`
	XTimestamp string `json:"x-timestamp"`
	XSignature string `json:"x-signature"`
	XNonce     string `json:"x-nonce"`
}

type CubeXAICompletionRequestBody struct {
	ModelId      string                                 `json:"mid"`
	Messages     []CubeXAICompletionRequestBodyMessages `json:"messages"`
	ModelType    string                                 `json:"model_type"`
	ModelVersion string                                 `json:"model_version"`
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
