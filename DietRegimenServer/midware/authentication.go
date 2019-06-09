package midware

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminVerify(c *gin.Context){
	token := c.GetHeader("Token")
	res := database.VerifyAdmin(token)
	if res == true{
		c.Next()
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		c.Abort()
	}
}
