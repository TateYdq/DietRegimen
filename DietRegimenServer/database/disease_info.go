package database

import (
	"errors"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/sirupsen/logrus"
)

type DiseaseInfo struct{
	DiseaseID int `json:"disease_id"  gorm:"column:disease_id;primary_key"`
	Name string `json:"name"`
	DiseaseKind string `json:"disease_kind"`
	Info string `json:"info"`
	Taboo string `json:"taboo"`
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
	diseaseID := request.DiseaseID
	return diseaseID,err
}

func GetDiseaseInfoByID(diseaseID int)(diseaseInfo DiseaseInfo,err error){
	err = DrDatabase.Model(DiseaseInfo{}).Where("disease_id = ?",diseaseID).First(&diseaseInfo).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetDiseaseInfoByID err,diseaseID:%v",diseaseID)
	}
	return diseaseInfo,err
}

func GetDiseaseLists(keyword string)(diseaseList[] DiseaseInfo,err error){
	if keyword == ""{
		err = DrDatabase.Model(DiseaseInfo{}).Scan(&diseaseList).Error
		if err != nil{
			logrus.WithError(err).Errorf("GetDiseaseLists err")
		}
	}else{
		keyword = "%"+keyword+"%"
		err = DrDatabase.Model(DiseaseInfo{}).Where("disease_kind like ? or info like ? or taboo like ?  or name like ? ",keyword,keyword,keyword,keyword).Scan(&diseaseList).Error
		if err != nil {
			logrus.WithError(err).Errorf("SearchByKeyWord err,keyword:v", keyword)
		}
	}
	return diseaseList,err

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

func UpdateDiseasePath(diseaseID int,path string)(err error){
	if diseaseID == 0{
		logrus.Errorf("diseaseID is equals to 0")
		return errors.New("diseaseID is equals to 0")
	}
	record := make(map[string]interface{})
	record["photo_path"] = path
	err = DrDatabase.Model(DiseaseInfo{}).Where("disease_id = ?",diseaseID).Updates(record).Error
	if err != nil{
		logrus.WithError(err).Errorf("UpdateFoodPath err,diseaseID:%v",diseaseID)
	}
	logrus.WithError(err).Errorf("UpdateFoodPath success,diseaseID:%v",diseaseID)
	return err
}