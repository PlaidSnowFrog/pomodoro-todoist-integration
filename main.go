package main

import (
	"PomodoroTimer/pomodoro"
	"PomodoroTimer/render"
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
	if len(tasks) == 0 {
		log.Fatal("No tasks found")
	}

	go pomodoro.Pomodoro(20, 3, 10, 3)

	render.RenderTasks(tasks)
}
