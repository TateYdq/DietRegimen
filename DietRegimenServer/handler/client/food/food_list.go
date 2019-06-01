package food

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchFood(c *gin.Context){
	defer func() {
		recover()
		c.JSON(http.StatusOK,gin.H{
			"code":utils.ServerError,
		})
		return
	}()
	keyword,err := GetKeyWords(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	foodList := Search(keyword)
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"food_list":foodList,
	})
	return
}
