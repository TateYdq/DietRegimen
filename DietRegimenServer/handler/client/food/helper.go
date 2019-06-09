package food

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
)


func KindIDAndKeySearch(keyword string,kindID int)(foodList[] database.FoodInfo,err error){
	return database.SearchByFoodKindIDAndKey(keyword,kindID)
}


func KindIDSearch(kindiD int)(foodList[] database.FoodInfo,err error){
	return database.SearchByFoodKindID(kindiD)
}

func KeySearch(keyword string)(foodList[] database.FoodInfo,err error){
	return database.SearchByKeyWord(keyword)
}

