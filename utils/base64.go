package utils

import "encoding/base64"

// Base64Encode base64 编码
func Base64Encode(data []byte) string {
	return base64.RawStdEncoding.EncodeToString(data)
}

// Base64Decode base64 解码
func Base64Decode(data string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(data)
}
