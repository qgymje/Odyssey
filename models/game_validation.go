package models

import "time"

type GameValidationState int

const (
	GameValidationStateUnValid GameValidationState = iota
	GameValidationStateValidating
	GameValidationStateFailed
	GameValidationStateSucceeded
)

var gameValidationStateDesc = [...]string{
	"还未开始", "正在审核", "失败未通过", "成功通过",
}

func (g GameValidationState) String() string {
	return gameValidationStateDesc[g]
}

// GameValidation model 表示一个比赛的完全进度
type GameValidation struct {
	Id        uint64
	Game      *Game
	Status    GameValidationState
	CreatedAt time.Time
}
