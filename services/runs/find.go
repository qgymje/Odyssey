package runs

import "Odyssey/models"

func FindOne(userID, runID int) (run *models.Run, err error) {
	/*
		where := map[string]interface{}{
			"id=?":      runID,
			"user_id=?": userID,
		}
		where2 := map[string]interface{}{
			"run_id=?": runID,
		}
	*/
	return
}

func Find(userId, runId uint64) (runs []models.Run, err error) {
	/*
		where := map[string]interface{}{
			"user_id=?": userId,
		}
	*/

	return
}
