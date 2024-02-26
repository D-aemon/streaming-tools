package utils

import "github.com/google/uuid"

func GetUUIDByLength(length int) string {
	return uuid.New().String()[:length]
}
