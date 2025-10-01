package utilities

import "time"

func GetDayAndTime() (time.Time, string) {
	location, _ := time.LoadLocation("Asia/Kolkata")
	now := time.Now().In(location)

	return now, now.Format("2006-01-02")
}
