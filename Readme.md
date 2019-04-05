# go cmdline app
## Purpose
This package is mean to help us quickly establish an cmdline app with basic functions like
- run a main function
- read flags in cmdline input arguments
- detect what user has input in stdin, and deal with them

## How to use
- just implement interfaces in app/interfaces.go
- and use main func to run them
- there is a template of main in main_template.go
- there is also a echo app example in examples folder
- and a template of app implementation
```go
import (
	"flag"

	"github.com/haorenfsa/go-cmdline-app/app"
	"github.com/haorenfsa/go-cmdline-app/templates"
)

// main is a template main func
func main() {
	templateAPP := templates.NewAPP()
	templateAPP.SetFlags()
	flag.Parse()
	go templateAPP.Run()
	app.StartDealWithCmdSets(templateAPP)
}
```

```go
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
	}
}
```