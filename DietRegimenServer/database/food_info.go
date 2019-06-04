package database

import (
	"errors"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/sirupsen/logrus"
)

type FoodInfo struct {
	FoodID int `json:"food_id" gorm:"column:food_id;primary_key"`
	Name string `json:"name"`
	FoodKindID int `json:"food_kind_id"`
	FoodKind  string `json:"food_kind"`
	Info   string `json:"info"`
	Effect  string `json:"effect"`
	Keyword  string `json:"keyword"`
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
	foodID := request.FoodID
	return foodID,err
}


func GetFoodInfoByID(foodID int)(foodInfo FoodInfo,err error){
	err = DrDatabase.Model(FoodInfo{}).Where("food_id = ?",foodID).First(&foodInfo).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetFoodInfoByID err,foodID:v",foodID)
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

func SearchByFoodKindID(foodKindID int)(foodList[] FoodInfo,err error){
	err = DrDatabase.Model(FoodInfo{}).Where("food_kind_id = ?",foodKindID).Scan(&foodList).Error
	if err != nil{
		logrus.WithError(err).Errorf("SearchByFoodKindID err,foodKindID:%v",foodKindID)
	}
	return foodList,err
}

func SearchByFoodKindIDAndKey(keyword string,foodKindID int)(foodList[] FoodInfo,err error){
	keyword = "%"+keyword+"%"
	err = DrDatabase.Model(FoodInfo{}).Where("food_kind_id = ? and (food_kind like ? or info like ?  or keyword like ? or effect like ? or name like ?) ",foodKindID,keyword,keyword,keyword,keyword,keyword).Scan(&foodList).Error
	if err != nil{
		logrus.WithError(err).Errorf("SearchByFoodKindIDAndKey err,foodKindID:v,keyword:v",foodKindID,keyword)
	}
	return foodList,err
}

func SearchByKeyWord(keyword string)(foodList[] FoodInfo,err error){
	if keyword == ""{
		err = DrDatabase.Model(FoodInfo{}).Scan(&foodList).Error
		if err != nil{
			logrus.WithError(err).Errorf("SearchByKeyWord err")
		}
	}else {
		keyword = "%"+keyword+"%"
		err = DrDatabase.Model(FoodInfo{}).Where("food_kind like ? or info like ? or keyword like ? or effect like ? or name like ? ",keyword,keyword,keyword,keyword,keyword).Scan(&foodList).Error
		if err != nil {
			logrus.WithError(err).Errorf("SearchByKeyWord err,keyword:v", keyword)
		}
	}
	return foodList,err
}