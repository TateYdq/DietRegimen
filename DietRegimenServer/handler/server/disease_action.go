package server

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)



func AddDisease(c *gin.Context){
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

type DiseaseFoodRecRequest struct {
	DiseaseID int `json:"disease_id"`
	FoodNames[] string `json:"food_names"`
}

func AddDiseaseFoodRec(c *gin.Context){
	var request DiseaseFoodRecRequest
	err := c.BindJSON(&request)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	database.AddFoodRecs(request.DiseaseID,request.FoodNames)
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
	})
	return
}