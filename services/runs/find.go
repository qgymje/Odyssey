package runs

import (
	"Odyssey/models"
	"Odyssey/utils"
)

func FindOne(userId, runId uint64) (*models.Run, error) {
	var err error
	where := map[string]interface{}{
		"id":      runId,
		"user_id": userId,
	}
	where2 := map[string]interface{}{
		"run_id": runId,
	}
	rs := []*models.Run{}
	if rs, err = models.FindRuns(where); err != nil {
		return nil, err
	}
	utils.Dump(rs)
	ls := models.Locations{}
	if ls, err = models.FindLocations(where2); err != nil {
		return nil, err
	}
	rs[0].Locations = ls
	return rs[0], nil
}

func Find(userId, runId uint64) ([]*models.Run, error) {
	where := map[string]interface{}{
		"user_id": userId,
	}

	return models.FindRuns(where)
}
