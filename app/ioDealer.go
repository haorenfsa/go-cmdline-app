package app

import (
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

// StartDealWithCmdSets starts deal with cmdline inputs
func StartDealWithCmdSets(runner CmdSetRunner) {
	cmdChan := make(chan CmdSet, 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		loopReadCmd(cmdChan)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		loopExecuteCmd(cmdChan, runner)
		wg.Done()
	}()
	wg.Wait()
}

func loopReadCmd(cmdChan chan CmdSet) {
	var cmdReader io.Reader
	cmdReader = os.Stdin
	buffer := make([]byte, 100)
	for {
		len, err := cmdReader.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		inputStr := string(buffer[:len])
		inputLine := strings.Trim(inputStr, "\n")
		cmds := strings.Split(inputLine, " ")
		cmdChan <- cmds
	}
}

func loopExecuteCmd(cmdChan chan CmdSet, runner CmdSetRunner) {
	for {
		select {
		case cmds := <-cmdChan:
			runner.RunOnCmdSet(cmds)
		}
	}
}
