package database

import "github.com/sirupsen/logrus"

type DiseaseFoodRec struct {
	DiseaseID int `json:"disease_id"`
	FoodID int `json:"food_id"`
	FoodName string `json:"food_name"`
}


func AddFoodRecs(diseaseID int,foodNames[] string){
	for _,foodName := range foodNames{
		object := DiseaseFoodRec{
			DiseaseID:diseaseID,
			FoodID:GetFoodIDByFoodName(foodName),
			FoodName:foodName,
		}
		err := DrDatabase.Create(&object).Error
		if err != nil{
			logrus.WithError(err).Errorf("AddFoodRecs err,diseaseID:%v,foodName:%v",diseaseID,foodName)
		}
	}
}