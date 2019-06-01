package user

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CollectDisease struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	DiseaseID int `json:"disease_id"`
	RecordTime string `json:"record_time"`
}


func GetCollectDisease(c *gin.Context){
	defer func() {
		recover()
		c.JSON(http.StatusOK,gin.H{
			"code":utils.ServerError,
		})
		return
	}()
	userID,err := helper.GetUserID(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	success := helper.VerifyToken(c)
	if !success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	collectDiseases, err := database.GetCollectDiseaseByUserID(userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": utils.ServerError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":         utils.Success,
		"collect_info": collectDiseases,
	})
	return
}