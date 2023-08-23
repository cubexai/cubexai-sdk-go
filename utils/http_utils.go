package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type HttpClient struct {
	AppId  string
	Secret string
	Header CubeXAIRequestHeader
}

var (
	timeout = time.Second * 3
)

func NewHttpClient(AppId, Secret string) *HttpClient {
	return &HttpClient{
		AppId:  AppId,
		Secret: Secret,
	}
}

func (h *HttpClient) DoGet(api string, params interface{}) (result []byte, err error) {
	// 创建一个URL对象
	u, err := url.Parse(api)
	if err != nil {
		return nil, err
	}

	// 使用url.Values来存储查询参数
	query := u.Query()

	v := reflect.ValueOf(params)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("json")
		key := strings.Split(tag, ",")[0] // 获取JSON标签的键
		if key != "" && field.IsValid() && field.CanInterface() {
			value := fmt.Sprintf("%v", field.Interface())
			query.Set(key, value)
		}
	}

	// 将查询参数附加到URL上
	u.RawQuery = query.Encode()

	// 创建一个HTTP GET请求
	req, _ := http.NewRequest("GET", u.String(), nil)

	err = h.generateSign(&params)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-AppId", h.Header.XAppId)
	req.Header.Add("X-Timestamp", h.Header.XTimestamp)
	req.Header.Add("X-Nonce", h.Header.XNonce)
	req.Header.Add("X-Signature", h.Header.XSignature)

	client := http.Client{
		Timeout: timeout,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, _ := io.ReadAll(res.Body)
		return body, nil
	} else {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error status code: %d，%s", res.StatusCode, string(body))
	}
}

func (h *HttpClient) DoPost(api string, body interface{}) (result []byte, err error) {
	// 将body转换为JSON格式
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	// 创建一个HTTP POST请求
	req, err := http.NewRequest("POST", api, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	// 添加JSON内容类型头部
	req.Header.Add("Content-Type", "application/json")

	err = h.generateSign(body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-AppId", h.Header.XAppId)
	req.Header.Add("X-Timestamp", h.Header.XTimestamp)
	req.Header.Add("X-Nonce", h.Header.XNonce)
	req.Header.Add("X-Signature", h.Header.XSignature)

	client := http.Client{
		Timeout: timeout,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return resBody, nil
	} else {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error status code: %d，%s", res.StatusCode, string(body))
	}
}

func (h *HttpClient) generateSign(data interface{}) error {

	uSignature := StructOrMapToSortedString(data)

	h.Header.XAppId = h.AppId
	h.Header.XTimestamp = strconv.FormatInt(time.Now().Unix(), 10)
	h.Header.XNonce = GenerateRandomString(18)
	h.Header.XSignature = GenerateSignature(uSignature, h.Header.XNonce, h.Header.XTimestamp, h.Secret)

	return nil
}
