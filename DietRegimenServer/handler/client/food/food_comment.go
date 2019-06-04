package food

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type FoodComment struct{
	ID int `json:"id"`
	FoodID int `json:"food_id"`
	UserID int `json:"user_id"`
	UserName string `json:"user_name"`
	Comment string `json:"comment"`
	RecordTime string `json:"record_time"`
}
type CommentFoodRequest struct{
	FoodID  int    `json:"food_id"`
	Content string `json:"content"`
}
func CommentFood(c *gin.Context) {
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

	var cfq CommentFoodRequest
	err := c.BindJSON(&cfq)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		logrus.WithError(err).Errorf("BindJson error")
		return
	}
	err = database.CreateFoodComment(cfq.FoodID,userID,cfq.Content)
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
	foodID,err := helper.GetFoodID(c)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Failed,
		})
		return
	}
	foodComments,err := database.GetCommentByFoodID(foodID)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Failed,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": utils.Success,
		"comment_list":foodComments,
	})
	return
}