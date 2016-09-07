package main

import (
	"flag"
	"log"

	"github.com/qgymje/Odyssey/commons/utils"
	"github.com/qgymje/Odyssey/servers/userCenter/rpc/models"

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
	initModel()
}

func initModel() {
	const driverName = "mysql"
	db := utils.InitDB()
	models.InitModels(db, driverName)
}

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	port := utils.GetConf().GetString("app.http_port")
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
