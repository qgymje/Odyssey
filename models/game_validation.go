package models

import "time"

type GameValidationState int

const (
	GameValidationStateInValid GameValidationState = iota
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

// GameValidation model 表示一个比赛的审核进度
type GameValidation struct {
	ID        int64
	GameID    int64
	Status    GameValidationState
	CreatedAt time.Time

	Game *Game
}
