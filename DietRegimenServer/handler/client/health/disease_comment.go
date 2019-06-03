package health

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DiseaseCommentInfo struct{
	ID int `json:"id"`
	DiseaseID int `json:"disease_id"`
	UserID int `json:"user_id"`
	UserName string `json:"user_name"`
	Comment string `json:"comment"`
	RecordTime string `json:"record_time"`
}

type CommentDiseaseRequest struct{
	DiseaseID  int `json:"disease_id"`
	Content string `json:"content"`
}
func CommentDisease(c *gin.Context) {
	defer func() {
		recover()
		c.JSON(http.StatusOK,gin.H{
			"code":utils.ServerError,
		})
		return
	}()
	if success := helper.VerifyToken(c);!success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	var cdq CommentDiseaseRequest
	err := c.BindJSON(&cdq)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	userID,err := helper.GetUserID(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	database.CreateDiseaseComment(cdq.DiseaseID,userID,cdq.Content)

}

func GetComment(c *gin.Context){
	defer func() {
		recover()
		c.JSON(http.StatusOK,gin.H{
			"code":utils.ServerError,
		})
		return
	}()
	diseaseID,err := helper.GetDiseaseID(c)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Failed,
		})
		return
	}
	if success := helper.VerifyToken(c);!success{
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Forbidden,
		})
		return
	}

	comments,err := database.GetCommentByDiseaseID(diseaseID)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Failed,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": utils.Success,
		"comment_list":comments,
	})
	return
}