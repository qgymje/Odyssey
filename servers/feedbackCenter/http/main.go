package main

import (
	"flag"
	"log"

	"github.com/qgymje/Odyssey/feedbackCenter/http/controllers"
	"github.com/qgymje/Odyssey/utils"

	"github.com/gin-gonic/gin"
)

var (
	env        = flag.String("env", "dev", "set env: dev, test, prod")
	syncdb     = flag.Bool("syncdb", false, "set syncdb")
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
		feedback := new(controllers.Feedback)
		v1.GET("/feedback", feedback.Index)
		v1.POST("/feedback", feedback.Create)
		v1.PUT("/feedback/reply", feedback.Reply)
		//v1.GET("/faq")
	}

	port := utils.GetConf().GetString("app.http_port")
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
