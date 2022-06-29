package main

import (
	"flag"
	"sync"

	log "github.com/ivan-kostko/conc/pkg/logger"
)

var (
	buffer          *int
	procNumber      *int
	fileName        *string
	commandTemplate *string
	logger          = log.DefaultLogger
)

func init() {
	buffer = flag.Int("buffer", 100, "size of the buffer between reader and executor")
	procNumber = flag.Int("n", 10, "number of concurrent executors")
	fileName = flag.String("f", "", "the name of source file. Omitted, if piped")
	commandTemplate = flag.String("cmd", "", "command template to be executed")
}

func main() {

	logger.Debug("parsing flags")
	flag.Parse()

	logger.Debug("setting channel")

	if buffer == nil {
		logger.Debug("As buffer was nil, setting it to a new integer")
		buffer = new(int)
	}

	commandChan := make(chan string, *buffer)

	var wg sync.WaitGroup

	logger.Debug("setting workers")
	for i := 0; i < *procNumber; i++ {

		logger.Debugf("setting worker %i", i)
		wg.Add(1)
		log := logger.WithField("workerNo", i)
		worker := NewWorker(commandChan, log)
		go func() {
			defer func() { log.Debug("worker defferring"); wg.Done() }()

			log.Debug("starting worker")
			worker()
		}()
	}

	go func() {
		logger.Debug("launching commander")
		NewCommander(commandChan)
		logger.Debug("commander is done")

	}()

	logger.Debug("waiting for workers to complete")
	wg.Wait()

	logger.Debug("workers to completed")

	return
}
