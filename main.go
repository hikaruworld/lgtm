package main

import (
	"os"

	"github.com/urfave/cli"
)

const version string = "0.1.0"

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "lgtm"
	app.Usage = "get lgtm image url(supported private repo)."
	app.Version = version
	app.Author = "hikaruworld"
	app.Email = "hikaruworld@gmail.com"
	app.Commands = Commands
	return app
}
