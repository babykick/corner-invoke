package main

import (
	"log"
	"os/exec"
	"syscall"
	"time"
	"unsafe"

	"github.com/go-vgo/robotgo"
)

const (
	DESKTOP_SWITCHDESKTOP = 0x0100 // The access to the desktop
	RIGHT_BOTTOM_X        = 3039
	RIGHT_BOTTOM_Y        = 899
)

func ScreenIsLocked() bool {

	// load user32.dll only once
	user32 := syscall.MustLoadDLL("user32.dll")
	openDesktop := user32.MustFindProc("OpenDesktopW")
	closeDesktop := user32.MustFindProc("CloseDesktop")
	switchDesktop := user32.MustFindProc("SwitchDesktop")

	var lpdzDesktopPtr uintptr = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Default"))) //string
	var dwFloatsPtr uintptr = 0                                                               //uint32
	var fInheritPtr uintptr = 0                                                               //boolean
	var dwDesiredAccessPtr uintptr = uintptr(DESKTOP_SWITCHDESKTOP)                           //uint32

	r1, _, _ := syscall.Syscall6(openDesktop.Addr(), 4, lpdzDesktopPtr, dwFloatsPtr, fInheritPtr, dwDesiredAccessPtr, 0, 0)
	if r1 == 0 {
		panic("get desktop locked status error")
	}

	res, _, _ := syscall.Syscall(switchDesktop.Addr(), 1, r1, 0, 0)
	// clean up
	syscall.Syscall(closeDesktop.Addr(), 1, r1, 0, 0)

	return res != 1
}

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

	for {
		x, y := robotgo.GetMousePos()

		if x == RIGHT_BOTTOM_X && y == RIGHT_BOTTOM_Y {
			lockScreen()
			robotgo.Move(0, 0)
		}
		time.Sleep(200 * time.Millisecond)
	}

}
