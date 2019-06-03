package server

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func AddDisease(c *gin.Context){
	if success := helper.VerifyAdminToken();!success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	logrus.Infof("verify success")
	var rq database.DiseaseInfo
	err := c.BindJSON(&rq)
	if err != nil{
		logrus.WithError(err).Error("AddDisease BindJSON failed")
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	diseaseID,err := database.CreateDiseaseAdmin(rq)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"disease_id":diseaseID,
	})
	return
}


func UpdateDisease(c *gin.Context){
	var rq database.DiseaseInfo
	if success := helper.VerifyAdminToken();!success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	logrus.Infof("verify success")
	err := c.BindJSON(&rq)
	if err != nil{
		logrus.WithError(err).Error("UpdateDisease BindJSON failed")
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	err = database.UpdateDiseaseInfo(rq)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
	})
	return
}