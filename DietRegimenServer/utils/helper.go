package utils

import (
	"os"
	"time"
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

func GetCurTime()(t string){
	t = time.Now().Format("2006-01-02 15:04:05")
	return t
}
