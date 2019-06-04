package user

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/secret"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	logger "github.com/sirupsen/logrus"
)

var(
	WechatUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code"
)
type WechatLoginRequestBody struct {
	Code int `json:"code"`
}

type WechatLoginReponseBody struct {
	Openid string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid string `json:"unionid"`
	Errcode int `json:"errcode"`
	Errmsg string `json:"errmsg"`
}

func UserLogin(c *gin.Context){
	defer func() {
		recover()
	}()
	var wxReqBody WechatLoginRequestBody
	err := c.BindJSON(&wxReqBody)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		logger.WithError(err).Errorf("params error")
		return
	}

	//小程序验证请求
	wxUrl := fmt.Sprintf(WechatUrl,secret.AppID,secret.AppSecret,wxReqBody.Code)
	wxClient := &http.Client{}
	wxResp,err := wxClient.Get(wxUrl)
	defer wxResp.Body.Close()
	body, err := ioutil.ReadAll(wxResp.Body)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		logger.WithError(err).Errorf("get body error")
		return
	}
	var wxRespBody WechatLoginReponseBody
	err = json.Unmarshal(body,&wxRespBody)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		logger.WithError(err).Errorf("get body error")
		return
	}
	if wxRespBody.Errcode != 0{
		c.JSON(http.StatusOK,gin.H{
			"code":utils.Failed,
		})
		logger.WithError(err).Errorf("auth.code2Session failed")
		return
	}
	//将openID加密为Token
	hash := sha256.New()
	hash.Write([]byte(wxRespBody.Openid))
	bytes := hash.Sum(nil)
	token := string(bytes)
	logger.WithField("token",token).WithField("sessionKey",wxRespBody.SessionKey)
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"token":token,
		"sessionID":wxRespBody.SessionKey,
	})
	return
}

