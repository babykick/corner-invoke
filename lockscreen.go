package main

import (
	"log"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

const (
	RIGHT_BOTTOM_X = 3039
	RIGHT_BOTTOM_Y = 899
)

func setup() {
	KillProcess("cornerinvoke")
}

func main() {
	setup()
	log.Println("press combo keys")

	hook.Register(hook.MouseDown, nil, func(e hook.Event) {
		if e.Button == 3 {
			robotgo.KeyTap("d", "cmd")
		}
	})
	s := hook.Start()
	<-hook.Process(s)

	for {
		x, y := robotgo.GetMousePos()

		if x == RIGHT_BOTTOM_X && y == RIGHT_BOTTOM_Y {
			LockScreen()
			robotgo.Move(0, 0)
		}
		time.Sleep(200 * time.Millisecond)
	}

}
