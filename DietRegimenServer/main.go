package main

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/handler/client/food"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/handler/client/health"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/handler/client/user"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/handler/server"
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
	config,err := utils.InitConfig()
	if err != nil {
		logrus.WithError(err).Errorf("init config failed")
		return
	}
	//数据库:
	err = database.InitMysql(config.Mysql.Database)
	if err != nil {
		logrus.WithError(err).Errorf("init passport mysql failed")
		return
	}
	dietRegimenPage := router.Group("/DietRegimen")
	{
		serverPage := dietRegimenPage.Group("/control1/admin")
		{
			serverPage.POST("/addUser", server.AddUser)
			serverPage.POST("/updateUser", server.UpdateUser)

			serverPage.POST("/addFood", server.AddFood)
			serverPage.POST("/updateFood", server.UpdateFood)

			serverPage.POST("/addFoodKind", server.AddFoodKind)
			serverPage.POST("/updateFoodKind", server.UpdateFoodKind)

			serverPage.POST("/addDisease", server.AddDisease)
			serverPage.POST("/updateDisease", server.UpdateDisease)

		}
		clientPage := dietRegimenPage.Group("/client")
		{
			userPage := clientPage.Group("/user")
			{
				userPage.POST("/userLogin",user.UserLogin)

				userPage.GET("/getUserInfo",user.GetUserInfo)
				userPage.POST("/updateUserInfo",user.UpdateUserInfo)

				userPage.GET("/collectFood",user.CollectFood)
				userPage.GET("/getCollectFood",user.GetCollectFood)

				userPage.GET("/collectDisease",user.CollectDisease)
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