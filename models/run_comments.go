package models

import "time"

// RunComment model 表示对一次跑步纪录的评论
type RunComment struct {
	ID              int64     `json:"run_comment_id"`
	RunID           int64     `db:"run_id" json:"run_id"`
	UserID          int64     `db:"user_id" json:"user_id"`
	ParentCommentID NullInt64 `db:"parent_comment_id" json:"parent_comment_id"` // 如果为空, 则表示对跑步纪录的评论, 不为空, 则为对用户的评论的评论

	Content   string    `json:"conente"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`

	CommentUser *User // 评论人的信息
	ReplyUser   *User // 如果是回复, 则有回复人信息, 如果是评论, 则为nil
}

// Create 创建一个评论/回复
func (rc *RunComment) Create() (err error) {
	rc.CreatedAt = time.Now()
	result, err := GetDB().NamedExec(`
insert into run_comments(run_id, user_id, parent_comment_id, content)
values(:run_id, :user_id, :parent_comment_id, :content, :created_at)
`, rc)
	if err != nil {
		return
	}
	if _, err = result.RowsAffected(); err != nil {
		return
	}
	if rc.ID, err = result.LastInsertId(); err != nil {
		return
	}
	return
}

func FindComments(runID int64, order string, limit, offset int) (comments []*RunComment, err error) {
	rows, err := GetDB().Queryx(`
select c.id as comment_id, c.run_id, c.user_id, c.parent_comment_id, c.content, c.created_at,
u.id, u.nickname, u.avatar, u2.id, u2.nickname, u2.avatar from run_comments as c
inner join users as u on c.user_id = u.id
left join run_comments as c2 on u.parent_comment_id = c2.id
left join users as u2 on c2.user_id = u2.id
where c.run_id = ? order by c.created_at desc limit ?, ?
`, runID, offset, limit)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		c := RunComment{
			CommentUser: &User{},
			ReplyUser:   &User{},
		}
		if err = rows.Scan(&c.ID,
			&c.RunID,
			&c.UserID,
			&c.ParentCommentID,
			&c.Content,
			&c.CreatedAt,
			&c.CommentUser.ID,
			&c.CommentUser.Nickname,
			&c.CommentUser.Avatar,
			&c.ReplyUser.ID,
			&c.ReplyUser.Nickname,
			&c.ReplyUser.Avatar,
		); err != nil {
			return nil, err
		}
		comments = append(comments, &c)
	}

	return comments, nil
}
