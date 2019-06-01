package database

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type UserInfo struct{
	UserID int `json:"user_id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	UserImagePath string `json:"user_image_path"`
	DiseasesFocus string `json:"diseases_focus"`
	Keywords string `json:"keywords"`
}

func GetUserInfoByID(userID int)(userInfo UserInfo,err error){
	err = DrDatabase.Where("user_id="+string(userID)).First(&userInfo).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetUserInfoByID err,userID:%v",userID)
	}
	return userInfo,err
}

func UpdateUserInfo(userInfo UserInfo)(err error){
	if userInfo.UserID == 0{
		return errors.New("userID is equals to 0")
	}
	record := make(map[string]interface{})
	if userInfo.Name != ""{
		record["name"] = userInfo.Name
	}
	if userInfo.Age != 0{
		record["age"] = userInfo.Age
	}
	if userInfo.Gender != ""{
		record["gender"] = userInfo.Gender
	}
	if userInfo.UserImagePath != ""{
		record["user_image_path"] = userInfo.UserImagePath
	}
	if userInfo.DiseasesFocus != ""{
		record["diseases_focus"] = userInfo.DiseasesFocus
	}
	if userInfo.Keywords != ""{
		record["keywords"] = userInfo.Keywords
	}
	err = DrDatabase.Where("user_id="+string(userInfo.UserID)).Updates(record).Error
	if err != nil{
		logrus.WithError(err).Errorf("updateUserInfo err,userID:%v",userInfo.UserID)
	}
	return err
}