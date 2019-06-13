package database

import (
	"errors"
	"fmt"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/ai"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/cache"
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
	//语音识别
	go func() {
		path,err := ai.CreateVoice("food", foodID, request.Info)
		if err != nil{
			logrus.WithError(err).Error("CreateFoodVoice failed")
			return
		}
		err = UpdateFoodField(foodID,path,"voice_path")
		if err != nil{
			logrus.WithError(err).Error("UpdateVoice failed")
			return
		}
	}()
	return foodID,err
}


func GetFoodInfoByID(foodID int)(foodInfo FoodInfo,err error){
	err = DrDatabase.Model(FoodInfo{}).Where("food_id = ?",foodID).First(&foodInfo).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetFoodInfoByID err,foodID:v",foodID)
	}
	go func() {
		err := DrDatabase.Raw("update food_info set view_count=view_count+1 where food_id = ?",foodID).Error
		if err != nil{
			logrus.Errorf("UpdateFoodView err")
		}
	}()
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
	err = DrDatabase.Model(FoodInfo{}).Order("view_count desc").Where("food_kind_id = ?",foodKindID).Scan(&foodList).Error
	if err != nil{
		logrus.WithError(err).Errorf("SearchByFoodKindID err,foodKindID:%v",foodKindID)
	}
	return foodList,err
}

func SearchByFoodKindIDAndKey(keyword string,foodKindID int)(foodList[] FoodInfo,err error){
	keyword = "%"+keyword+"%"
	err = DrDatabase.Model(FoodInfo{}).Order("view_count desc").Where("food_kind_id = ? and (food_kind like ? or info like ?  or keyword like ? or effect like ? or name like ?) ",foodKindID,keyword,keyword,keyword,keyword,keyword).Scan(&foodList).Error
	if err != nil{
		logrus.WithError(err).Errorf("SearchByFoodKindIDAndKey err,foodKindID:%v,keyword:%v",foodKindID,keyword)
	}
	return foodList,err
}

func SearchByKeyWord(keyword string)(foodList[] FoodInfo,err error){
	if keyword == ""{
		err = DrDatabase.Model(FoodInfo{}).Order("view_count desc").Scan(&foodList).Error
		if err != nil{
			logrus.WithError(err).Errorf("SearchByKeyWord err")
		}
	}else {
		keyword = "%"+keyword+"%"
		err = DrDatabase.Model(FoodInfo{}).Order("view_count desc").Where("food_kind like ? or info like ? or keyword like ? or effect like ? or name like ? ",keyword,keyword,keyword,keyword,keyword).Scan(&foodList).Error
		if err != nil {
			logrus.WithError(err).Errorf("SearchByKeyWord err,keyword:%v", keyword)
		}
	}
	return foodList,err
}

func UpdateFoodField(foodID int, value string,field string)(err error){
	if foodID == 0{
		logrus.Errorf("foodID is equals to 0")
		return errors.New("foodID is equals to 0")
	}
	record := make(map[string]interface{})
	record[field] = value
	err = DrDatabase.Model(FoodInfo{}).Where("food_id = ?",foodID).Updates(record).Error
	if err != nil{
		logrus.WithError(err).Errorf("UpdateFoodField err,foodID:%v",foodID)
	}
	logrus.Infof("UpdateFoodField success,foodID:%v",foodID)
	return err
}

func UpdateFoodCollect(foodID int){
	if foodID == 0{
		logrus.Errorf("foodID is equals to 0")
		return
	}
	err := DrDatabase.Raw("update food_info set collect_count=collect_count+1 where food_id = ?",foodID).Error
	if err != nil{
		logrus.Errorf("UpdateFoodCollect err")
	}
}

func UpdateFoodView(foodID int)(){
	if foodID == 0{
		logrus.Errorf("foodID is equals to 0")
		return
	}
	err := DrDatabase.Raw("update food_info set view_count=view_count+1 where food_id = ?",foodID).Error
	if err != nil{
		logrus.Errorf("UpdateFoodView err")
	}
}

func GetFoodNameByFoodID(foodID int)(string){
	key := fmt.Sprintf(cache.FoodIDToNameKey,foodID)
	value,err := cache.Get(key)
	if err != nil || value == ""{
		var foodInfo FoodInfo
		err := DrDatabase.Model(FoodInfo{}).Where("food_id = ?",).First(&foodInfo).Error
		if err != nil{
			logrus.WithError(err).Errorf("GetFoodNameByFoodID err,foodID:%v",foodID)
			return ""
		}
		go cache.Set(key,foodInfo.Name,0)
		return foodInfo.Name
	}else{
		return value
	}
}

func GetFoodIDByFoodName(foodName string)(int){
	var foodInfo FoodInfo
	err := DrDatabase.Model(FoodInfo{}).Where("name = ?",foodName).First(&foodInfo).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetFoodIDByFoodName err,name:%v",foodName)
		return 0
	}
	return foodInfo.FoodID
}

