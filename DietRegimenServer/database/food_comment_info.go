package database

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/sirupsen/logrus"
)

type FoodCommentInfo struct{
	ID int `json:"id"`
	FoodID int `json:"food_id"`
	UserID int `json:"user_id"`
	UserName string `json:"user_name"`
	Comment string `json:"comment"`
	RecordTime string `json:"record_time"`
}

func GetCommentByFoodID(foodID int)(foodComments[] FoodCommentInfo,err error){
	err = DrDatabase.Model(FoodCommentInfo{}).Where("food_id = ?",foodID).Find(&foodComments).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetFoodInfoByID err,foodID:%v",foodID)
	}
	return foodComments,err
}

func CreateFoodComment(foodID int,userID int, content string)(err error){
	name,err := GetUserNameByID(userID)
	if err != nil{
		return err
	}
	fc := FoodCommentInfo{
		FoodID:foodID,
		UserID:userID,
		UserName:name,
		Comment:content,
		RecordTime:utils.GetCurTime(),
	}
	err = DrDatabase.Model(FoodCommentInfo{}).Create(fc).Error
	if err != nil{
		logrus.WithError(err).Errorf("CreateFoodComment err,foodID:%v,userID:%v",foodID,userID)
		return err
	}
	return nil
}
