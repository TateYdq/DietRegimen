package database

import "github.com/sirupsen/logrus"

type FoodKind struct {
	KindID int `json:"kind_id"`
	KindName string `json:"kind_name"`
	KindInfo string `json:"kind_info"`
	PhotoPath string `json:"photo_path"`
	ViewCount int `json:"view_count"`
}

func GetFoodKindByID(kindID int)(foodKind FoodKind,err error) {
	err = DrDatabase.Where("kind_id="+string(kindID)).First(&foodKind).Error
	if err != nil{
		logrus.WithError(err).Errorf("GetFoodKindByID err,kindID:%v",kindID)
	}
	return foodKind,err
}