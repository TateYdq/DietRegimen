package food

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
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
	var cfq CommentFoodRequest
	err := c.BindJSON(&cfq)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}

}

func GetComment(c *gin.Context){
	defer func() {
		recover()
		c.JSON(http.StatusOK,gin.H{
			"code":utils.ServerError,
		})
		return
	}()
	foodID,err := helper.GetFoodID(c)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Failed,
		})
		return
	}
	success := helper.VerifyToken(c)
	if !success{
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Forbidden,
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