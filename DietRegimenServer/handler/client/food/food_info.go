package food

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)
func GetFoodDetails(c *gin.Context){
	defer func() {
		recover()
	}()
	foodID,err := helper.GetFoodID(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	foodInfo,err := database.GetFoodInfoByID(foodID)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	go func() {
		database.UpdateFoodView(foodID)
		success,userID:= helper.VerifyToken(c)
		if !success{
			return
		}
		go database.AddUserScore(userID,utils.ScoreUserLook)
		go database.AddUserAndFoodScore(userID,foodID,utils.ScoreFoodLook)
	}()
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"food_detail":foodInfo,
	})
	return
}

func IsCollected(c *gin.Context){
	success,userID:= helper.VerifyToken(c)
	if !success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	foodID,err := helper.GetFoodID(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	result := database.IsUserCollectedFood(userID,foodID)
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"result":result,
	})
}

func CancelCollected(c *gin.Context){
	success,userID:= helper.VerifyToken(c)
	if !success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	foodID,err := helper.GetFoodID(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	result := database.DeleteCollectedFood(userID,foodID)
	go func() {
		if result{
			database.DecreaseFoodCollectCount(foodID)
		}
	}()
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"result":result,
	})
}
