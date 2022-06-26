package service

import (
	"time"

	"github.com/Astronaut-X-X/TaskList/back_end/model"
)

func GetTodayDailyDetailService(userid uint) ([]model.DailyDetail, bool) {
	week := time.Now().Weekday()
	weekPlan, ok := model.SelecetWeekPlanByUserIdAndWeek(userid, week)
	if !ok {
		return nil, false
	}
	if weekPlan.ID == 0 {
		return nil, false
	}
	dailyPlanID := weekPlan.DailyPlanId
	dailyPlans, ok := model.SelecetDailyDetailByDailyPlanId(dailyPlanID)
	if !ok {
		return nil, false
	}
	return dailyPlans, true
}
