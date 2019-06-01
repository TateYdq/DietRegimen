package database

import "github.com/sirupsen/logrus"

type DiseaseInfo struct{
	DiseaseID int `json:"disease_id"`
	Name string `json:"name"`
	DiseaseKind string `json:"disease_kind"`
	Info string `json:"info"`
	PhotoPath string `json:"photo_path"`
	VoicePath string `json:"voice_path"`
	ViewCount int `json:"view_count"`
	CollectCount string `json:"collect_count"`
}

func GetDiseaseInfoByID(diseaseID int)(diseaseInfo DiseaseInfo,err error){
	err = DrDatabase.Model(DiseaseInfo{}).Where("disease_id="+string(diseaseID)).First(&diseaseInfo).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetDiseaseInfoByID err,diseaseID:%v",diseaseID)
	}
	return diseaseInfo,err
}