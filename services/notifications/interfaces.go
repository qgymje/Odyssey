package notifications

import "Odyssey/models"

// Notice 接口定义添加notice事件接口, 面向用户类
type Notice interface {
	Type() models.EventType     // 事件名字, 比如回复用户反馈
	Messages() map[int64]string // user_id: message
	// 通过手机Notification or Email or SMS?
}
