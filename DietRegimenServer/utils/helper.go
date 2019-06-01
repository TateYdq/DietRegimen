package utils

import (
	"os"
)

const(
	Success = 2000 //成功
	Forbidden = 4003// 禁止访问
	Failed = 5000 // 失败
	ServerError = 5003 // 5003服务器异常
)

func ENV() string {
	env := os.Getenv("ENV")
	if env != "" {
		return env
	}
	return "test"
}
