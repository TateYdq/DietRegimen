package database

import "github.com/sirupsen/logrus"

type AdminInfo struct {
	ID int `json:"id"`
	Token string `json:"token"`
}


func VerifyAdmin(token string)(bool){
	var count int64
	DrDatabase.Model(AdminInfo{}).Where("token = ?",token).Count(&count)
	if count == 1{
		logrus.Infof("admin verify success")
		return true
	}
	return false
}