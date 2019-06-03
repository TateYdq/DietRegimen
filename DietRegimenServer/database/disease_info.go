package database

import (
	"errors"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/sirupsen/logrus"
)

type DiseaseInfo struct{
	ID int
	DiseaseID int `json:"disease_id"`
	Name string `json:"name"`
	DiseaseKind string `json:"disease_kind"`
	Info string `json:"info"`
	PhotoPath string `json:"photo_path"`
	VoicePath string `json:"voice_path"`
	ViewCount int `json:"view_count"`
	CollectCount int `json:"collect_count"`
}

func CreateDiseaseAdmin(request DiseaseInfo)(int,error) {
	db := DrDatabase.Create(&request)
	err := db.Error
	if  err != nil {
		logrus.WithError(err).Error("CreateDiseaseAdmin failed")
	}
	diseaseID := request.ID
	return diseaseID,err
}

func GetDiseaseInfoByID(diseaseID int)(diseaseInfo DiseaseInfo,err error){
	err = DrDatabase.Model(DiseaseInfo{}).Where("disease_id = ?",diseaseID).First(&diseaseInfo).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetDiseaseInfoByID err,diseaseID:%v",diseaseID)
	}
	return diseaseInfo,err
}

func UpdateDiseaseInfo(request DiseaseInfo)(err error){
	if request.DiseaseID == 0{
		logrus.Errorf("diseaseID is equals to 0")
		return errors.New("diseaseID is equals to 0")
	}
	record := utils.ScanStruct(request,[]string{"diseaseID"})
	err = DrDatabase.Model(DiseaseInfo{}).Where("disease_id = ?",request.DiseaseID).Updates(record).Error
	if err != nil{
		logrus.WithError(err).Errorf("UpdateDiseaseInfo err,diseaseID:%v",request.DiseaseID)
	}
	return err

}