package models

/*
import (
	"Odyssey/utils"
	"time"

	"gopkg.in/pg.v4/orm"
)

// Run model 表示一个用户的一次跑步的纪录
type Run struct {
	TableName struct{} `sql:"runs,alias:runs" json:"-"`
	ID        int      `json:"run_id"`
	UserID    int      `json:"user_id"`
	Distance  float64  `json:"distance"`
	Duration  int      `json:"duration"`
	//Setps     int       `json:"steps"` // 步数
	IsPublic     bool           `json:"is_public"`
	Comment      string         `sql:",null" json:"comment"`
	RunLocations []*RunLocation `json:"run_locaitons"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `sql:",null" json:"-"`
	DeletedAt time.Time `sql:",null" json:"-"`
}

// Create 创建一条跑步纪录, 需要RunLocation数据
func (r *Run) Create() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.run.Create error: ", err)
		}
	}()

	r.CreatedAt = time.Now()
	err = GetDB().Create(r)

	return
}

// FindRuns 查找跑步纪录
func FindRuns(columns []string, relations map[string]string, where map[string]interface{}, order string, limit int, offset int) (runs []*Run, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.run.FindRuns error: ", err)
		}
	}()

	for key, val := range relations {
		query = query.Relation(key, func(*orm.Query) *orm.Query {
			return query.Where(val)
		})
	}
	for key, val := range where {
		query = query.Where(key, val)
	}
	err = query.Order(order).Limit(limit).Offset(offset).Find(&runs)

	return
}
*/
