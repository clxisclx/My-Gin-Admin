package utils

import (
	"encoding/base64"
	"golang.org/x/crypto/argon2"
	"golang.org/x/exp/rand"
)

// GenerateSalt 生成盐
func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return Base64Encode(salt), nil
}

// HashPassword 密码加密
func HashPassword(password string, salt []byte) string {
	// 使用 Argon2id 哈希密码
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32) // 1 次迭代，64 MB 内存，4 并发，32 字节输出
	return Base64Encode(hash)
}

// ComparePassword 比较密码
func ComparePassword(hashedPassword string, password string, salt string) bool {
	// 使用相同的参数对输入的密码进行哈希
	newHash := HashPassword(password, []byte(salt))
	// 比较哈希值是否相同
	return newHash == hashedPassword
}

// Base64Encode base64 编码
func Base64Encode(data []byte) string {
	return base64.RawStdEncoding.EncodeToString(data)
}

// Base64Decode base64 解码
func Base64Decode(data string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(data)
}
