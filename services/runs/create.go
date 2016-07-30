package runs

import (
	"Odyssey/forms"
	"Odyssey/models"
	"Odyssey/utils"
	"database/sql"
	"encoding/json"
	"time"
)

// Run 表示Services里的Run业务集合
type Run struct {
	runModel     *models.Run
	rawLocations string // formatted in json
}

// NewRun 创建一个Run业务单元
func NewRun(form *forms.RunForm) *Run {
	r := new(Run)
	r.runModel = &models.Run{
		UserID:       form.UserID,
		Distance:     form.Distance,
		Duration:     form.Duration,
		IsPublic:     form.IsPublic,
		Comment:      sql.NullString{String: form.Comment},
		RunLocations: []*models.RunLocation{},
	}
	r.rawLocations = form.RunLocations
	return r
}

// Do dodododo...
func (r *Run) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.Run.Do error", err)
		}
	}()

	if err = r.validLocations(); err != nil {
		return
	}

	if err = r.save(); err != nil {
		return
	}

	return
}

func (r *Run) save() (err error) {
	err = r.runModel.Create()
	return
}

func (r *Run) validLocations() (err error) {
	err = json.Unmarshal([]byte(r.rawLocations), &r.runModel.RunLocations)
	return
}

type RunInfo struct {
	RunID     int64     `json:"run_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *Run) RunInfo() *RunInfo {
	return &RunInfo{
		RunID:     r.runModel.ID,
		UserID:    r.runModel.UserID,
		CreatedAt: r.runModel.CreatedAt,
	}
}
