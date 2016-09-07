package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/qgymje/Odyssey/commons/protobufs/sms"
	"github.com/qgymje/Odyssey/commons/utils"
	"github.com/qgymje/Odyssey/servers/smsCenter/rpc/models"
	"google.golang.org/grpc"
)

const (
	port = "localhost:9595"
)

var (
	configPath = flag.String("conf", "./configs/", "set config path")
	env        = flag.String("env", "dev", "set env: dev, test, prod")
)

func initEnv() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	flag.Parse()
	log.Println("current env is: ", *env)
	utils.SetEnv(*env)
}

func init() {
	initEnv()
	utils.InitConfig(*configPath)
	utils.InitLogger()
	utils.InitRander()
	db := utils.InitMysql()
	models.InitModels(db, "mysql")
}

func main() {
	prot := utils.GetConf().String("app.rpc_port")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSMSServer(s, &sms.SMSServer{})
	s.Serve(lis)
}
