package main

import (
	"os/exec"

	log "github.com/ivan-kostko/conc/pkg/logger"
)

func NewWorker(commands chan string, log log.FieldLogger) func() {
	return func() {
		for command := range commands {

			log.WithField("command", command)

			log.Debug("Running command")
			if outpt, err := exec.Command(command).Output(); err != nil {
				log.Error(err)
			} else {
				log.Debugf("Command returned %s\r\n", outpt)
			}
		}
	}
}
