package util

import "time"

func getWeek() time.Weekday {
	return time.Now().Weekday()
}
