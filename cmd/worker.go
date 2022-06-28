package main

import (
	"os/exec"
	log "github.com/ivan-kostko/conc/pkg/logger"
)

func NewWorker(commands chan string, log log.Logger) {
	for _, command := range <-commands {
		
		log.WithField("command", command)
		log.Debug("Preparing command")
		cmd := exec.Command(string(command))
		
		log.Debug("Running command")
		if err := cmd.Run() {
			log.Error(err)
		}
		log.Debugf("Command returned %s", cmd.Output())

	}
}
