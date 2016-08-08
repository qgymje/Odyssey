package runs

import (
	"Odyssey/models"
	"Odyssey/utils"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Run 表示Services里的Run业务集合
type Run struct {
	runModel     *models.Run
	rawLocations string // configatted in json
}

type RunConfig struct {
	UserID       int64
	Distance     float64
	Duration     int
	IsPublic     bool
	Comment      string
	RunLocations string
}

// NewRun 创建一个Run业务单元
func NewRun(config *RunConfig) *Run {
	r := new(Run)
	r.runModel = &models.Run{
		UserID:       config.UserID,
		Distance:     config.Distance,
		Duration:     config.Duration,
		IsPublic:     config.IsPublic,
		Comment:      models.ToNullString(config.Comment),
		RunLocations: []*models.RunLocation{},
	}
	r.rawLocations = config.RunLocations
	return r
}

// Do dodododo...
func (r *Run) Do() (err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "services.Run.Do error")
			utils.GetLog().Error("%+v", err)
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
