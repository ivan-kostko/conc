package main

import (
	"flag"
	log "github.com/ivan-kostko/conc/pkg/logger"
)

var (
	buffer          *int
	procNumber      *int
	commandTemplate *string
	log = log.DefaultLogger
)

func init() {
	buffer = flag.Int("buffer", 100, "size of the buffer between reader and executor")
	procNumber = flag.Int("n", 10, "number of concurrent executors")
	commandTemplate = flag.String("cmd", "", "command template to be executed")
}

func main() {
	flag.Parse()

	log := 

	commandChan := make(chan string, *buffer)

	wg := waitGroup

	for i := 0; i< *procNumber; i++ {
		wg.Add(1)
		go func(){
			
		}
	}

	return
}

func NewCommander(commandChan chan string)  {
	
	defer func(){ close(commandChan) }
	source := []string{"env", "ls", "pwd"}

	for _, command := range source {
		commandChan <- command
	}
}
