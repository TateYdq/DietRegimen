package health

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDiseaseDetails(c *gin.Context){
	defer func() {
		recover()
	}()
	diseaseID,err := helper.GetDiseaseID(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	diseaseInfo,err := database.GetDiseaseInfoByID(diseaseID)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	go func() {
		database.UpdateDiseaseView(diseaseID)
		success,userID:= helper.VerifyToken(c)
		if !success{
			return
		}
		go database.AddUserScore(userID,utils.ScoreUserLook)
		go database.AddUserAndDiseaseScore(userID,diseaseID,utils.ScoreDiseaseLook)
	}()
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"disease_detail":diseaseInfo,
	})
	return
}

func GetDiseasesLists(c *gin.Context){
	defer func() {
		recover()
	}()
	keyword,_ := helper.GetKeyWord(c)
	list,err:= database.GetDiseaseLists(keyword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Failed,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"disease_list":list,
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
	diseaseID,err := helper.GetDiseaseID(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	result := database.IsUserCollectedDisease(userID,diseaseID)
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
	diseaseID,err := helper.GetDiseaseID(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	result := database.DeleteCollectedDisease(userID,diseaseID)
	go func() {
		if result{
			database.DecreaseDiseaseCollectCount(diseaseID)
		}
	}()
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"result":result,
	})
}