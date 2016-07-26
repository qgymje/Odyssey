package models

// RouteAnnotation model 表示比赛线路上的标注的点, 比如早餐点, 休息站, 等等,
type RouteAnnotation struct {
	Id uint64 `json:"route_annotation_id"`
	// 对应的线路图
	Route *Route `json:"route_id"`
	// 此annotation名字
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Image    string `json:"image"`
	// 支持多个地点
	Location []float64
	Mark     string `json:"mark"`
}
