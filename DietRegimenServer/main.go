package main

import (
	"github.com/TateYdq/DietRegimen/DietRegimenServer/cache"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/database"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/handler/client/food"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/handler/client/health"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/handler/client/recommend"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/handler/client/user"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/handler/file_interact"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/handler/server"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/midware"
	"github.com/TateYdq/DietRegimen/DietRegimenServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	config, err := utils.InitConfig()
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
	err = cache.Init(config.Redis.Addr, config.Redis.Password, config.Redis.DbToken)
	if err != nil {
		logrus.WithError(err).Errorf("init cache failed")
		return
	}
	dietRegimenPage := router.Group("/")
	{
		serverPage := dietRegimenPage.Group("/control1/admin")
		serverPage.Use(midware.AdminVerify)
		{
			serverPage.POST("/addUser", server.AddUser)
			serverPage.POST("/updateUser", server.UpdateUser)

			serverPage.POST("/addFood", server.AddFood)
			serverPage.POST("/updateFood", server.UpdateFood)

			serverPage.POST("/addFoodKind", server.AddFoodKind)
			serverPage.POST("/updateFoodKind", server.UpdateFoodKind)

			serverPage.POST("/addDisease", server.AddDisease)
			serverPage.POST("/updateDisease", server.UpdateDisease)
			serverPage.POST("/addDiseaseFoodRec", server.AddDiseaseFoodRec)

			serverPage.POST("/uploadImage", file_interact.Fileupload)

			serverPage.POST("/addQuestion", server.AddQuestion)

		}
		clientPage := dietRegimenPage.Group("/client")
		{
			userPage := clientPage.Group("/user")
			{
				userPage.POST("/userLogin", user.UserLogin)

				userPage.GET("/getUserInfo", user.GetUserInfo)
				userPage.POST("/updateUserInfo", user.UpdateUserInfo)

				userPage.POST("/collectFood", user.CollectFood)
				userPage.GET("/getCollectFood", user.GetCollectFood)

				userPage.POST("/collectDisease", user.CollectDisease)
				userPage.GET("/getCollectDisease", user.GetCollectDisease)

				userPage.POST("/uploadUserImage", user.UploadUserImage)
			}
			foodPage := clientPage.Group("/food")
			{
				foodPage.GET("/getFoodCategory", food.GetFoodCategory)

				foodPage.GET("/getFoodDetails", food.GetFoodDetails)

				foodPage.POST("/commentFood", food.CommentFood)
				foodPage.GET("/getComment", food.GetComment)

				foodPage.GET("/searchFood", food.SearchFood)
			}
			healthPage := clientPage.Group("/health")
			{
				healthPage.GET("/getDiseaseDetails", health.GetDiseaseDetails)

				healthPage.GET("/getDiseasesLists", health.GetDiseasesLists)

				healthPage.GET("/getComment", health.GetComment)
				healthPage.POST("/commentDisease", health.CommentDisease)
			}
			recommendPage := clientPage.Group("/recommend")
			{
				recommendPage.GET("/getQuestionnaire", recommend.GetQuestionnaire)
				recommendPage.POST("/submitQuestionnaire", recommend.SubmitQuestionnaire)

				recommendPage.POST("/getRecInfo", recommend.GetRecInfo)
			}
		}
		filePage := dietRegimenPage.Group("/file")
		{
			filePage.GET("/getImage", file_interact.GetImage)
			filePage.GET("/getVoice", file_interact.GetVoice)
		}

	}
	router.Run(":8080")

}
