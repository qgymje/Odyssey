package main

import (
	"Odyssey/models"
	"Odyssey/utils"
	"flag"
	"log"

	"github.com/streadway/amqp"
)

var (
	env = flag.String("env", "dev", "设置运行环境, 有dev, test, prod三种配置环境")
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

var connection *amqp.Connection
var channel *amqp.Channel

func main() {
	connection, channel = utils.GetAMQP()

	var (
		exchnage   = "smssender"
		queue      = exchange
		bindingKey = queue
	)
	//func (me *Channel) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args Table) (Queue, error)
	q, err := channel.QueueDeclare(
		queue, // queue name
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	//func (me *Channel) QueueBind(name, key, exchange string, noWait bool, args Table) error
	if err = channel.QueueBind(
		q.Name,
		bindingKey, // here is the binding key
		exchnage,   // exchange name
		false,
		nil,
	); err != nil {
		log.Fatal(err)
	}

	go func() {
		for d := range msgs {
			handleMessage(d.Body)
			d.Ack(false)
		}
	}()

	done := make(chan bool)
	log.Printf("[*] Waiting for messages. To exit press CTRL+C ")
	<-done
}

func handleMessage(msg []byte) {
	log.Printf("[x]received msg: %s", msg)
}
