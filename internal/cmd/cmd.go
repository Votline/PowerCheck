package cmd

import (
	"log"
	"os/exec"
	"runtime"
)

func Shutdown() error {
	switch runtime.GOOS {
	case "linux", "darwin":
		return exec.Command("shutdown", "-h", "now").Run()
	case "windows":
		return exec.Command("shutdown", "/s", "/t", "0").Run()
	default:
		log.Println("unsupported platform")
	}
	return nil
}

func Suspend() error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("dbus-send", "--system", "--print-reply",
			"--dest=org.freedesktop.login1",
			"/org/freedesktop/login1",
			"org.freedesktop.login1.Manager.Suspend",
			"boolean:true").Run()
	case "darwin":
		return exec.Command("pmset", "sleepnow").Run()
	case "windows":
		return exec.Command("rundll32.exe", "powrprof.dll,SetSuspendState", "0,1,0").Run()
	default:
		log.Println("unsupported platform")
	}
	return nil
}
