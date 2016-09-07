package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/qgymje/Odyssey/commons/utils"
	"tech.cloudzen/account_center/http/controllers"
)

var (
	env        = flag.String("env", "dev", "set env: dev, test, prod")
	configPath = flag.String("conf", "./configs/", "config path")
)

func initEnv() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.Parse()
	log.Println("current env is: ", *env)
	utils.SetEnv(*env)
}
func init() {
	initEnv()
	utils.InitConfig(*configPath)
	utils.InitLogger()
	utils.InitRander()
}

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		user := new(controllers.User)
		v1.POST("/register", user.Register)
		v1.POST("/login", user.Login)
		v1.DELETE("/logout", user.Logout)
		//v1.POST("/foundpassword", user.FoundPassword) //找回密码
		//v1.POST("/resetpassword", user.ResetPassword) // 修改密码
	}

	port := utils.GetConf().GetString("app.http_port")
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
