package games

import (
	"Odyssey/models"
	"time"
)

type Game struct {
	gameModel *models.Game
}

type GameConfig struct {
	Name               string
	Slogan             string
	MaximumParticipant int
	Cost               float32
	RegisterTime       time.Time
	StartTime          time.Time
	Duration           int
	Distance           float32
}
