package runs

import "Odyssey/models"

func Find(userId, runId uint64) ([]*models.Run, error) {
	where := map[string]interface{}{
		"id":      runId,
		"user_id": userId,
	}

	return models.FindRuns(where)
}
