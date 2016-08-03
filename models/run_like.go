package models

import "time"

// RunLike 表示一个赞
type RunLike struct {
	ID         int64     `json:"like_id"`
	RunID      int64     `db:"run_id" json:"run_id"`
	UserID     int64     `db:"user_id" json:"user_id"`
	IsCanceled bool      `db:"is_canceled" json:"-"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`

	User *User `json:"user"`
}

func (l *RunLike) Create() (err error) {
	l.CreatedAt = time.Now()
	result, err := GetDB().NamedExec(`insert into run_likes(run_id, user_id, is_canceled, created_at) value(:run_id, :user_id, :is_canceled, :created_at)`, l)
	if err != nil {
		return
	}
	if _, err = result.RowsAffected(); err != nil {
		return
	}
	if l.ID, err = result.LastInsertId(); err != nil {
		return
	}
	return
}

func FindRunLikes(runID int64, order string, limit, offset int) (likes []*RunLike, err error) {
	rows, err := GetDB().Queryx(`select r.*, u.id, u.nickname, u.avatar from run_likes as r inner join users as u on r.user_id = u.id where r.run_id = ? and is_canceled =  false ordery by ? limit ?,?`, runID, order, offset, limit)
	if err != nil {
		return
	}
	for rows.Next() {
		l := RunLike{
			User: &User{},
		}
		if err = rows.Scan(&l.ID,
			&l.RunID,
			&l.UserID,
			&l.IsCanceled,
			&l.CreatedAt,
			&l.User.ID,
			&l.User.Nickname,
			&l.User.Avatar,
		); err != nil {
			return
		}
		likes = append(likes, &l)
	}
	return
}
