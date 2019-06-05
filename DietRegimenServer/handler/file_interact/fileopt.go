package file_interact

import (
	"encoding/hex"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)
/**文件上传下载操作页面**/
func Fileopthtml(c *gin.Context){
	c.HTML(http.StatusOK, "fileopt.html", gin.H{
		"title": "GIN: 文件上传下载操作布局页面",
	})
}


func UploadTable(num int,id int,path string)(err error){
	if num == TypeUser{
		return database.UpdateUserPath(id,path)
	} else if num == TypeFood{
		return database.UpdateFoodPath(id,path)
	}else if num == TypeKind{
		return database.UpdateFoodKindPath(id,path)
	}else if num == TypeDisease{
		return database.UpdateDiseasePath(id,path)
	}
	return nil
}


/**上传方法**/
func Fileupload(c *gin.Context){
	//得到上传的文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": utils.Failed,
		})
		return
	}
	numGet := c.Request.Form["num"][0]
	num,err := strconv.Atoi(numGet)
	if err != nil || num == 0{
		c.JSON(http.StatusOK,gin.H{
			"code": utils.Failed,
			"msg":"please give num",
		})
		return
	}
	idGet := c.Request.Form["id"][0]
	id,err := strconv.Atoi(idGet)
	if err != nil || id == 0{
		c.JSON(http.StatusOK,gin.H{
			"code": utils.Failed,
			"msg":"please give id",
		})
		return
	}
	logrus.Infof("upload num:%v,id:%v",num,id)
	//文件的名称
	filename := header.Filename
	var suffix string
	if index := strings.LastIndex(filename,".");index != -1{
		suffix = filename[index:]
	}else{
		suffix = ""
	}
	filename = hex.EncodeToString([]byte(strconv.FormatInt(time.Now().Unix(),10)))
	path := filename + suffix
	logrus.WithField("path",path)
	//创建文件
	out, err := os.Create(utils.StaticUploadPath+path)
	if err != nil {
		logrus.WithError(err).Errorf("create file err")
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	err = UploadTable(num,id,path)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code": utils.Failed,
			"msg":"update table err",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code": utils.Success,
		"path":path,
	})
	return
}

/**下载方法**/
func Filedown(c *gin.Context){
	//暂时没有提供方法
}
