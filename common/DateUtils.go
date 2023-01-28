package common

import (
	"time"
)

type DateUtils struct {
}

func (that DateUtils) GetCommonTimeStr() string {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	return timeStr
}
func (that DateUtils) GetCommonTimeStrAndMillisecond() string {
	timeStr := time.Now().Format("2006-01-02 15:04:05.555")
	return timeStr
}
func (that DateUtils) GetCommonDate() string {
	timeStr := time.Now().Format("2006-01-02")
	return timeStr
}
