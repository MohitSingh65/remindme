package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"
)

func sendNotification(message string) {
	cmd := exec.Command("notify-send", "remindme", message)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error sending notification.", err)
	}
}

func setReminder(message string, duration time.Duration) {
	cmd := exec.Command("sh", "-c", fmt.Sprintf(
		"nohup sleep %d && notify-send remindme '%s' &",
		int(duration.Seconds()),
		message,
	))
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting reminder process:", err)
	} else {
		fmt.Println("Reminder process started successfully!")
	}
}

func parseTime(input string) (time.Duration, error) {
	layout := "15:04"
	targetTime, err := time.Parse(layout, input)
	if err != nil {
		return 0, fmt.Errorf("invalid time format. HH:MM (24-hour format)")
	}
	now := time.Now()

	targetTime = time.Date(now.Year(), now.Month(), now.Day(), targetTime.Hour(), targetTime.Minute(), 0, 0, now.Location())

	if targetTime.Before(now) {
		targetTime = targetTime.Add(24 * time.Hour)
	}

	duration := targetTime.Sub(now)
	return duration, nil
}

func main() {
	message := flag.String("message", "", "The reminder message")
	minutes := flag.Int("minutes", 0, "Time in minutes for the reminder")
	hours := flag.Int("hours", 0, "Time in hours for the reminder")
	exactTime := flag.String("at", "", "Exact time for the reminder (HH:MM)")
	flag.Parse()

	var duration time.Duration
	var err error

	if *exactTime != "" {
		duration, err = parseTime(*exactTime)
		if err != nil {
			fmt.Println(err)
			return
		}

	} else if *minutes > 0 || *hours > 0 {
		duration = time.Duration(*hours)*time.Hour + time.Duration(*minutes)*time.Minute
	} else {
		fmt.Println("Error: You must specify a time for the reminder.")
		flag.Usage()
		return
	}
	fmt.Printf("Reminder set for %v from now.\n", duration)
	setReminder(*message, duration)
}
