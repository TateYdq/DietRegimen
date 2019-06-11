package database

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/sirupsen/logrus"
)

type UserCollectFoodInfo struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	FoodID int `json:"food_id"`
	RecordTime string `json:"record_time"`
}

func GetCollectFoodByID(userID int)(collectFood[] UserCollectFoodInfo,err error){
	//注意此处为取多行，非一行
	err = DrDatabase.Model(UserCollectFoodInfo{}).Where("user_id = ?",userID).Find(&collectFood).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetCollectFoodByID err,userID:%v",userID)
	}
	return collectFood,err
}

func CreateCollectFoodInfo(userID int,foodID int)(err error){
	var request UserCollectFoodInfo
	request.UserID = userID
	request.FoodID = foodID
	request.RecordTime = utils.GetCurTime()
	db := DrDatabase.Create(&request)
	err = db.Error
	if err != nil{
		logrus.WithError(err).Errorf("CreateCollectDiseaseInfo err,userID:%v,foodID:%v",userID,foodID)
		return err
	}
	go AddUserAndFoodScore(userID,foodID,utils.ScoreFoodCollect)
	return err
}
