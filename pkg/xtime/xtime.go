package xtime

import (
	"time"
)

// 通常情况下要确保程序部署环境的时区是否正确
// 1. 主机 Centos
// 2. 容器 Alpine Linux

const (
	TimeLayOut = "2006-01-02 15:04:05"
)

func Now() time.Time {
	// 1. UTC
	// 2. LOCAL
	// 3. IANA [ZONEINFO environment variable]
	// var zoneSources = []string{"/usr/share/zoneinfo/", "/usr/share/lib/zoneinfo/", "/usr/lib/locale/TZ/", runtime.GOROOT() + "/lib/time/zoneinfo.zip",
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

func FormatTime(t time.Time, format string) string {
	return t.Format(format)
}

func MoveTime(t time.Time, d string) string {
	duration, _ := time.ParseDuration(d)
	return FormatTime(Now().Add(duration), TimeLayOut)
}

func ParseTime(t string) time.Time {
	// 时区
	location, _ := time.LoadLocation("Asia/Shanghai")
	// 解析也需要注意时区 	time.Parse() 默认使用UTC
	tm, _ := time.ParseInLocation(TimeLayOut, t, location)
	return time.Unix(tm.Unix(), 0).In(location)
}
