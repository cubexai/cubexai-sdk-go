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
	OpenId       string                                 `json:"open_id"`
	AppId        string                                 `json:"app_id"`
	X            uint64                                 `json:"x"`
	Y            uint64                                 `json:"y"`
}

type CubeXAICompletionRequestBodyMessages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CubeXAICompletionResponse struct {
	Code    uint64                        `json:"code"`
	Message string                        `json:"message"`
	Data    CubeXAICompletionResponseBody `json:"data"`
}

type CubeXAICompletionResponseBody struct {
	Aid string `json:"aid"`
}

type CubeXAIMessageRequest struct {
	Aid string `json:"aid"`
}

type CubeXAIMessageResponse struct {
	Code    uint64                     `json:"code"`
	Message string                     `json:"message"`
	Data    CubeXAIMessageResponseBody `json:"data"`
}

type CubeXAIMessageResponseBody struct {
	Role     string `json:"role"`
	Content  string `json:"content"`
	Balance  uint64 `json:"balance"`
	Quantity uint64 `json:"quantity"`
	Status   string `json:"status"`
}
