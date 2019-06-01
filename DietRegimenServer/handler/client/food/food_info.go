package food

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func GetFoodDetails(c *gin.Context){
	defer func() {
		recover()
		c.JSON(http.StatusOK,gin.H{
			"code":utils.ServerError,
		})
		return
	}()
	foodID,err := helper.GetFoodID(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	userInfo,err := database.GetFoodInfoByID(foodID)
	//TODO:ViewCount计数,CollectCount计数
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Success,
			"user_info":userInfo,
		})
		return
	}
}
