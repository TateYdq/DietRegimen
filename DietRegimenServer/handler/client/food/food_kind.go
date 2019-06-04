package food

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
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


func GetFoodCategory(c *gin.Context){
	defer func() {
		recover()
	}()
	foodCategory,err:= database.GetFoodKind()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Failed,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"food_category":foodCategory,
	})
	return
}
