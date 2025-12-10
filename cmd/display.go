package cmd

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/urfave/cli/v3"
)

func displayBySystem(file *os.File) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", file.Name())
	case "darwin":
		cmd = exec.Command("open", file.Name())
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", file.Name())
	default:
		return fmt.Errorf("unsupported platform")
	}

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("error starting command: %v", err)
	}

	return nil
}

func DisplayComic(number int) error {
	var comic XKCDComic
	var err error

	if number == 0 {
		comic, err = fetchLatestComic()
		if err != nil {
			return err
		}
	} else {
		comic, err = fetchComicByNumber(number)
		if err != nil {
			return err
		}
	}

	resp, err := http.Get(comic.Img)
	if err != nil {
		return fmt.Errorf("error downloading comic: %v", err)
	}

	defer resp.Body.Close()

	file, err := os.CreateTemp("", "xkcd_"+strconv.Itoa(comic.Num)+"_*.png")
	if err != nil {
		return fmt.Errorf("error creating temporary file: %v", err)
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Println("========================")
	fmt.Println("Comic number : ", comic.Num)
	fmt.Println("Title        : ", comic.Title)
	fmt.Println("========================")

	err = displayBySystem(file)
	if err != nil {
		return fmt.Errorf("error displaying comic: %v", err)
	}

	return nil
}

func displayGivenComic(ctx context.Context, cmd *cli.Command) error {
	number, err := strconv.Atoi(cmd.Args().First())
	if err != nil {
		return fmt.Errorf("error parsing comic number: %v", err)
	}
	if number < 1 {
		fmt.Println("Comic number should be greater than 0")
	}
	return DisplayComic(number)
}

func DisplayCommand() *cli.Command {
	return &cli.Command{
		Name:  "display",
		Usage: "Display any xkcd comic without downloading",
		OnUsageError: func(ctx context.Context, cmd *cli.Command, err error, isSubcommand bool) error {
			fmt.Fprintf(os.Stderr, "xkcd-cli: %v\n", err)
			return cli.Exit("See '--help' for usage", 2)
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if !cmd.Args().Present() {
				cli.ShowSubcommandHelp(cmd)
				return nil
			} else {
				_, err := strconv.Atoi(cmd.Args().First())
				if err != nil {
					fmt.Fprintf(os.Stderr, "Unknown command. See 'xkcd display help' for available commands\n")
					return nil
				}
				return displayGivenComic(ctx, cmd)
			}
		},
		Commands: []*cli.Command{
			{
				Name:  "latest",
				Usage: "Display the latest comic",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return DisplayComic(0)
				},
			},
			{
				Name:  "random",
				Usage: "Display a random comic",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					number, err := pickRandomComic()
					if err != nil {
						return err
					}
					return DisplayComic(number)
				},
			},
			{
				Name:  "[number]",
				Usage: "Display a specific comic by number",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return displayGivenComic(ctx, cmd)
				},
			},
		},
	}
}
