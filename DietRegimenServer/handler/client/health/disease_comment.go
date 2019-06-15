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

type DiseaseCommentResponse struct {
	ID int `json:"id"`
	DiseaseID int `json:"disease_id"`
	UserID int `json:"user_id"`
	UserName string `json:"user_name"`
	Comment string `json:"comment"`
	RecordTime string `json:"record_time"`
	UserImagePath string `json:"user_image_path"`
}

type CommentDiseaseRequest struct{
	DiseaseID  int `json:"disease_id"`
	Content string `json:"content"`
}
func CommentDisease(c *gin.Context) {
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
	var cdq CommentDiseaseRequest
	err := c.BindJSON(&cdq)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	err = database.CreateDiseaseComment(cdq.DiseaseID,userID,cdq.Content)
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

func GetComment(c *gin.Context){
	defer func() {
		recover()
	}()
	diseaseID,err := helper.GetDiseaseID(c)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Failed,
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
	var responseList[] DiseaseCommentResponse
	for _,comment := range comments {
		response := DiseaseCommentResponse{
			ID: comment.ID,
			DiseaseID: comment.DiseaseID,
			UserID: comment.UserID,
			UserName: comment.UserName,
			Comment: comment.Comment,
			RecordTime: comment.RecordTime,
		}
		userImagePath,err:= database.GetUserImagePath(comment.UserID)
		if err == nil{
			response.UserImagePath = userImagePath
		}
		responseList = append(responseList, response)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": utils.Success,
		"comment_list":responseList,
	})
	return
}