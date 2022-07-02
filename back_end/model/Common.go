package model

import "time"

type PageInfo struct {
	Page      int       `json:"page" form:"page"`
	PageSize  int       `json:"pageSize" form:"pageSize"`
	Keyword   string    `json:"keyword" form:"keyword"`
	StartTime time.Time `json:"start_time" form:"start_time"`
	EndTime   time.Time `json:"end_time" form:"end_time"`
}

type PageData struct {
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
	Total    int           `json:"total"`
	Data     []interface{} `json:"data"`
}
