package runs

import (
	"Odyssey/models"
	"Odyssey/utils"
	"math"

	"github.com/pkg/errors"
)

func Find(userID int64, pageNum, pageSize int) (runs []*models.Run, err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "services.runs.Find error")
			utils.GetLog().Error("%+v", err)
		}
	}()
	offset := math.Max(float64(pageNum-1), 0.0)

	runs, err = models.FindRunsByUserID(userID, "id DESC", pageSize, int(offset))

	return
}

func FindOne(userID, runID int64) (run *models.Run, err error) {
	return models.FindRunByID(userID, runID)
}
