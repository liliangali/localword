package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/spf13/cast"
	"time"
)

func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return string(hex.EncodeToString(m.Sum(nil)))
}
func GetUnix() int64 {
	return time.Now().Unix()
}

func GetDayLimit() int64 {
	now := time.Now()
	tomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	diff := tomorrow.Sub(now)
	return cast.ToInt64(diff.Seconds()) - 10
}

func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}
