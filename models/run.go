package models

import "time"

// 一个跑步的纪录
type Run struct {
	Id        uint64
	UserId    uint64
	Distance  float64
	Duration  float64
	Locations []Location

	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt time.Time
}

// 仿照iOS CLLocation的结构
type Location struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
	TimeStamp time.Time
	Course    float64 //Direction
	Speed     float64
}
