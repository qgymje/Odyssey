package controllers

type RunLikeBinding struct {
	RunID int64 `form:"run_id" binding:"required"`

	*Base
}
