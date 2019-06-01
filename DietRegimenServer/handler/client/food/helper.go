package food

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func GetKeyWords(c *gin.Context)(keyword string,err error){
	keyword = c.Query("keyword")
	if keyword == "" {
		return "",errors.New("get keywords error")
	}
	return keyword,nil
}
//TODO:搜索算法待写
func Search(keyword string)(foodList[] FoodInfo){
	return nil
}