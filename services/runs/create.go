package runs

/*
// Run 表示Services里的Run业务集合
type Run struct {
	runModel *models.Run

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
		Comment:      form.Comment,
		RunLocations: []*models.RunLocation{},
	}
	r.rawLocations = form.RunLocations
	return r
}

// Do dodododo...
func (r *Run) Do() (err error) {
	if err = r.validLocations(); err != nil {
		return
	}

	if err = r.save(); err != nil {
		return
	}

	return
}

func (r *Run) save() (err error) {
	if err = r.runModel.Create(); err != nil {
		return
	}
	if err = models.CreateRunLocations(r.runModel.ID, r.runModel.RunLocations); err != nil {
		return
	}
	return
}

func (r *Run) validLocations() (err error) {
	if err = json.Unmarshal([]byte(r.rawLocations), &r.runModel.RunLocations); err != nil {
		return
	}
	return
}

type RunInfo struct {
	RunID     int       `json:"run_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *Run) RunInfo() *RunInfo {
	return &RunInfo{
		RunID:     r.runModel.ID,
		CreatedAt: r.runModel.CreatedAt,
	}
}

*/
