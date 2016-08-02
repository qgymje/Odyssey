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
	utils.InitConfig("./configs/")
	utils.InitLogger()
	utils.InitRander()
	models.InitModels()
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.FakedLogin())

	v1 := r.Group("/api/v1")
	{

		user := new(controllers.User)
		v1.POST("/smscode", user.SMSCode)
		v1.POST("/register", user.Register)
		v1.POST("/login", user.Login)
		v1.DELETE("/logout", user.Logout)

		run := new(controllers.Run)
		v1.POST("/run/:user_id", run.Create)
		v1.GET("/run/:user_id", run.Index)
		v1.GET("/run/:user_id/:run_id", run.Show)

		feedback := new(controllers.Feedback)
		v1.GET("/feedback", feedback.Index)
		v1.POST("/feedback", feedback.Create)
		v1.POST("/feedback/:feedback_id/reply", feedback.Reply)

		/*
			v1.Get("/user", user.Profile)
			v1.Get("/user/games", user.Games)
			v1.Get("/user/friends", user.Friends)
			v1.Get("/user/groups", user.Groups)

				group := new(controllers.Group)
				v1.Get("/group", group.Index)
				v1.Get("/group/:group_id", group.Show)
				v1.Post("/group", group.Create)
				v1.Post("/group/:group_id/join", group.Join)
				v1.Delete("/group/:group_id/quit", group.Quit)

				runLike := new(controller.RunLike)
				v1.Post("/run/:run_id/like)
				v1.Post("/run/:run_id/unlike)

				runComment := new(controller.RunComment)
				v1.Get("/run/comment", runComment.Show)
				v1.Post("/run/comment/:run_id", runComment.Comment)
				v1.Post("/run/comment/:run_id/reply/:comment_id", runComment.Reply)

				game := new(controllers.Game)
				v1.Get("/game", game.Index)
				v1.Get("/game/:game_id", game.Show)
				v1.Post("/game", game.Create)
				v1.Put("/game/edit", game.Update)
				v1.Delete("/game", game.Destroy)

				register := new(controllers.Register)
				v1.Post("/register/:game_id", register.Create)
				v1.Post("/register/:game_id/pay", register.Pay)
				v1.Post("/register/:game_id/cancel", register.PayCancel)
				v1.Post("/register/:game_id/refund", register.PayRefund)

		*/
	}

	port := utils.GetConf().GetString("app.http_port")
	r.Run(":" + port)
}
