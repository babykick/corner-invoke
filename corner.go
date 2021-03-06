package main

import (
	"fmt"
	"log"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var (
	RIGHT_BOTTOM_X int16
	RIGHT_BOTTOM_Y int16
	RIGHT_TOP_X    int16
	RIGHT_TOP_Y    int16
	LEFT_BOTTOM_X  int16
	LEFT_BOTTOM_Y  int16
	LEFT_TOP_X     int16 = 0
	LEFT_TOP_Y     int16 = 0
)

func setup() {
	KillProcess("corner-invoke")
	sx, sy := robotgo.GetScreenSize()
	fmt.Println("get screen size: ", sx, sy)
	RIGHT_BOTTOM_X, RIGHT_BOTTOM_Y = int16(sx-1), int16(sy-1)
}

func handleMouseDown() {
	hook.Register(hook.MouseDown, nil, func(e hook.Event) {
		log.Println(e)
		if e.Button == 3 {
			HideAllWindows()
		}
	})

	hook.Register(hook.MouseWheel, nil, func(e hook.Event) {
		log.Println(e)

		HideAllWindows()

	})

}

func handleMouseMove() {
	hook.Register(hook.MouseMove, nil, func(e hook.Event) {
		// log.Println(e)
		if e.X == RIGHT_BOTTOM_X && e.Y == RIGHT_BOTTOM_Y {
			LockScreen()
			robotgo.Move(0, 0)
		}
	})
}

func main() {
	setup()
	handleMouseDown()

	handleMouseMove()
	s := hook.Start()
	<-hook.Process(s)
}
