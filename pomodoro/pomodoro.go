package pomodoro

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"
)

func countDown(totaltime int, message string) {
	for i := 0; i < totaltime; i++ {
		time.Sleep(1 * time.Minute)
		fmt.Println(message+"		time left: ", totaltime-i)
	}
}

func beep(times int) {
	os := runtime.GOOS
	var cmdStr string

	switch os {
	case "darwin":
		cmdStr = "afplay"
	case "linux":
		cmdStr = "paplay"
	case "windows":
		cmdStr = "start"
	}

	for i := 0; i < times; i++ {
		cmd := exec.Command(cmdStr, "assets/beep-06.wav")
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		err = cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Pomodoro(focus, shortBreak, longBreak, interval int) {
	breakCounter := 0

	for {
		if breakCounter < interval {
			go beep(3)
			countDown(focus, "Focus")
			go beep(3)
			countDown(shortBreak, "Break")
			breakCounter++
		} else {
			countDown(focus, "Focus")
			go beep(3)
			countDown(longBreak, "Break")
			breakCounter = 0
		}
	}

}
