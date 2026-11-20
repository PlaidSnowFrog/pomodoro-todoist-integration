package render

import (
	"PomodoroTimer/pomodoro"
	"PomodoroTimer/todoist"
	"time"

	"github.com/gdamore/tcell/v2"
)

func drawText(s tcell.Screen, x, y int, text string, style tcell.Style) {
	for _, r := range text {
		s.SetContent(x, y, r, nil, style)
		x++
	}
}

func RenderTasks(tasks []todoist.Task) {
	// Initialize ncurses
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}

	go func() {
		for range time.Tick(time.Second) {
			s.PostEvent(tcell.NewEventInterrupt(nil))
		}
	}()

	defer s.Fini()

	checked := make([]bool, len(tasks))
	cursor := 0
	normal := tcell.StyleDefault
	highlight := tcell.StyleDefault.Reverse(true)

	for {
		s.Clear()
		if len(tasks) == 0 {
			drawText(s, 0, 0, "No tasks for today (q to quit)", normal)
		}
		for i, task := range tasks {
			box := "[ ]"
			if checked[i] {
				box = "[x]"
			}
			style := normal
			if i == cursor {
				style = highlight
			}
			drawText(s, 0, i, box+" "+task.Content, style)
		}

		_, h := s.Size()
		drawText(s, 0, h-1, pomodoro.Status, normal)

		s.Show()

		// Get Possible inputs
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize: // Resize
			s.Sync()
		case *tcell.EventKey: // Keyboard Keys
			switch {
			// Vim motions/arrow keys
			case ev.Key() == tcell.KeyCtrlC || ev.Rune() == 'q':
				return
			case ev.Key() == tcell.KeyUp || ev.Rune() == 'k':
				if cursor > 0 {
					cursor--
				}
			case ev.Key() == tcell.KeyDown || ev.Rune() == 'j':
				if cursor < len(tasks)-1 {
					cursor++
				}
			case ev.Rune() == ' ': // space
				if len(tasks) > 0 {
					checked[cursor] = !checked[cursor]
				}
			case ev.Rune() == 'p': // pause
				pomodoro.Paused.Store(!pomodoro.Paused.Load())
			}
		}
	}
}
