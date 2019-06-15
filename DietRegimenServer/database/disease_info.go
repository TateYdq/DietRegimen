package database

import (
	"errors"
	"fmt"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/ai"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/cache"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/sirupsen/logrus"
)

type DiseaseInfo struct{
	DiseaseID int `json:"disease_id"  gorm:"column:disease_id;primary_key"`
	Name string `json:"name"`
	DiseaseKind string `json:"disease_kind"`
	Info string `json:"info"`
	Taboo string `json:"taboo"`
	RecommendFood string `json:"recommend_food"`
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
	//语音识别
	go func() {
		content := ""
		if request.Taboo !="" {
			content = "禁忌: "+request.Taboo+" "
		}
		if request.RecommendFood != ""{
			content = "推荐食物: "+request.RecommendFood+" "
		}
		if request.Info != ""{
			content = "介绍: "+request.Info
		}
		path, err := ai.CreateVoice("disease", request.DiseaseID,content)
		if err != nil{
			logrus.WithError(err).Error("CreateDiseaseVoice failed")
			return
		}
		err = UpdateDiseaseField(diseaseID,path,"voice_path")
		if err != nil{
			logrus.WithError(err).Error("UpdateVoice failed")
			return
		}
	}()
	return diseaseID,err
}


func GetDiseaseInfoByID(diseaseID int)(diseaseInfo DiseaseInfo,err error){
	err = DrDatabase.Model(DiseaseInfo{}).Where("disease_id = ?",diseaseID).First(&diseaseInfo).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetDiseaseInfoByID err,diseaseID:%v",diseaseID)
	}
	go func() {
		err := DrDatabase.Exec("update disease_info set view_count=view_count+1 where disease_id = ?",diseaseID).Error
		if err != nil{
			logrus.Errorf("UpdateDiseaseViewCount err")
		}
	}()
	return diseaseInfo,err
}

func GetDiseaseLists(keyword string)(diseaseList[] DiseaseInfo,err error){
	if keyword == ""{
		err = DrDatabase.Model(DiseaseInfo{}).Order("view_count desc").Scan(&diseaseList).Error
		if err != nil{
			logrus.WithError(err).Errorf("GetDiseaseLists err")
		}
	}else{
		keyword = "%"+keyword+"%"
		err = DrDatabase.Model(DiseaseInfo{}).Order("view_count desc").Where("disease_kind like ? or info like ? or taboo like ?  or name like ? ",keyword,keyword,keyword,keyword).Scan(&diseaseList).Error
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

func UpdateDiseaseField(diseaseID int,path string,field string)(err error){
	if diseaseID == 0{
		logrus.Errorf("diseaseID is equals to 0")
		return errors.New("diseaseID is equals to 0")
	}
	record := make(map[string]interface{})
	record[field] = path
	err = DrDatabase.Model(DiseaseInfo{}).Where("disease_id = ?",diseaseID).Updates(record).Error
	if err != nil{
		logrus.WithError(err).Errorf("UpdateDiseaseField err,diseaseID:%v",diseaseID)
	}
	logrus.Infof("UpdateDiseaseField success,diseaseID:%v",diseaseID)
	return err
}

func UpdateDiseaseCollectCount(diseaseID int){
	if diseaseID == 0{
		logrus.Errorf("diseaseID is equals to 0")
		return
	}
	err := DrDatabase.Exec("update disease_info set collect_count=collect_count+1 where disease_id = ?",diseaseID).Error
	if err != nil{
		logrus.Errorf("UpdateDiseaseCollectCount err")
	}
}

func UpdateDiseaseView(diseaseID int)(){
	if diseaseID == 0{
		logrus.Errorf("diseaseID is equals to 0")
		return
	}
	err := DrDatabase.Exec("update disease_info set view_count=view_count+1 where disease_id = ?",diseaseID).Error
	if err != nil{
		logrus.Errorf("UpdateDiseaseView err")
	}
}

func GetDiseaseNameByDiseaseID(diseaseID int)(string){
	key := fmt.Sprintf(cache.DiseaseIDToNameKey,diseaseID)
	value,err := cache.Get(key)
	if err != nil || value == ""{
		var diseaseInfo DiseaseInfo
		err := DrDatabase.Model(DiseaseInfo{}).Where("disease_id = ?",).First(&diseaseInfo).Error
		if err != nil{
			logrus.WithError(err).Errorf("GetDiseaseNameByDiseaseID err,diseaseID:%v",diseaseID)
			return ""
		}
		go cache.Set(key,diseaseInfo.Name,0)
		return diseaseInfo.Name
	}else{
		return value
	}
}

func DecreaseDiseaseCollectCount(diseaseID int){
	err := DrDatabase.Exec("update disease_info set collect_count = collect_count-1 where disease_id = ? and collect_count > 0",diseaseID).Error
	if err != nil{
		logrus.WithError(err).Errorf("DecreaseCollectCount err,diseaseID = %v",diseaseID)
	}
}

func CreateDiseaseVoice() {
	var infoList[] DiseaseInfo
	DrDatabase.Model(&DiseaseInfo{}).Scan(&infoList)
	for _,info:= range infoList{
		content := ""
		if info.Taboo !="" {
			content = "禁忌: "+info.Taboo+" "
		}
		if info.RecommendFood != ""{
			content = "推荐食物: "+info.RecommendFood+" "
		}
		if info.Info != ""{
			content = "介绍: "+info.Info
		}
		path, err := ai.CreateVoice("disease", info.DiseaseID,content)
		if err != nil {
			logrus.WithError(err).Error("CreateDiseaseVoice failed")
			return
		}
		err = UpdateDiseaseField(info.DiseaseID, path, "voice_path")
		if err != nil {
			logrus.WithError(err).Error("UpdateVoice failed")
			return
		}
	}

}