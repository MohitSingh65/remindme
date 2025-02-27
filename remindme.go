package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"
)

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

func main() {
	message := flag.String("message", "", "The reminder message")
	minutes := flag.Int("minutes", 0, "Time in minutes for the reminder")
	hours := flag.Int("hours", 0, "Time in hours for the reminder")
	flag.Parse()

	if *message == "" {
		fmt.Println("Reminder message is required.")
		flag.Usage()
		return
	}

	if *minutes == 0 && *hours == 0 {
		fmt.Println("Error: You must specify a time for the reminder.")
		flag.Usage()
		return
	}

	duration := time.Duration(*hours)*time.Hour + time.Duration(*minutes)*time.Minute

	fmt.Printf("Reminder set for %v from now.\n", duration)

	setReminder(*message, duration)
}
