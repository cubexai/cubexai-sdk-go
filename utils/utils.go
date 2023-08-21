package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
)

/*
@feature         生成随机字符串
@params var sLen 字符串长度
*/
func GenerateRandomString(sLen uint8) string {

	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, sLen)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

/*
@feature 将结构体Key按字典序排序，生成签名字符串
@params var v  interface{}
*/
func StructOrMapToSortedString(v interface{}) string {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	var fields []string

	switch val.Kind() {
	case reflect.Struct:
		typeOfT := val.Type()
		for i := 0; i < val.NumField(); i++ {
			fields = append(fields, fieldToString(typeOfT.Field(i), val.Field(i)))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			fields = append(fields, fmt.Sprintf("%s=%s", key.Interface(), valueToString(val.MapIndex(key))))
		}
	default:
		return "Unsupported type"
	}

	sort.Strings(fields)
	return strings.Join(fields, "&")
}

/*
@feature 计算签名
@params var data string      签名字符串
@params var Nonce string     随机字符串
@params var Timestamp string 时间戳
@params var secret string    秘钥
@return signature string     签名值
*/
func GenerateSignature(data string, Nonce string, Timestamp string, secret string) string {

	data = data + Nonce + Timestamp

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func valueToString(value reflect.Value) string {
	switch value.Kind() {
	case reflect.Struct:
		return StructOrMapToSortedString(value.Interface())
	case reflect.Slice:
		var sliceStrings []string
		for i := 0; i < value.Len(); i++ {
			sliceStrings = append(sliceStrings, valueToString(value.Index(i)))
		}
		return strings.Join(sliceStrings, ",")
	case reflect.Map:
		return StructOrMapToSortedString(value.Interface())
	default:
		return fmt.Sprintf("%v", value.Interface())
	}
}

func fieldToString(field reflect.StructField, value reflect.Value) string {
	tag := field.Tag.Get("json")
	key := tag
	if key == "" {
		key = field.Name
	}
	return fmt.Sprintf("%s=%s", key, valueToString(value))
}
