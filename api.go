package openapisdkgo

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type OpenApiClient struct {
	AccessKey string
	SecretKey string
}

func NewClient(accessKey, secretKey string) *OpenApiClient {
	// 创建 openApiClient
	openApiClient := &OpenApiClient{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	return openApiClient
}

func (c *OpenApiClient) AddHeaders(req *http.Request, userRequestParams string) {
	// 请求体
	body := userRequestParams

	// Content-Type
	req.Header.Set("Content-Type", "application/json")

	// accessKey，用户标识
	// TODO: 从配置中获取
	accessKey := c.AccessKey
	req.Header.Add("accessKey", accessKey)

	// nonce，随机数
	// 设置随机种子，保证每次运行都有不同的随机数序列
	rand.Seed(time.Now().UnixNano())

	// 生成小于 10000 的随机数
	nonce := rand.Int63n(10000)
	req.Header.Add("nonce", strconv.FormatInt(nonce, 10))

	// timestamp，当前时间戳
	timestamp := time.Now().Unix()
	req.Header.Add("timestamp", strconv.FormatInt(timestamp, 10))

	// 签名（请求体 + secretKey）
	// TODO: secretKey 从配置中获取
	secretKey := c.SecretKey
	sign := GenSign(body, secretKey)
	req.Header.Add("sign", sign)

	// 请求体
	req.Header.Add("body", userRequestParams)

}

func (c *OpenApiClient) SayHelloUsingGet() string {
	client := &http.Client{}

	// 构造请求
	req, err := http.NewRequest("GET", "http://localhost:3000/api/hello", nil)
	if err != nil {
		return ""
	}

	// 添加请求头
	c.AddHeaders(req, "")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	// 读取结果
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	// 返回结果
	return string(bytes)
}

func (c *OpenApiClient) GetNameByIdUsingPost(userRequestParams string) string {
	client := &http.Client{}

	// 构造请求
	req, err := http.NewRequest("POST", "http://localhost:3000/api/name", strings.NewReader(userRequestParams))
	if err != nil {
		return ""
	}

	// 添加请求头
	c.AddHeaders(req, userRequestParams)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	// 读取结果
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "567"
	}

	// 返回结果
	return string(bytes)

}
