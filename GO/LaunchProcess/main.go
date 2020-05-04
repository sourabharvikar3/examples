// This example is tested on windows 10 x64

package main

import (
	"fmt"
	"os/exec"
)

func LaunchNotepad() {
	appToLaunch := "c:\\Windows\\System32\\notepad.exe"
	cmd := exec.Command(appToLaunch)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("failed to launch %s, error - %s\n", appToLaunch, err)
	}
}

// Launches CMD with pointer to current dir
func LaunchCmd() {
	appToLaunch := "c:\\Windows\\System32\\conhost.exe"
	cmd := exec.Command(appToLaunch)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("failed to launch %s, error - %s\n", appToLaunch, err)
	}
}

func main() {
	LaunchNotepad()
	LaunchCmd()
}
