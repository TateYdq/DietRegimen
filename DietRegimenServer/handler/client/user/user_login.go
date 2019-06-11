package user

import (
	"encoding/json"
	"fmt"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/secret"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)




var(
	WechatUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code"
)
type WechatLoginRequestBody struct {
	Code string `json:"code"`
}

type WechatLoginReponseBody struct {
	Openid string `json:"openid"`
	SessionKey string `json:"session_key"`
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
	wxUrl := fmt.Sprintf(WechatUrl,utils.ConfigInstance.Wechat.AppID,utils.ConfigInstance.Wechat.AppSecret,wxReqBody.Code)
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
		logger.Errorf("auth.code2Session failed,Errcode: %v",wxRespBody.Errcode)
		return
	}
	userInfo,err := database.GetOrCreateUserInfoByOpenID(wxRespBody.Openid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Failed,
		})
		return
	}
	//距离上一次登录超过24小时奖励积分
	go database.AddUserScore(userInfo.UserID,utils.ScoreUserLogin)
	token,err := secret.CreateToken(wxRespBody.Openid,userInfo.UserID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Failed,
		})
		return
	}
	logger.Infof("openID:%v,token:%v",wxRespBody.Openid,token)
	c.JSON(http.StatusOK,gin.H{
		"code":utils.Success,
		"token":token,
	})
	return
}

