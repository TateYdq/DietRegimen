package recommend

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRecInfo(c *gin.Context){
	success,userID:= helper.VerifyToken(c)
	if !success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	foodInfoList,err:= database.GetRecFoodByUserID(userID)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	diseaseInfoList,err:= database.GetRecDiseaseByUserID(userID)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"food_list":foodInfoList,
		"disease_list":diseaseInfoList,
	})
	return
}