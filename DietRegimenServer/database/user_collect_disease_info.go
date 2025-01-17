package database

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/sirupsen/logrus"
)

type UserCollectDiseaseInfo struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	DiseaseID int `json:"disease_id"`
	RecordTime string `json:"record_time"`
}

func GetCollectDiseaseByUserID(userID int)(collectDisease[] UserCollectDiseaseInfo,err error){
	//注意此处为取多行，非一行
	err = DrDatabase.Model(UserCollectDiseaseInfo{}).Where("user_id = ?",userID).Find(&collectDisease).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetCollectDiseaseByUserID err,userID:%v",userID)
	}
	return collectDisease,err
}


func CreateCollectDiseaseInfo(userID int,diseaseID int)(err error){
	var request UserCollectDiseaseInfo
	request.UserID = userID
	request.DiseaseID = diseaseID
	request.RecordTime = utils.GetCurTime()
	db := DrDatabase.Create(&request)
	err = db.Error
	if err != nil{
		logrus.WithError(err).Errorf("CreateCollectDiseaseInfo err,userID:%v,diseaseID:%v",userID,diseaseID)
		return err
	}
	go AddUserAndDiseaseScore(userID,diseaseID,utils.ScoreDiseaseCollect)
	return err
}

func IsUserCollectedDisease(userID int,diseaseID int)(bool){
	count := 0
	DrDatabase.Model(UserCollectDiseaseInfo{}).Where("user_id = ? and disease_id = ?",userID,diseaseID).Count(&count)
	if count == 1{
		return true
	}
	return false
}

func DeleteCollectedDisease(userID int,disease_id int)(bool){
	err := DrDatabase.Where("user_id = ? and disease_id = ?",userID,disease_id).Delete(UserCollectDiseaseInfo{}).Error
	if err != nil {
		return false
	}else {
		return true
	}
}