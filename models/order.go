package models

import "time"

type Order struct {
	ID        int64
	GameID    int64 `db:"game_id"`
	Payment   *Payment
	CreatedAt time.Time
}
