package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func GetKeyWord(c *gin.Context)(keyword string,err error){
	keyword = c.Query("keyword")
	if keyword == "" {
		return "",errors.New("get keywords error")
	}
	return keyword,nil
}

