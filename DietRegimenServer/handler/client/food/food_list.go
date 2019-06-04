package food

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func SearchFood(c *gin.Context){
	defer func() {
		recover()
	}()
	keyword,err1 := helper.GetKeyWord(c)
	kindID,err2 := helper.GetKindID(c)
	var err error
	var foodList[] database.FoodInfo
	if err1 == nil{
		if err2 == nil{
			foodList,err = KindIDAndKeySearch(keyword,kindID)
		}else{
			foodList,err = KeySearch(keyword)
		}
	}else{
		if err2 == nil{
			foodList,err = KindIDSearch(kindID)
		}else{
			foodList,err = KeySearch("")
		}
	}
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		logrus.WithError(err).Errorf("search error")
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"food_list":foodList,
	})
	return
}
