package runs
/*
import (
	"Odyssey/models"
	"math"
)

func Find(userID, pageNum, pageSize int) (runs []*models.Run, err error) {
	columns := []string{
		"runs.*", "RunLocations",
	}

	relations := map[string]string{}

	where := map[string]interface{}{
		"user_id=?": userID,
	}

	offset := math.Max(float64(pageNum-1), 0.0)
	runs, err = models.FindRuns(columns, relations, where, "id DESC", pageSize, int(offset))

	return
}

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

*/
