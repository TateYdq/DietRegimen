package recommend

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SubmitQuestionnaireRequest struct {
	AnswerSheets[] database.AnswerSheet `json:"answer_sheets"`
}
func SubmitQuestionnaire(c *gin.Context){
	success,userID:= helper.VerifyToken(c)
	if !success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	var request SubmitQuestionnaireRequest
	err := c.BindJSON(&request)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	go func() {
		for _,answer := range request.AnswerSheets{
			//ID号明确，防止越权。
			answer.UserID = userID
			database.CreateQuestionSheet(answer)
		}
	}()
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
	})
	return
}

func GetQuestionnaire(c *gin.Context){
	success,userID:= helper.VerifyToken(c)
	if !success{
		userID = 0
	}
	questionLists,err := database.SelectQuestionInfo(userID)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"question_lists":questionLists,
	})
	return
}

