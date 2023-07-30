package util

import "time"

// NowTimestamp 获取当前时间戳
func NowTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
