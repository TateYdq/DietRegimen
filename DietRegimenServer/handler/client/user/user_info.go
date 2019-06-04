package user

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UpdateUserInfoRequest struct {
	Token string `json:"token"`
	UserInfo database.UserInfo `json:"user_info"`
}
//需验证token
func GetUserInfo(c *gin.Context){
	defer func() {
		recover()
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
	defer func() {
		recover()
	}()
	if success := helper.VerifyToken(c);!success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	var request UpdateUserInfoRequest
	err := c.BindJSON(&request)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.ServerError,
		})
		logrus.WithError(err).Errorf("BindJson error")
		return
	}
	userID,err := helper.GetUserID(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	userInfo := request.UserInfo
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
