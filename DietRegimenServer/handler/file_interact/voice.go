package file_interact

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func GetVoice(c *gin.Context){
	filename := c.Query("path")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./static/voice/"+filename)
}