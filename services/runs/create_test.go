package runs

import (
	"Odyssey/models"
	"encoding/json"
	"testing"
)

func TestLocaitons(t *testing.T) {
	locs := `[
    {"latitude":31.003800, "longitude":121.223080, "altitude": 5.2, "timestamp": "2016-07-28T17:49:09.088899+08:00"},
    {"latitude":31.203800, "longitude":121.233080, "alttitude": 5.2, "timestams": "2016-07-28T17:52:09.088899+08:00"},
    {"latitude":31.403800, "longitude":121.243080, "altitude": 5.2, "timestamp": "2016-07-28T17:55:09.088899+08:00"}
   ]`

	var locations []models.RunLocation
	err := json.Unmarshal([]byte(locs), &locations)
	if err != nil {
		t.Error("locs format error:", err)
	}

}
