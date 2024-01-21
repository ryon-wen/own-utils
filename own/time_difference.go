package own

import (
	"fmt"
	"log"
	"time"

	"github.com/beego/beego/toolbox"
)

func GetTimeDifference() time.Duration {
	// To set the expiration time
	today := time.Now()
	addTime := today.AddDate(0, 0, 1).Format(time.DateOnly) + " 00:00:00"
	tomorrow, _ := time.ParseInLocation(time.DateTime, addTime, today.Location())
	return tomorrow.Sub(today)
}

func HolidayTimeDifference(interval int) {
	// Get the difference between the current time and the holiday time (2024-1-27 12:00:00)
	tk := toolbox.NewTask("tk", fmt.Sprintf("0/%d * * * * *", interval), getTime)
	toolbox.AddTask(tk.Taskname, tk)
	toolbox.StartTask()
}

func getTime() error {
	now := time.Now()
	holiday := time.Date(2024, 1, 27, 12, 0, 0, 0, now.Location())
	difference := time.Unix(holiday.Unix()-now.Unix(), 0).UTC()
	day := difference.Day() - 1
	hour := difference.Hour()
	minute := difference.Minute()
	second := difference.Second()
	log.Println(fmt.Sprintf("距离春节放假还有：%d天%d时%d分%d秒！", day, hour, minute, second))
	return nil
}
