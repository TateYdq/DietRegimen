package database

import (
	"errors"
	"fmt"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/cache"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type UserInfo struct{
	UserID int `json:"user_id" gorm:"column:user_id;primary_key"`
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender" gorm:"default:'male'"`
	UserImagePath string `json:"user_image_path"`
	DiseasesFocus string `json:"diseases_focus"`
	Keywords string `json:"keywords"`
	OpenID string `json:"open_id"`
	UserScore int `json:"user_score"`
	NoAttention int `json:"no_attention" gorm:"default:0"`
	UpdateTime string `json:"update_time"`
}


type WechatLoginRequestBody struct {
	Code string `json:"code"`
	NickName string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	Gender string `json:"gender"`
}

func CreateUserAdmin(request UserInfo)(int,error){
	db := DrDatabase.Create(&request)
	err := db.Error
	if  err != nil {
		logrus.WithError(err).Error("CreateUserAdmin failed")
	}
	userID := request.UserID
	return userID,err
}


//暂时userID就是openID
func GetOrCreateUserInfoByOpenID(openID string,request WechatLoginRequestBody)(userInfo UserInfo,err error){
	userInfo,err = GetUserInfoByOpenID(openID)
	//如果openID不存在的话,证明为首次登录
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			logrus.WithError(err).Error("GetOrCreateInfoByOpenID get userInfo err,then try to create one ")
			userImagePath,err := helper.GetJpegImg(request.AvatarUrl)
			if err != nil{
				logrus.WithError(err).Error("GetUserImage err")
				return userInfo,err
			}
			userInfo,err = CreateUserByOpenID(openID,request.NickName,request.Gender,userImagePath)
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

func GetUserInfoByOpenID(openID string)(userInfo UserInfo,err error){
	err = DrDatabase.Model(UserInfo{}).Where("open_id = ?",openID).First(&userInfo).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetUserInfoByID err,openID:%v",openID)
	}
	return userInfo,err
}



//创建用户，性别默认为男性,默认60岁
func CreateUserByOpenID(openID string,name string,gender string,userImagePath string)(newUserInfo UserInfo,err error){
	newUserInfo = UserInfo{
		OpenID: openID,
		Name: name,
		Gender: gender,
		Age: 60,
		NoAttention:0,
		UserImagePath: userImagePath,
	}
	err = DrDatabase.Model(UserInfo{}).Create(&newUserInfo).Error
	return newUserInfo,err
}

func GetUserInfoByUserID(userID int)(userInfo UserInfo,err error){
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
	record["no_attention"] = userInfo.NoAttention
	err = DrDatabase.Model(UserInfo{}).Where("user_id = ?",userInfo.UserID).Updates(record).Error
	if err != nil{
		logrus.WithError(err).Errorf("updateUserInfo err,userID:%v",userInfo.UserID)
	}
	return err
}

func UpdateUserPath(userID int,path string)(err error){
	if userID == 0{
		logrus.Errorf("userID is equals to 0")
		return errors.New("userID is equals to 0")
	}
	record := make(map[string]interface{})
	record["user_image_path"] = path
	err = DrDatabase.Model(UserInfo{}).Where("user_id = ?",userID).Updates(record).Error
	if err != nil{
		logrus.WithError(err).Errorf("UpdateUserPath err,userID:%v",userID)
	}
	logrus.Infof("UpdateUserPath success,userID:%v",userID)
	return err
}

func AddUserScore(userID int,score int)(err error){
	//限制每天得到的分数
	if score == utils.ScoreUserLook{
		key := fmt.Sprint(cache.UserLookFreqKey,userID)
		value,err := cache.CreateOrSetAddOne(key,utils.DurationUserScore)
		if err != nil{
			return err
		}
		if value > utils.LimitUserLookAwardDay{
			return nil
		}
	}else if score == utils.ScoreUserComment{
		key := fmt.Sprint(cache.UserCommentFreqKey,userID)
		value,err :=  cache.CreateOrSetAddOne(key,utils.DurationUserScore)
		if err != nil{
			return err
		}
		if value > utils.LimitUserCommentAwardDay{
			return nil
		}
	}else if score == utils.ScoreUserLogin{
		key := fmt.Sprint(cache.UserLoginFreqKey,userID)
		value,err :=  cache.CreateOrSetAddOne(key,utils.DurationUserScore)
		if err != nil{
			return err
		}
		if value > utils.LimitUserLoginAwardDay{
			return nil
		}
	}else{
		return nil
	}
	err = DrDatabase.Exec("update user_info set user_score = user_score + ? where user_id = ? ",score,userID).Error
	if err != nil{
		logrus.WithError(err).Errorf("AddUserAndFoodScore,userID:%v",userID)
	}
	return err
}

func GetUserImagePath(userID int)(path string,err error){
	var userInfo UserInfo
	err = DrDatabase.Model(UserInfo{}).Where("user_id = ?",userID).First(&userInfo).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetUserImagePath,userID:%v",userID)
		return "",err
	}
	return userInfo.UserImagePath,nil
}