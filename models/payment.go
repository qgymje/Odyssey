package models

import "time"

// PayState 描述用户报名比赛的支付状态
type PayState int

const (
	PayStateUnpayed PayState = iota
	PayStateFailed
	PayStatePayed
	PayStateRefunded
)

var payStateDesc = [...]string{
	"未支付", "支付失败", "已经支付", "已退款",
}

func (p PayState) String() string {
	return payStateDesc[p]
}

// Payment 纪录一次支付操作
type Payment struct {
	Id     uint64 `json:"payment_id"`
	Order  *Order
	Status PayState
	// 支付异常时候的额外信息
	Mark      string
	CreatedAt time.Time
}
