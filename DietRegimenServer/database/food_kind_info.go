package database

import (
	"errors"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/sirupsen/logrus"
)

type FoodKindInfo struct {
	ID int
	KindID int `json:"kind_id"`
	KindName string `json:"kind_name"`
	KindInfo string `json:"kind_info"`
	PhotoPath string `json:"photo_path"`
	ViewCount int `json:"view_count"`
}

func CreateFoodKindAdmin(request FoodKindInfo)(int,error){
	db := DrDatabase.Create(&request)
	err := db.Error
	if  err != nil {
		logrus.WithError(err).Error("CreateFoodKindAdmin failed")
	}
	kindID := request.ID
	return kindID,err
}


func GetFoodKindByID(kindID int)(foodKind FoodKindInfo,err error) {
	err = DrDatabase.Model(FoodKindInfo{}).Where("kind_id = ? ",kindID).First(&foodKind).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetFoodKindByID err,kindID:%v",kindID)
	}
	return foodKind,err
}

func UpdateFoodKindInfo(request FoodKindInfo)(err error){
	if request.KindID == 0{
		logrus.Errorf("kindID is equals to 0")
		return errors.New("kindID is equals to 0")
	}
	record := utils.ScanStruct(request,[]string{"kindID"})
	err = DrDatabase.Model(FoodKindInfo{}).Where("kind_id = ?",request.KindID).Updates(record).Error
	if err != nil{
		logrus.WithError(err).Errorf("UpdateFoodKindInfo err,kindID:%v",request.KindID)
	}
	return err

}
