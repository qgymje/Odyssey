package models

// Oraginazer model 表示一个组织者, 代表个人或者一个团队
type Oraginazer struct {
	Id   uint64 `json:"oraginzer_id"`
	Name string `json:"name"`
}
