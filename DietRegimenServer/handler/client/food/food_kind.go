package food

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FoodKind struct {
	KindID int `json:"kind_id"`
	KindName string `json:"kind_name"`
	KindInfo string `json:"kind_info"`
	PhotoPath string `json:"photo_path"`
	ViewCount int `json:"view_count"`
}

func GetFoodKinds(c *gin.Context){
	defer func() {
		recover()
		c.JSON(http.StatusOK,gin.H{
			"code":utils.ServerError,
		})
		return
	}()
	kindID,err := helper.GetFoodKindID(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	foodKind,err := database.GetFoodKindByID(kindID)
	//TODO:ViewCount计数
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Success,
			"food_category":foodKind,
		})
		return
	}
}