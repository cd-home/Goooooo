package xtime

import (
	"syscall"
	"testing"
	"time"
)

func TestLocalTime(t *testing.T) {
	t.Log(Now())
}

func TestFormatTime(t *testing.T) {
	t.Log(FormatTime(Now(), TimeLayOut))
}

func TestMoveTime(t *testing.T) {
	t.Log(MoveTime(Now(), "-2h"))
}

func TestParseTime(t *testing.T) {
	timeString := "2022-06-14 9:55:01"
	t.Log(ParseTime(timeString).Add(time.Hour))
}

func TestZoneInfo(t *testing.T) {
	env, _ := syscall.Getenv("ZONEINFO")
	t.Log(env)
}
