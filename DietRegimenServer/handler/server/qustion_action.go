package server

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func AddQuestion(c *gin.Context){
	var request database.QuestionInfo
	err := c.BindJSON(&request)
	if err != nil{
		logrus.WithError(err).Errorf("BindJSON err")
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	questionID,err := database.AddQuestion(request)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"question_id":questionID,
	})
	return
}