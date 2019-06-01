package health

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DiseaseInfo struct{
	DiseaseID int `json:"disease_id"`
	Name string `json:"name"`
	DiseaseKind string `json:"disease_kind"`
	Info string `json:"info"`
	PhotoPath string `json:"photo_path"`
	VoicePath string `json:"voice_path"`
	ViewCount int `json:"view_count"`
	CollectCount string `json:"collect_count"`
}

func GetDiseaseDetail(c *gin.Context){
	defer func() {
		recover()
		c.JSON(http.StatusOK,gin.H{
			"code":utils.ServerError,
		})
		return
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