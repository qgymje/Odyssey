package main

import (
	"Odyssey/controllers"
	"Odyssey/middlewares"
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
	utils.InitConfig()
	utils.InitLogger()
	utils.InitRander()
	models.InitModels()
}

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		user := new(controllers.User)
		v1.POST("/smscode", middlewares.Token(), user.SMSCode)
		v1.POST("/sign_up", middlewares.Token(), user.SignUp)
		v1.POST("/sign_in", middlewares.Token(), user.SignIn)
		v1.DELETE("/sign_out", middlewares.Token(), user.SignOut)
		v1.DELETE("/delete_account", middlewares.Token(), user.DeleteAccount)

		run := new(controllers.Run)
		v1.POST("/run/create", run.Create)
		v1.GET("/run/:user_id", run.Read)
		v1.GET("/run/:user_id/:run_id", run.ReadOne)

		feedback := new(controllers.Feedback)
		v1.POST("/feedback/create", feedback.Create)
		v1.GET("/feedbacks", feedback.Read)
	}

	r.Run(":8081")
}
