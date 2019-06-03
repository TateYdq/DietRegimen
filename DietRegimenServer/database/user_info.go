package database

import (
	"errors"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type UserInfo struct{
	ID int
	UserID int `json:"user_id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender",gorm:"default:'male'"`
	UserImagePath string `json:"user_image_path"`
	DiseasesFocus string `json:"diseases_focus"`
	Keywords string `json:"keywords"`
}



func CreateUserAdmin(request UserInfo)(int,error){
	db := DrDatabase.Create(&request)
	err := db.Error
	if  err != nil {
		logrus.WithError(err).Error("CreateUserAdmin failed")
	}
	userID := request.ID
	return userID,err
}


//暂时userID就是openID
func GetOrCreateUserInfoByOpenID(openID int)(userInfo UserInfo,err error){
	userInfo,err = GetUserInfoByID(openID)
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			logrus.WithError(err).Error("GetOrCreateInfoByOpenID get userInfo err,then try to create one ")
			userInfo,err = CreateUserByOpenID(openID)
			if err != nil{
				logrus.WithError(err).Error("GetOrCreateInfoByOpenID create err")
				return userInfo,err
			}
			return userInfo,nil
		}else{
			logrus.WithError(err).Error("GetOrCreateInfoByOpenID err")
			return userInfo,err
		}
	}else{
		return userInfo,nil
	}
}



//创建用户，性别默认为男性,TODO:名字，id
func CreateUserByOpenID(openID int)(newUserInfo UserInfo,err error){
	newUserInfo = UserInfo{
		UserID: openID,
		Gender: "male",
	}
	err = DrDatabase.Model(UserInfo{}).Create(newUserInfo).Error
	return newUserInfo,err
}

func GetUserInfoByID(userID int)(userInfo UserInfo,err error){
	err = DrDatabase.Model(UserInfo{}).Where("user_id = ?",userID).First(&userInfo).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetUserInfoByID err,userID:%v",userID)
	}
	return userInfo,err
}

func GetUserNameByID(userID int)(name string,err error){
	var userInfo UserInfo
	err = DrDatabase.Model(UserInfo{}).Where("user_id = ?",userID).First(&userInfo).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetUserInfoByID err,userID:%v",userID)
	}
	return userInfo.Name,err
}

func UpdateUserInfo(userInfo UserInfo)(err error){
	if userInfo.UserID == 0{
		logrus.Errorf("userID is equals to 0")
		return errors.New("userID is equals to 0")
	}
	record := utils.ScanStruct(userInfo,[]string{"userID"})
	err = DrDatabase.Model(UserInfo{}).Where("user_id = ?",userInfo.UserID).Updates(record).Error
	if err != nil{
		logrus.WithError(err).Errorf("updateUserInfo err,userID:%v",userInfo.UserID)
	}
	return err
}