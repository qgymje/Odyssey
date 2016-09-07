package smsClient

import (
	"context"

	pb "github.com/qgymje/Odyssey/commons/protobufs/sms"
	"google.golang.org/grpc"
	"tech.cloudzen/utils"
)

const (
	address = "localhost:3001"
)

type SMSClient struct {
	conn   *grpc.ClientConn
	client pb.SMSClient
}

func NewSMSClient() *SMSClient {
	// 在请求连接之前, 先向服务器请求以获得服务器地址
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		// 是否增加一个用于通知的功能? 统一在error处理
		utils.GetLog().Error("sms grpc server can not connect :%v", err)
	}
	sms := new(SMSClient)
	sms.conn = conn
	sms.client = pb.NewSMSClient(sms.conn)

	return sms
}

func (sms *SMSClient) Close() error {
	return sms.Close()
}

func (sms *SMSClient) Send(in *pb.SendInfo) (*pb.Status, error) {
	defer sms.Close()
	return sms.client.Send(context.Background(), in)
}
