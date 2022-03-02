//go:build darwin
// +build darwin

package main

func ScreenIsLocked() bool {
	return true
}

func HideAllWindows() {
	// robotgo.KeyTap("d", "cmd")
}

func LockScreen() {

}
