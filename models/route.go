package models

// Route model 描述一个比赛的路线
type Route struct {
	ID uint64 `json:"route_id"`
	// 线路对应的比赛
	Game *Game `json:"game_id"`
	// 集合点
	GatheringLatitude  float64 `json:"gathering_latitude"`
	GatheringLongitude float64 `json:"gathering_longitude"`
	// 起跑点
	StartLatitude  float64 `json:"start_latitude"`
	StartLongitude float64 `json:"start_longitude"`
	// 结束点
	FinishLatitude  float64 `json:"finish_latitude"`
	FinishLongitude float64 `json:"finish_longitude"`

	// 用于保存路线地图的关键点, 用于显示线路图
	Locations []float64 // 定义为数组, 奇数位为latitude, 偶数位为logitude, 在pg里也保存也array
	// 用于显示地图上一些点, 比如厕所点, 领奖点等
	Annotations []*RouteAnnotation
}
