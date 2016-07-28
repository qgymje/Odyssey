package models

import (
	"Odyssey/utils"
	"time"
)

// Run model 表示一个用户的一次跑步的纪录
type Run struct {
	TableName struct{} `sql:"runs"`
	ID        int      `json:"run_id"`
	User      *User    `json:"user_id"`
	Distance  float64  `json:"distance"`
	Duration  int      `json:"duration"`
	//Setps     int       `json:"steps"` // 步数
	IsPublic     bool           `json:"is_public"`
	Comment      string         `json:"comment"`
	RunLocations []*RunLocation `json:"run_locaitons"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// Create 创建一条跑步纪录, 需要RunLocation数据
func (r *Run) Create() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.run.Create error: %v", err)
		}
	}()

	r.CreatedAt = time.Now()
	err = GetDB().Create(r)

	return
}

// FindRuns 查找跑步纪录
func FindRuns(where map[string]interface{}, order string, limit int, offset int) (runs []*Run, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.run.FindRun error: %v", err)
		}
	}()

	query := GetDB().Model(&runs)
	for key, val := range where {
		query = query.Where(key, val)
	}
	err = query.Order(order).Limit(limit).Offset(offset).Select()

	return
}
