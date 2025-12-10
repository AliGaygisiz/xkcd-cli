package main

import (
	"context"
	"fmt"
	"os"

	"github.com/AliGaygisiz/xkcd-cli/cmd"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{}
	app.Name = "xkcd-cli"
	app.Version = "0.3.1"
	app.HideVersion = false
	app.Usage = "A CLI Application for XKCD comics"
	app.CommandNotFound = func(ctx context.Context, cmd *cli.Command, command string) {
		fmt.Fprintf(os.Stderr, "xkcd-cli: unknown command '%s'. See 'xkcd-cli --help'\n", command)
	}

	app.OnUsageError = func(ctx context.Context, cmd *cli.Command, err error, isSubcommand bool) error {
		fmt.Fprintf(os.Stderr, "xkcd-cli: %v\n", err)
		return cli.Exit("See '--help' for usage", 2)
	}

	app.Commands = []*cli.Command{
		cmd.DisplayCommand(),
		cmd.GetCommand(),
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		panic(err)
	}
}
