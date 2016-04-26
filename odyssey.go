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
	log.SetFlags(log.Ltime | log.Lshortfile)
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
		v1.GET("/sign_up", middlewares.Token(), user.SignUp)
		v1.GET("/sign_in", middlewares.Token(), user.SignIn)
		v1.GET("/sign_out", middlewares.Token(), user.SignOut)
	}

	r.Run(":8081")
}
