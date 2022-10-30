package controller

type GetListArgs struct {
	Page           int    `form:"page"`
	Count          int    `form:"count"`
	Title          string `form:"title"`
	Content        string `form:"content"`
	StartTimeStart int64  `form:"start_time_start"`
	StartTimeEnd   int64  `form:"start_time_end"`
	EndTimeStart   int64  `form:"end_time_start"`
	EndTimeEnd     int64  `form:"end_time_end"`
}
