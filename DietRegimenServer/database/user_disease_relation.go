package database

import (
	"github.com/sirupsen/logrus"
)

type UserDiseaseRelation struct {
	UserID int `json:"user_id"`
	DiseaseID int `json:"disease_id"`
	DiseaseName string `json:"disease_name"`
	Score int `json:"score"`
}


func CreateUserAndDiseaseScore(userID int,diseaseID int){
	diseaseName := GetDiseaseNameByDiseaseID(diseaseID)
	object := UserDiseaseRelation{
		UserID:userID,
		DiseaseID:diseaseID,
		DiseaseName:diseaseName,
		Score:0,
	}
	err := DrDatabase.Create(&object).Error
	if err != nil{
		logrus.WithError(err).Errorf("CreateUserAndDiseaseScore,userID:%v,diseaseID:%v",userID,diseaseID)
	}
}

func AddUserAndDiseaseScore(userID int,diseaseID int,score int)(err error){
	var count int
	DrDatabase.Model(UserDiseaseRelation{}).Where("user_id = ? and disease_id = ? ",userID,diseaseID).Count(&count)
	if count == 0{
		CreateUserAndDiseaseScore(userID,diseaseID)
	}

	err = DrDatabase.Raw("update user_disease_relation set score = score + ? where user_id = ? and disease_id = ? ",score,userID,diseaseID).Error
	if err != nil{
		logrus.WithError(err).Errorf("AddUserAndDiseaseScore,userID:%v,diseaseID:%v",userID,diseaseID)
	}
	return err
}


func GetRecDiseaseByUserID(userID int)(diseaseInfo[] DiseaseInfo,err error){
	var relations []UserDiseaseRelation
	err = DrDatabase.Model(UserDiseaseRelation{}).Where("user_id = ? ",userID).Order("score desc").Scan(&relations).Error
	if err != nil {
		logrus.WithError(err).Errorf("GetRecDiseaseByUserID,userID:%v",userID)
		return diseaseInfo,err
	}
	var diseaseIDs []int
	for _,relation := range relations{
		diseaseIDs = append(diseaseIDs, relation.DiseaseID)
	}
	err = DrDatabase.Model(DiseaseInfo{}).Where("disease_id in (?)",diseaseIDs).Scan(&diseaseInfo).Error
	if err != nil {
		logrus.WithError(err).Errorf("GetRecDiseaseByUserID,userID:%v",userID)
	}
	return diseaseInfo,err
}
