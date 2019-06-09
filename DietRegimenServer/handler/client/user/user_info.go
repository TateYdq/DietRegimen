package user

import (
	"encoding/hex"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/helper"
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

type UpdateUserInfoRequest struct {
	Token string `json:"token"`
	UserInfo database.UserInfo `json:"user_info"`
}
//需验证token
func GetUserInfo(c *gin.Context){
	defer func() {
		recover()
	}()
	success,userID:= helper.VerifyToken(c)
	if !success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	userInfo,err := database.GetUserInfoByUserID(userID)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Success,
			"user_info":userInfo,
		})
		return
	}
}
//需验证token
func UpdateUserInfo(c *gin.Context){
	defer func() {
		recover()
	}()
	success,userID:= helper.VerifyToken(c)
	if !success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	var request UpdateUserInfoRequest
	err := c.BindJSON(&request)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.ServerError,
		})
		logrus.WithError(err).Errorf("BindJson error")
		return
	}
	userInfo := request.UserInfo
	userInfo.UserID = userID
	err = database.UpdateUserInfo(userInfo)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Success,
		})
		return
	}
}

func UploadUserImage(c *gin.Context){
	defer func() {
		recover()
	}()
	success,userID:= helper.VerifyToken(c)
	if !success{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Forbidden,
		})
		return
	}
	//得到上传的文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": utils.Failed,
		})
		return
	}
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
	err = database.UpdateUserPath(userID,path)
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
