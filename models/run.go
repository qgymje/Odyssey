package models

import (
	"database/sql"
	"time"
)

// Run model 表示一个用户的一次跑步的纪录
type Run struct {
	ID        int64
	UserID    int64 `db:"user_id"`
	Distance  float64
	Duration  int
	IsPublic  bool `db:"is_public"`
	Comment   sql.NullString
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt NullTime  `db:"deleted_at"`

	RunLocations []*RunLocation `db:"-"`
}

// Create 创建一条跑步纪录, 需要RunLocation数据
func (r *Run) Create() (err error) {
	now := time.Now()
	r.CreatedAt = now
	r.UpdatedAt = now

	result, err := GetDB().NamedExec(`insert into runs(user_id, distance, duration, is_public, comment, created_at, updated_at) values(:user_id, :distance, :duration, :is_public, :comment, :created_at, :updated_at)`, r)
	if err != nil {
		return
	}
	if _, err = result.RowsAffected(); err != nil {
		return
	}
	if r.ID, err = result.LastInsertId(); err != nil {
		return
	}
	if err = CreateRunLocations(r.ID, r.RunLocations); err != nil {
		return
	}

	return
}

func FindRunsByUserID(userID int64, orderby string, limit, offset int) (runs []*Run, err error) {
	if err = GetDB().Select(runs, `select runs.* where user_id = ? order by ? limti = ? offset = ?;`, userID, orderby, limit, offset); err != nil {
		return
	}

	runIDs := []int64{}
	for _, r := range runs {
		runIDs = append(runIDs, r.ID)
	}

	var runLocations []*RunLocation
	if err = GetDB().Select(runLocations, `select run_locations.* from run_locations where run_id in (?);`, runIDs); err != nil {
		return
	}

	runIDLocations := map[int64][]*RunLocation{}
	for _, rl := range runLocations {
		runIDLocations[rl.RunID] = append(runIDLocations[rl.RunID], rl)
	}

	for i, r := range runs {
		runs[i].RunLocations = runIDLocations[r.ID]
	}

	return
}

func FindRunByID(runID int64) (run *Run, err error) {
	return
}
