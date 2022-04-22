package domain

import (
	"time"
)

// Domain Model And Interface
type User struct {
	UserName     string
	Password     string
	Age          uint8
	Gender       string
	NickName     string
	Avatar       string
	Birthday     time.Time
	Introduction string
	Email        string
	Phone        string
	State        string
	RegisterTime time.Time
	LoginTime    time.Time
	LastIp       string
}
