package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Input struct {
	LineIn string
}

func getSource() *bufio.Scanner, error {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {

		return bufio.NewScanner(os.Stdin)

	} else {
		
		if 
		
	}

}

func NewCommander(commandChan chan string) {

	defer func() { close(commandChan) }()
	
	
	for _, command := range source {
		commandChan <- command
	}
}
