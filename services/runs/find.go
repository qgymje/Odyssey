package runs

import (
	"Odyssey/models"
	"math"
)

func Find(userID int64, pageNum, pageSize int) (runs []*models.Run, err error) {
	offset := math.Max(float64(pageNum-1), 0.0)

	runs, err = models.FindRunsByUserID(userID, "id DESC", pageSize, int(offset))

	return
}

func FindOne(userID, runID int64) (run *models.Run, err error) {
	return models.FindRunByID(userID, runID)
}
