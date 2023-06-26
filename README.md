# Open-API Go SDK

## 快速开始
### 1. 引入 sdk
```Go
go get github.com/jianping5/open-api-sdk-go
```

### 2. 编写配置文件
```yaml
openapi:
  accessKey: 你的 accessKey
  secretKey: 你的 secretKey
```

### 3. 读取配置
```Go
type Conf struct {
	OpenApi OpenApi `yaml:"openapi"`
}

type OpenApi struct {
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
}

var Config Conf

func init() {
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("[config init error] ioutil.ReadFile 配置文件读取失败 " + err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		fmt.Println("[config init error] yaml.Unmarshal 配置文件解析失败 " + err.Error())
		return
	}
}
```

### 4. 接口调用示例代码
```Go
func main() {
	// get key pairs from config
	accessKey := config.Config.OpenApi.AccessKey
	secretKey := config.Config.OpenApi.SecretKey

	// New openApiClient
	openApiClient := openapisdkgo.NewClient(accessKey, secretKey)
	
	// demo1：invoke "GET" api
	res := openApiClient.SayHelloUsingGet()
	fmt.Println(res)

	// demo2: invoke "POST" api
	res = openApiClient.GetNameByIdUsingPost("{\"id\": 1}")
	fmt.Println(res)
}
```

## API 文档
### 1. 你好
方法名：SayHelloUsingGet

方法类型：GET

请求参数：无

响应参数：data 响应结果

### 2. 根据 ID 获取姓名
方法名：GetNameByIdUsingPost

方法类型：POST

请求参数：id int

响应参数：data 响应结果

