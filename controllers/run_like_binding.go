package controllers

type RunLikeForm struct {
	RunID int64 `form:"run_id" binding:"required"`

	*Base
}
