package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/urfave/cli/v2"
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

func displayGivenComic(c *cli.Context) error {
	number, err := strconv.Atoi(c.Args().First())
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
		Usage: "Display any xkcd comic withoud downloading",
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			fmt.Fprintf(os.Stderr, "xkcd-cli: %v\n", err)
			return cli.Exit("See '--help' for usage", 2)
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				cli.ShowSubcommandHelp(c)
				return nil
			} else {
				_, err := strconv.Atoi(c.Args().First())
				if err != nil {
					fmt.Println("Unknown command. See 'xkcd get help' for available commands")
					os.Exit(1)
				}
				return displayGivenComic(c)
			}
		},
		Subcommands: []*cli.Command{
			{
				Name:  "latest",
				Usage: "Display the latest comic",
				Action: func(c *cli.Context) error {
					return DisplayComic(0)
				},
			},
			{
				Name:  "random",
				Usage: "Display a random comic",
				Action: func(c *cli.Context) error {
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
				Action: func(c *cli.Context) error {
					return displayGivenComic(c)
				},
			},
		},
	}
}
