package models

import "time"

// 一个跑步的纪录
type Run struct {
	Id        uint64
	UserId    uint64
	Distance  float64
	Duration  float64
	IsPublic  bool   //是否发布?
	Comment   string // 自己的评价
	Locations []Location

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// 仿照iOS CLLocation的结构
type Location struct {
	Latitude  float64   `json:"lat"`
	Longitude float64   `json:"lng"`
	Altitude  float64   `json:"alt"`
	TimeStamp time.Time `json: "ts"`
	Course    float64   `json:"course"`
	Speed     float64   `json:"speed"`
}
