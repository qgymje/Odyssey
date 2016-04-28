package runs

import (
	"Odyssey/forms"
	"Odyssey/models"
	"encoding/json"
	"time"
)

type Run struct {
	runModel *models.Run

	rawLocations string // formatted in json
}

func NewRun(form *forms.RunForm) *Run {
	r := new(Run)
	r.runModel = &models.Run{
		UserId:    form.UserId,
		Distance:  form.Distance,
		Duration:  form.Duration,
		Locations: models.Locations{},
	}
	r.rawLocations = form.Locations
	return r
}

func (r *Run) Do() error {
	if err := r.validLocations(); err != nil {
		return err
	}

	if err := r.save(); err != nil {
		return err
	}
	return nil
}

func (r *Run) save() error {
	if err := r.runModel.Create(); err != nil {
		return err
	}

	if err := r.runModel.Locations.Create(r.runModel.Id); err != nil {
		return nil
	}
	return nil
}

func (r *Run) validLocations() error {
	if err := json.Unmarshal([]byte(r.rawLocations), &r.runModel.Locations); err != nil {
		return err
	}
	return nil
}

type RunInfo struct {
	RunId     uint64    `json:"run_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *Run) RunInfo() *RunInfo {
	return &RunInfo{
		RunId:     r.runModel.Id,
		CreatedAt: r.runModel.CreatedAt,
	}
}
