package main

import (
	"fmt"
	"os"

	"github.com/AliGaygisiz/xkcd-cli/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "xkcd-cli"
	app.Version = "0.1.0"
	app.HideVersion = false
	app.Usage = "A CLI Application for XKCD comics"
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(os.Stderr, "xkcd-cli: unknown command '%s'. See 'xkcd-cli --help'\n", command)
	}
	app.Commands = []*cli.Command{
		cmd.GetCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}