package main

import (
	"log"
	"os/exec"

	hook "github.com/robotn/gohook"
)

const (
	RIGHT_BOTTOM_X = 3039
	RIGHT_BOTTOM_Y = 899
)

func lockScreen() {
	if isLocked := ScreenIsLocked(); !isLocked {
		cmd := exec.Command("rundll32.exe", "user32.dll,LockWorkStation")
		err := cmd.Start()

		if err != nil {
			log.Printf("lockScreen err:%v", err)
		}
	}
}

func main() {
	log.Println("press combo keys")
	// robotgo.KeyTap("m", "alt", "command")
	hook.Register(hook.MouseDown, nil, func(e hook.Event) {
		log.Println(e)
	})
	s := hook.Start()
	<-hook.Process(s)
	// for {
	// 	x, y := robotgo.GetMousePos()

	// 	if x == RIGHT_BOTTOM_X && y == RIGHT_BOTTOM_Y {
	// 		lockScreen()
	// 		robotgo.Move(0, 0)
	// 	}
	// 	time.Sleep(200 * time.Millisecond)
	// }

}
