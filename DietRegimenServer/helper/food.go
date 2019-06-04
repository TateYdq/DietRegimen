package helper

import (
	"github.com/gin-gonic/gin"
	"errors"
	"strconv"
)

const(
	MethodGet = "get"
	MethodPost = "post"

	Token = "token"
	SessionID = "sessionID"
)

//获取食物ID
func GetFoodID(c *gin.Context)(foodID int,err error){
	id := c.Query("food_id")
	if id == "" {
		return 0,errors.New("get food_id error")
	}
	foodID,err = strconv.Atoi(id)
	return foodID,err
}

//获取食物ID
func GetFoodKindID(c *gin.Context)(kindID int,err error){
	id := c.Query("food_kind_id")
	if id == "" {
		return 0,errors.New("get food_kind_id error")
	}
	kindID,err = strconv.Atoi(id)
	return kindID,err
}

//获取食物ID2
func GetKindID(c *gin.Context)(kindID int,err error){
	id := c.Query("kind_id")
	if id == "" {
		return 0,errors.New("get kind_id error")
	}
	kindID,err = strconv.Atoi(id)
	return kindID,err
}

