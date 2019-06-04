package user

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserCollectDiseaseInfo struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	DiseaseID int `json:"disease_id"`
	RecordTime string `json:"record_time"`
}
type CollectDiseaseRequest struct {
	Token string `json:"token"`
	DiseaseID int `json:"disease_id"`
}

func CollectDisease(c *gin.Context){
	defer func() {
		recover()
	}()
	success,userID:= helper.VerifyToken(c)
	if !success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	var request CollectDiseaseRequest
	err := c.BindJSON(&request)
	if err != nil{
		logrus.WithError(err).Errorf("BindJson error")
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Failed,
		})
		return
	}
	err = database.CreateCollectDiseaseInfo(userID,request.DiseaseID)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Failed,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": utils.Success,
	})
	return
}


func GetCollectDisease(c *gin.Context){
	defer func() {
		recover()
	}()
	success,userID:= helper.VerifyToken(c)
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