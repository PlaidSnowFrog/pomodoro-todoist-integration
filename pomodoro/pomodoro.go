package pomodoro

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"sync/atomic"
	"time"
)

var Paused atomic.Bool
var Status string

func countDown(totaltime int, message string) {
	for left := totaltime * 60; left > 0; {
		Status = fmt.Sprintf("%s     time left: %02d:%02d (press p to pause)", message, left/60, left%60)
		if Paused.Load() {
			Status = fmt.Sprintf("%s     time left: %02d:%02d (press p to resume)", message, left/60, left%60)
		} else {
			left--
		}
		time.Sleep(1 * time.Second)
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
