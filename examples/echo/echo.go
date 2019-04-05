package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/haorenfsa/go-cmdline-app/app"
)

// Flags are the flags given in cmdline
type Flags struct {
	Banner *string
}

// APP is a example implementation of APP interface
type APP struct {
	flags *Flags
}

// NewAPP is the constuctor of APP
func NewAPP() *APP {
	return &APP{
		flags: &Flags{},
	}
}

// SetFlags implements APP.SetFlags
func (a *APP) SetFlags() {
	a.flags.Banner = flag.String("b", "==============Thank you for using cmdline ==============\n", "the starting banner will be print to stdout at begin")
}

// Run implements APP.Run
func (a *APP) Run() {
	fmt.Print(*a.flags.Banner)
}

// RunOnCmdSet implements CmdSetRunner.RunOnCmdSet
func (a *APP) RunOnCmdSet(cmdSet app.CmdSet) {
	if cmdSet[0] == "quit" {
		os.Exit(0)
	} else {
		for _, cmd := range cmdSet {
			fmt.Print(cmd + " ")
		}
		fmt.Print("\n")
	}
}

func main() {
	echoAPP := NewAPP()
	echoAPP.SetFlags()
	flag.Parse()
	go echoAPP.Run()
	app.StartDealWithCmdSets(echoAPP)
}
