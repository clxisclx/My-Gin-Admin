package utils

import "github.com/google/uuid"

// GenerateUuid 生成uuid
func GenerateUuid() string {
	return uuid.New().String()
}
