package database

import (
	"errors"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/sirupsen/logrus"
)

type FoodInfo struct {
	ID int
	FoodID int `json:"food_id"`
	Name string `json:"name"`
	FoodKindID int `json:"food_kind_id"`
	FoodKind  string `json:"food_kind"`
	Info   string `json:"info"`
	Effect  string `json:"effect"`
	Keyword  string `json:"keyword"`
	Taboo string `json:"taboo"`
	ViewCount int `json:"view_count"`
	CollectCount int `json:"collect_count"`
	PhotoPath  string `json:"photo_path"`
	VoicePath string `json:"voice_path"`
}

func CreateFoodAdmin(request FoodInfo)(int,error){
	db := DrDatabase.Create(&request)
	err := db.Error
	if  err != nil {
		logrus.WithError(err).Error("CreateFoodAdmin failed")
	}
	foodID := request.ID
	return foodID,err
}


func GetFoodInfoByID(foodID int)(foodInfo FoodInfo,err error){
	err = DrDatabase.Model(FoodInfo{}).Where("food_id = ?",foodID).First(&foodInfo).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetFoodInfoByID err,foodID:%v",foodID)
	}
	return foodInfo,err
}

func UpdateFoodInfo(request FoodInfo)(err error){
	if request.FoodID == 0{
		logrus.Errorf("foodID is equals to 0")
		return errors.New("foodID is equals to 0")
	}
	record := utils.ScanStruct(request,[]string{"foodID"})
	err = DrDatabase.Model(FoodInfo{}).Where("food_id = ?",request.FoodID).Updates(record).Error
	if err != nil{
		logrus.WithError(err).Errorf("UpdateFoodInfo err,foodID:%v",request.FoodID)
	}
	return err

}
