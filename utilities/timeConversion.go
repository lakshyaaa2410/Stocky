package utilities

import "time"

func GetDayAndTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return time.Now().In(loc)
}
