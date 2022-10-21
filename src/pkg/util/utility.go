package util

import (
	"strings"
	"time"
)

func TimeStamp() string {
	now := time.Now()
	return now.Format("20060102150405")
}

func TrimByteToString(b []byte) string {
	str := string(b)
	return strings.Join(strings.Fields(str), " ")
}
