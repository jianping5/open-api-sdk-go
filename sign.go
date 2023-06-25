package openapisdkgo

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenSign(s1, s2 string) string {

	// 创建一个 SHA-256 哈希对象
	hash := sha256.New()

	// 将字符串转换为字节数组，然后写入哈希对象
	hash.Write([]byte(s1))
	hash.Write([]byte(s2))

	// 计算哈希值并转换为十六进制字符串
	hashValue := hash.Sum(nil)
	hashString := hex.EncodeToString(hashValue)

	return hashString
}