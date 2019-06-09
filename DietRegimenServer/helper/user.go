package helper

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/cache"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


//获取用户ID
func GetUserID(token string)(userID int,err error){
	userID,err = cache.GetInt(token)
	if err != nil{
		logrus.WithError(err).Errorf("GetUserID err,token is %v",token)
	}
	return userID,err
}


//验证token
func VerifyToken(c *gin.Context)(success bool,userID int){
	success = false
	token := c.GetHeader("Token")
	logrus.Infof("token is %v",token)
	if token == ""{
		return success,0
	}
	//测试环境，123为测试token,id为2的用户为测试用户
	if utils.IsStagingENV(){
		if token == "123"{
			return true,2
		}
	}
	userID,err := GetUserID(token)
	if err == nil{
		success = true
	}
	return success,userID
}
