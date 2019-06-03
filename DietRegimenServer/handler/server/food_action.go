package server

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func AddFood(c *gin.Context){
	if success := helper.VerifyAdminToken();!success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	logrus.Infof("verify success")
	var rq database.FoodInfo
	err := c.BindJSON(&rq)
	if err != nil{
		logrus.WithError(err).Error("AddFood BindJSON failed")
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	foodID,err := database.CreateFoodAdmin(rq)
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


func UpdateFood(c *gin.Context){
	var rq database.FoodInfo
	if success := helper.VerifyAdminToken();!success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	logrus.Infof("verify success")
	err := c.BindJSON(&rq)
	if err != nil{
		logrus.WithError(err).Error("UpdateFood BindJSON failed")
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	err = database.UpdateFoodInfo(rq)
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