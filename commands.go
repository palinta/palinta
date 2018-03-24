package main

import (
	"log"
	"os/exec"
)

var isRunning = false

func startImperialMarch() {
	if isRunning {
		return
	}
	cmd := exec.Command("spotify", "play", "uri", "spotify:track:6ht3MI8a40UT5haN7t9StM")
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		return
	}
	isRunning = true
}

func stopImperialMarch() {
	if !isRunning {
		return
	}
	cmd := exec.Command("spotify", "stop")
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		return
	}
	isRunning = false
}
