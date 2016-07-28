package async

type ExchangeType int

const (
	Direct ExchangeType = iota + 1
	Fanout
	Topic
)

// Queue 表示一个rabbitmq封装的实例
type Queue struct {
	exchange     string
	exchangeType ExchangeType
	routingKey   string
}
