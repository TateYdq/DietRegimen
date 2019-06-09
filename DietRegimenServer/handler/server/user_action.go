package server

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)



func AddUser(c *gin.Context){
	var rq database.UserInfo
	err := c.BindJSON(&rq)
	if err != nil{
		logrus.WithError(err).Error("AddUser BindJSON failed")
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	userID,err := database.CreateUserAdmin(rq)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"user_id":userID,
	})
	return
}

func UpdateUser(c *gin.Context){
	var rq database.UserInfo
	err := c.BindJSON(&rq)
	if err != nil{
		logrus.WithError(err).Error("UpdateUser BindJSON failed")
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	err = database.UpdateUserInfo(rq)
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