package helper

import (
	"github.com/gin-gonic/gin"
)


//获取用户ID
func GetUserID(c *gin.Context)(userID int,err error){
	return 1,nil
	//return 0,errors.New("invalid");
}
func GetToken(c *gin.Context)(string){
	//if methodName == MethodGet{
	token := c.Query(Token)
	return token
	//}else{
	//	TOKEN
	//}
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
