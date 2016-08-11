package main

import (
	"Odyssey/controllers"
	"Odyssey/controllers/middlewares"
	"Odyssey/models"
	"Odyssey/utils"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
)

var (
	env        = flag.String("env", "dev", "设置运行环境, 有dev, test, prod三种配置环境")
	syncdb     = flag.Bool("syncdb", false, "set syncdb")
	cpuprofile = flag.String("cpuprofile", "", "cpu profile")
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
}

func main() {
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	const driverName = "mysql"

	c := utils.GetConf().GetStringMapString("database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset-utf8&parseTime=True&loc=Local", c["username"], c["password"], c["host"], c["port"], c["dbname"])

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		panic("connect db failed.")
	}
	defer db.Close()

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)

	models.InitModels(db, driverName)

	m := utils.GetConf().GetStringMapString("mongodb")
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{m["host"]},
		Timeout:  60 * time.Second,
		Database: m["dbname"],
		Username: m["username"],
		Password: m["password"],
	}
	mongoSession, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic("connect mongodb failed")
	}
	defer mongoSession.Clone()

	mongoSession.SetMode(mgo.Monotonic, true)
	models.InitMongodb(mongoSession)

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

		follow := new(controllers.UserFollow)
		v1.POST("/user/follow", follow.Follow)
		v1.PUT("/user/unfollow", follow.Unfollow)
		v1.GET("/user/followers", follow.Followers)

		feedback := new(controllers.Feedback)
		v1.GET("/feedback", feedback.Index)
		v1.POST("/feedback", feedback.Create)
		v1.PUT("/feedback/reply", feedback.Reply)

		run := new(controllers.Run)
		v1.POST("/run", run.Create)                    // 上传一条跑步纪录
		v1.GET("/run/user/:user_id", run.Index)        // 某用户跑步纪录列表
		v1.GET("/user/run/:user_id/:run_id", run.Show) // 某一用户的某一跑步纪录

		runLike := new(controllers.RunLike)
		v1.POST("/run/like", runLike.Like)
		v1.PUT("/run/unlike", runLike.Unlike)

		runComment := new(controllers.RunComment)
		v1.GET("/run/comments/:run_id", runComment.Index)
		v1.GET("/comment/:comment_id", runComment.Show)
		v1.POST("/comment", runComment.Comment)

		game := new(controllers.Game)
		v1.GET("/game", game.Index)
		v1.GET("/game/:game_id", game.Show)
		v1.POST("/game", game.Create)
		v1.PUT("/game/edit", game.Update)
		v1.DELETE("/game", game.Destroy)

		order := new(controllers.Order)
		v1.POST("/game/order/:game_id", order.Create)
		v1.POST("/order/pay/:register_id", order.Pay)
		v1.POST("/order/cancel/:register_id", order.PayCancel)
		v1.POST("/order/refund/:register_id", order.PayRefund)

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
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
