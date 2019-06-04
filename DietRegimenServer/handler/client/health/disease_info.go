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
	//TODO:ViewCount计数,CollectCount计数
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Success,
			"disease_detail":diseaseInfo,
		})
		return
	}
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