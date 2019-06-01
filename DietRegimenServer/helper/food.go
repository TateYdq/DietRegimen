package helper

import (
	"github.com/gin-gonic/gin"
	"errors"
)

const(
	MethodGet = "get"
	MethodPost = "post"

	Token = "token"
	SessionID = "sessionID"
)

//获取食物ID
func GetFoodID(c *gin.Context)(foodID int,err error){
	return 0,errors.New("invalid");
}

//获取食物ID
func GetFoodKindID(c *gin.Context)(kindID int,err error){
	return 0,errors.New("invalid");
}
