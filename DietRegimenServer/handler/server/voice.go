package server

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/gin-gonic/gin"
)

//更换所有音频
func CreateAllVoice(c *gin.Context){
	database.CreateFoodVoice()
	database.CreateDiseaseVoice()
}

