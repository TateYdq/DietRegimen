package database

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/sirupsen/logrus"
)

type UserFoodRelation struct {
	UserID int `json:"user_id"`
	FoodID int `json:"food_id"`
	FoodName string `json:"food_name"`
	Score int `json:"score"`
}


func CreateUserAndFoodScore(userID int,foodID int){
	foodName := GetFoodNameByFoodID(foodID)
	relation := UserFoodRelation{
		UserID:userID,
		FoodID:foodID,
		FoodName:foodName,
		Score:0,
	}
	err := DrDatabase.Create(&relation).Error
	if err != nil{
		logrus.WithError(err).Errorf("CreateUserAndFoodScore,userID:%v,foodID:%v",userID,foodID)
	}
}

func AddUserAndFoodScore(userID int,foodID int,score int)(err error){
	var count int
	DrDatabase.Model(UserFoodRelation{}).Where("user_id = ? and food_id = ? ",userID,foodID).Count(&count)
	if count == 0{
		CreateUserAndFoodScore(userID,foodID)
	}
	err = DrDatabase.Exec("update user_food_relation set score = score + ? where user_id = ? and food_id = ? ",score,userID,foodID).Error
	if err != nil{
		logrus.WithError(err).Errorf("AddUserAndFoodScore,userID:%v,foodID:%v",userID,foodID)
	}
	return err
}


func GetRecFoodByUserID(userID int)(foodInfo[] FoodInfo,err error){
	var relations []UserFoodRelation
	var count int
	db := DrDatabase.Model(UserFoodRelation{}).Where("user_id = ? ",userID).Order("score desc").Count(&count)
	if count < 3{
		err = DrDatabase.Model(FoodInfo{}).Order("collect_count desc").Limit(utils.RecNum).Scan(&foodInfo).Error
		if err != nil {
			logrus.WithError(err).Errorf("GetRecFoodByUserID,userID:%v",userID)
		}
		return foodInfo,err
	}else{
		err := db.Limit(utils.RecNum).Scan(&relations).Error
		if err != nil {
			logrus.WithError(err).Errorf("GetRecFoodByUserID,userID:%v",userID)
			return foodInfo,err
		}
		var foodIDs []int
		for _,relation := range relations{
			foodIDs = append(foodIDs, relation.FoodID)
		}

		err = DrDatabase.Model(FoodInfo{}).Where("food_id in (?)",foodIDs).Scan(&foodInfo).Error
		if err != nil {
			logrus.WithError(err).Errorf("GetRecFoodByUserID,userID:%v",userID)
		}
		return foodInfo,err

	}
}