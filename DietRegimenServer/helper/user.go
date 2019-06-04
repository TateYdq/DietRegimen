package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


//获取用户ID
func GetUserID(c *gin.Context)(userID int,err error){
	return 3,nil
	//return 0,errors.New("invalid");
}
func GetToken(c *gin.Context)(string){
	token := c.Query(Token)
	logrus.Infof("token is %v",token)
	return token
}

func GetSession(c *gin.Context)(string){
	sessionID := c.GetHeader(SessionID)
	return sessionID
}


//验证token

func VerifyToken(c *gin.Context)(success bool){
	GetToken(c)
	return true
}
