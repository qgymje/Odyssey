package main

import (
	"Odyssey/controllers"
	"Odyssey/models"
	"Odyssey/utils"
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	env    = flag.String("env", "dev", "设置运行环境, 有dev, test, prod三种配置环境")
	syncdb = flag.Bool("syncdb", false, "set syncdb")
)

func initEnv() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.Parse()
	log.Println("当前运行环境为: ", *env)
	utils.SetEnv(*env)
}

func init() {
	initEnv()
	utils.InitConfig("./configs/")
	utils.InitLogger()
	utils.InitRander()
	models.InitModels()
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		user := new(controllers.User)
		v1.POST("/smscode", user.SMSCode)
		v1.POST("/sign_up", user.SignUp)
		v1.POST("/sign_in", user.SignIn)
		v1.DELETE("/sign_out", user.SignOut)

		run := new(controllers.Run)
		v1.POST("/run/create", run.Create)
		v1.GET("/run/:user_id", run.Read)
		v1.GET("/run/:user_id/:run_id", run.ReadOne)

		/*
			feedback := new(controllers.Feedback)
			v1.POST("/feedback/create", feedback.Create)
			v1.GET("/feedbacks", feedback.Read)
		*/

	}

	port := utils.GetConf().GetString("app.http_port")
	r.Run(":" + port)
}
