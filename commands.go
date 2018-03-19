package main

import (
	"log"
	"os/exec"
)

func startImperialMarch() {
	cmd := exec.Command("spotify", "play", "uri", "spotify:track:6ht3MI8a40UT5haN7t9StM")
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}

func stopImperialMarch() {
	cmd := exec.Command("spotify", "stop")
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}
