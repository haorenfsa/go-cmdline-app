package main

import (
	"flag"

	APP "github.com/haorenfsa/go-cmdline-app/app"
	"github.com/haorenfsa/go-cmdline-app/examples"
)

// main is a template main func
func main() {
	app := examples.NewAPP()
	app.SetFlags()
	flag.Parse()
	go app.Run()
	APP.StartDealWithCmdSets(app)
}
