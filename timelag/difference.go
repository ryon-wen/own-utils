package timelag

import (
	"time"
)

func GetTimeDifference() time.Duration {
	// To set the expiration time
	today := time.Now()
	addTime := today.AddDate(0, 0, 1).Format(time.DateOnly) + " 00:00:00"
	tomorrow, _ := time.ParseInLocation(time.DateTime, addTime, today.Location())
	return tomorrow.Sub(today)
}
