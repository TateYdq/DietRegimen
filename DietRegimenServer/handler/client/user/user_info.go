package user

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserInfo struct{
	UserID int `json:"user_id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	UserImagePath string `json:"user_image_path"`
	DiseasesHistory string `json:"diseases_focus"`
	Keywords string `json:"keywords"`
}
//需验证token
func GetUserInfo(c *gin.Context){
	defer func() {
		recover()
		c.JSON(http.StatusOK,gin.H{
			"code":utils.ServerError,
		})
		return
	}()
	userID,err := helper.GetUserID(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}
	success := helper.VerifyToken(c)
	if !success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	userInfo,err := database.GetOrCreateUserInfoByOpenID(userID)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Success,
			"user_info":userInfo,
		})
		return
	}
}
//需验证token
func UpdateUserInfo(c *gin.Context){
	var userInfo database.UserInfo
	err := c.BindJSON(&userInfo)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.ServerError,
		})
		logrus.WithError(err).Errorf("params error")
		return
	}
	userID,err := helper.GetUserID(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	success := helper.VerifyToken(c)
	if !success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	userInfo.UserID = userID
	err = database.UpdateUserInfo(userInfo)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Success,
		})
		return
	}
}
