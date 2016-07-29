package runs

import (
	"Odyssey/models"
	"encoding/json"
	"testing"
)

func TestLocaitons(t *testing.T) {
	locs := `[
    {"lat":31.003800, "lng":121.223080, "alt": 5.2, "ts": "2016-07-28T17:49:09.088899+08:00"},
    {"lat":31.203800, "lng":121.233080, "alt": 5.2, "ts": "2016-07-28T17:52:09.088899+08:00"},
    {"lat":31.403800, "lng":121.243080, "alt": 5.2, "ts": "2016-07-28T17:55:09.088899+08:00"}
   ]`

	var locations []models.RunLocation
	err := json.Unmarshal([]byte(locs), &locations)
	if err != nil {
		t.Error("locs format error:", err)
	}

}
