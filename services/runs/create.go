package runs

import "Odyssey/models"

type Run struct {
	runModel *models.Run

	rawLocations string // formatted in json
}

/*
func NewRun(form *forms.RunForm) *Run {
	r := new(Run)
	r.runModel = &models.Run{
		UserId:       form.UserId,
		Distance:     form.Distance,
		Duration:     form.Duration,
		IsPublic:     form.IsPublic,
		Comment:      form.Comment,
		RunLocations: models.RunLocations{},
	}
	r.rawLocations = form.RunLocations
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

	utils.Dump(r)
	if err := r.runModel.RunLocations.Create(r.runModel.Id); err != nil {
		return err
	}
	return nil
}

func (r *Run) validLocations() error {
	if err := json.Unmarshal([]byte(r.rawLocations), &r.runModel.RunLocations); err != nil {
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

*/
