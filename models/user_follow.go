package models

import "time"

// Follow 表示关注纪录表
type Follow struct {
	FromUserID int64     `db:"from_user_id" json:"from_user_id"`
	ToUserID   int64     `db:"to_user_id" json:"to_user_id"`
	IsCanceled bool      `db:"is_canceled" json:"-"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

func (f *Follow) Follow() (err error) {
	f.IsCanceled = false
	return f.Create()
}

func (f *Follow) UnFollow() (err error) {
	f.IsCanceled = true
	return f.Create()
}

func (f *Follow) Create() (err error) {
	f.CreatedAt = time.Now()
	result, err := GetDB().NamedExec(`
replace into user_follows(from_user_id, to_user_id, is_canceled, created_at) values(:from_user_id, :to_user_id, :is_canceled, :created_at)
`, f)
	if err != nil {
		return
	}
	if _, err = result.RowsAffected(); err != nil {
		return
	}
	return
}

func FindFollowers(userID int64, order string, limit, offset int) (users []*User, err error) {
	rows, err := GetDB().Queryx(`
select u.* from user_follows as f
inner join users as u on f.from_user_id = f.id
where f.to_user_id = ? ordery by ? limit ?, ?
`, userID, order, offset, limit)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		u := User{}
		if err = rows.StructScan(&u); err != nil {
			return
		}
		users = append(users, &u)
	}
	return
}

// FindCommonFollowers 查找共同关注的人
func FindCommonFollowers(userA, userB int64) (users []*User, err error) {
	rows, err := GetDB().Queryx(`
select u.*  from
(select to_user_id from user_follows where from_user_id = ?) as a
inner join
(select to_user_id from user_follows where from_user_id = ?) as b
on a.to_user_id = b.to_user_id
inner join users as u on a.to_user_id = u.id
`, userA, userB)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		u := User{}
		if err = rows.StructScan(&u); err != nil {
			return
		}
		users = append(users, &u)
	}
	return
}
