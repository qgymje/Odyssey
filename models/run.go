package models

import (
	"log"
	"time"
)

// Run model 表示一个用户的一次跑步的纪录
type Run struct {
	ID        int64      `json:"run_id"`
	UserID    int64      `db:"user_id" json:"user_id"`
	Distance  float64    `json:"distance"`
	Duration  int        `json:"duration"`
	IsPublic  bool       `db:"is_public" json:"is_public"`
	Comment   NullString `json:"comment"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt NullTime   `db:"deleted_at" json:"-"`

	RunLocations []*RunLocation `db:"-" json:"locations"`
	User         *User          `db:"-" json:"-"`
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
	if err = GetDB().Select(&runs, `select * from runs where user_id = ? order by ? limit ?,?;`, userID, orderby, offset, limit); err != nil {
		log.Println(err)
		return
	}

	runIDs := []int64{}
	for _, r := range runs {
		runIDs = append(runIDs, r.ID)
	}

	runLocations, err := FindRunLocations(runIDs)
	if err != nil {
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

func FindRunByID(userID, runID int64) (*Run, error) {
	var run Run
	var err error
	if err = GetDB().Get(&run, `select runs.* from runs where id = ? and user_id = ?`, runID, userID); err != nil {
		log.Println(err)
		return nil, err
	}
	runIDs := []int64{run.ID}
	var runLocations []*RunLocation
	if runLocations, err = FindRunLocations(runIDs); err != nil {
		log.Println(err)
		return nil, err
	}
	run.RunLocations = runLocations
	return &run, nil
}
