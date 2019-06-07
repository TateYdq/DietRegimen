package server

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)


func AddFoodKind(c *gin.Context){
	var rq database.FoodKindInfo
	err := c.BindJSON(&rq)
	if err != nil{
		logrus.WithError(err).Error("AddFoodKind BindJSON failed")
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	foodID,err := database.CreateFoodKindAdmin(rq)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"food_id":foodID,
	})
	return
}


func UpdateFoodKind(c *gin.Context){
	var rq database.FoodKindInfo
	err := c.BindJSON(&rq)
	if err != nil{
		logrus.WithError(err).Error("UpdateFoodKind BindJSON failed")
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	err = database.UpdateFoodKindInfo(rq)
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