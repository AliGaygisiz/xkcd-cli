package main

import (
	"os"

	"github.com/AliGaygisiz/xkcd-cli/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "xkcd-cli"
	app.Usage = "A CLI Application for XKCD comics"
	app.Commands = []*cli.Command{
		cmd.GetCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
