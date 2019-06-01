package database

import "github.com/sirupsen/logrus"

type CollectFood struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	FoodID int `json:"food_id"`
	RecordTime string `json:"record_time"`
}

func GetCollectFoodByID(userID int)(collectFood[] CollectFood,err error){
	err = DrDatabase.Where("user_id="+string(userID)).Find(&collectFood).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetCollectFoodByID err,userID:%v",userID)
	}
	return collectFood,err
}
