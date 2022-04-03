package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {

	s := gocron.NewScheduler(time.Local)

	s.Every(5).Seconds().Do(func() {
		fmt.Println("Every 5 seconds: ", time.Now())
	})

	s.Cron("*/1 * * * *").Do(everyMinuteTask)

	s.Every(1).Day().At("14:59").Do(everyDayTask)

	s.Every(1).Day().Hour().Do(everyHourTask)

	// Multiple time
	s.Every(1).Day().At("14:59").At("15:00").Do(func() {
		everyDayTask()
	})

	s.StartAsync()
	s.StartBlocking()
}

func everyMinuteTask() {
	fmt.Println("Every minute: ", time.Now())
}

func everyDayTask() {
	fmt.Println("Every day at : ", time.Now())
}

func everyHourTask() {
	fmt.Println("Every hour: ", time.Now())
}
