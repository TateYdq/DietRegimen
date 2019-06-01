package main

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/handler/client/food"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/handler/client/health"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/handler/client/user"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	router:=gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"Hello World")
	})
	//数据库:
	if utils.ENV() != "test" {
		err := database.InitMysql("root:@tcp(127.0.0.1:3306)/lark_passport?charset=utf8mb4&parseTime=True&loc=Local&readTimeout=500ms")
		if err != nil {
			logrus.WithError(err).Errorf("init passport mysql failed")
			return
		}
	}
	dietRegimenPage := router.Group("/DietRegimen")
	{
		dietRegimenPage.Group("/server")
		{

		}
		clientPage := dietRegimenPage.Group("/client")
		{
			userPage := clientPage.Group("/user")
			{
				userPage.POST("/userLogin",user.UserLogin)
				userPage.GET("/getUsertInfo",user.GetUserInfo)
				userPage.POST("/updateUserInfo",user.UpdateUserInfo)
				userPage.GET("/getCollectFood",user.GetCollectFood)
				userPage.GET("/getCollectDisease",user.GetCollectDisease)
			}
			foodPage := clientPage.Group("/food")
			{
				foodPage.GET("/getFoodCategory",food.GetFoodCategory)
				foodPage.GET("/getFoodDetails",food.GetFoodDetails)
				foodPage.POST("/commentFood",food.CommentFood)
				foodPage.GET("/getComment",food.GetComment)
				foodPage.POST("/searchFood",food.SearchFood)

			}
			healthPage := clientPage.Group("/health")
			{
				healthPage.GET("/getDiseaseDetail",health.GetDiseaseDetail)
				healthPage.GET("/getComment",health.GetComment)
				healthPage.GET("/commentDisease",health.CommentDisease)
			}

		}
	}


	router.Run(":8080")

}