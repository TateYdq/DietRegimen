package helper

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/cache"
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
	if token == ""{
		return success,0
	}
	logrus.Infof("token is %v",token)
	userID,err := GetUserID(token)
	if err == nil{
		success = true
	}
	return success,userID
}
