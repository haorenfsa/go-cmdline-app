package main

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
