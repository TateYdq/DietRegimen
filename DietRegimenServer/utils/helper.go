package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
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
	return "staging"
}

func IsOnlineENV() bool {
	return ENV() == "online"
}

func IsStagingENV() bool {
	return ENV() == "staging"
}


func GetCurTime()(t string){
	t = time.Now().Format("2006-01-02 15:04:05")
	return t
}


func ScanStruct(s interface{},excludeKey[] string)(map[string]interface{}){
	results := make(map[string]interface{})
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	for k := 0; k < t.NumField(); k++ {
		isExclude := false
		for _,key := range excludeKey{
			if t.Field(k).Name  == key {
				isExclude = true
				break;
			}
		}
		if isExclude{
			continue
		}
		switch v.Field(k).Kind() {
		case reflect.String:
			if v.Field(k).String() != ""{
				results[t.Field(k).Name] = v.Field(k).String()
				logrus.Infof("add %v - %v \n",t.Field(k).Name,v.Field(k).String())
			}
			break
		case reflect.Int:
			if v.Field(k).Int() != 0 {
				results[t.Field(k).Name] = v.Field(k).Int()
				logrus.Infof("add %v - %v \n",t.Field(k).Name,v.Field(k).Int())
			}
			break
		}
	}
	return results;
}