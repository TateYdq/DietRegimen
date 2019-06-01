package database

import "github.com/sirupsen/logrus"

type CollectDisease struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	DiseaseID int `json:"disease_id"`
	RecordTime string `json:"record_time"`
}

func GetCollectDiseaseByUserID(userID int)(collectDisease[] CollectDisease,err error){
	err = DrDatabase.Where("user_id="+string(userID)).Find(&collectDisease).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetCollectDiseaseByUserID err,userID:%v",userID)
	}
	return collectDisease,err
}