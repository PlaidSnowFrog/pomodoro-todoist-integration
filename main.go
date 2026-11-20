package main

import (
	"PomodoroTimer/pomodoro"
	"PomodoroTimer/todoist"
	"log"
)

func main() {
	tasks := todoist.GetTasks()

	// Logging
	log.Println(tasks)
	for _, task := range tasks {
		log.Println(task.Content)
	}

	go pomodoro.Pomodoro(20, 3, 10, 3)

}
