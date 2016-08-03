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
		v1.POST("/foundpassword", user.FoundPassword) //找回密码
		v1.POST("/resetpassword", user.ResetPassword) // 修改密码

		feedback := new(controllers.Feedback)
		v1.GET("/feedback", feedback.Index)
		v1.POST("/feedback", feedback.Create)
		v1.POST("/feedback/reply/:feedback_id", feedback.Reply)

		run := new(controllers.Run)
		v1.POST("/run", run.Create)                    // 上传一条跑步纪录
		v1.GET("/run/user/:user_id", run.Index)        // 某用户跑步纪录列表
		v1.GET("/user/run/:user_id/:run_id", run.Show) // 某一用户的某一跑步纪录

		runLike := new(controllers.RunLike)
		v1.POST("/run/like/:run_id", runLike.Like)
		v1.POST("/run/unlike/:run_id", runLike.Unlike)

		runComment := new(controllers.RunComment)
		v1.GET("/run/comments/:run_id", runComment.Index)
		v1.GET("/comment/:comment_id", runComment.Show)
		v1.POST("/comment", runComment.Comment)
		v1.POST("/comment/reply/:comment_id", runComment.Reply)

		game := new(controllers.Game)
		v1.GET("/game", game.Index)
		v1.GET("/game/:game_id", game.Show)
		v1.POST("/game", game.Create)
		v1.PUT("/game/edit", game.Update)
		v1.DELETE("/game", game.Destroy)

		register := new(controllers.Register)
		v1.POST("/game/register/:game_id", register.Create)
		v1.POST("/register/pay/:register_id", register.Pay)
		v1.POST("/register/cancel/:register_id", register.PayCancel)
		v1.POST("/register/refund/:register_id", register.PayRefund)

		v1.GET("/user/profile/:user_id", user.Profile)
		v1.GET("/user/around", user.Around) // 用户周围的人
		v1.GET("/user/games", user.Games)
		v1.GET("/user/friends", user.Friends)
		v1.GET("/user/groups", user.Groups)

		group := new(controllers.Group)
		v1.GET("/group", group.Index)
		v1.GET("/group/:group_id", group.Show)
		v1.POST("/group", group.Create)
		v1.PUT("/group/:group_id", group.Update)
		v1.POST("/group/join/:group_id", group.Join)
		v1.DELETE("/group/quit/:group_id", group.Quit)

	}

	port := utils.GetConf().GetString("app.http_port")
	r.Run(":" + port)
}
