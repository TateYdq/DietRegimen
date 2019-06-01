package database

import "github.com/sirupsen/logrus"

type FoodInfo struct {
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


func GetFoodInfoByID(foodID int)(foodInfo FoodInfo,err error){
	err = DrDatabase.Where("food_id="+string(foodID)).First(&foodInfo).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetFoodInfoByID err,foodID:%v",foodID)
	}
	return foodInfo,err
}