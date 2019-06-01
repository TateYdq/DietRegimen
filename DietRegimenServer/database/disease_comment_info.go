package database

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/sirupsen/logrus"
)


type DiseaseCommentInfo struct{
	ID int `json:"id"`
	DiseaseID int `json:"disease_id"`
	UserID int `json:"user_id"`
	UserName string `json:"user_name"`
	Comment string `json:"comment"`
	RecordTime string `json:"record_time"`
}

func GetCommentByDiseaseID(diseaseID int)(diseaseComments[] DiseaseCommentInfo,err error){
	err = DrDatabase.Model(DiseaseCommentInfo{}).Where("disease_id"+string(diseaseID)).Find(&diseaseComments).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetCommentByDiseaseID err,diseaseID:%v",diseaseID)
	}
	return diseaseComments,err
}

func CreateDiseaseComment(diseaseID int,userID int, content string)(err error){
	name,err := GetUserNameByID(userID)
	if err != nil{
		return err
	}
	fc := DiseaseCommentInfo{
		DiseaseID:diseaseID,
		UserID:userID,
		UserName:name,
		Comment:content,
		RecordTime:utils.GetCurTime(),
	}
	err = DrDatabase.Model(DiseaseCommentInfo{}).Create(fc).Error
	if err != nil{
		logrus.WithError(err).Errorf("CreateDiseaseComment err,diseaseID:%v,userID:%v",diseaseID,userID)
		return err
	}
	return nil
}
