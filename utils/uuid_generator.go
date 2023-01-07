package utils

import (
	"strconv"
	"time"
)

// GenerateUuid Generate a random uuid (NOT RFC4122 compliant)
func GenerateUuid() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
