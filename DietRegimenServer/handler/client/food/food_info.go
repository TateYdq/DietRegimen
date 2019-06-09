package food

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFoodDetails(c *gin.Context){
	defer func() {
		recover()
	}()
	foodID,err := helper.GetFoodID(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	userInfo,err := database.GetFoodInfoByID(foodID)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	go func() {
		database.UpdateFoodView(foodID)
	}()
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"user_info":userInfo,
	})
	return
}
