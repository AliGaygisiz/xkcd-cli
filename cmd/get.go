package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

func downloadAndSave(comic XKCDComic) error {
	fileName := "xkcd_" + strconv.Itoa(comic.Num) + ".png"

	_, err := os.Stat(fileName)

	if !os.IsNotExist(err) {
		fmt.Println("File already exists")
		fmt.Println("Do you want to overwrite it? (y/n)")
		input := ""

		if _, err = fmt.Scanln(&input); err != nil {
			return fmt.Errorf("error reading input: %v", err)
		}

		if input != "y" && input != "Y" {
			os.Exit(0)
		}
	}

	resp, err := http.Get(comic.Img)
	if err != nil {
		return fmt.Errorf("error downloading comic: %v", err)
	}

	defer resp.Body.Close()

	out, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}

func getComic(number int) error {
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

	fmt.Println("========================")
	fmt.Println("Comic number : ", comic.Num)
	fmt.Println("Title        : ", comic.Title)
	fmt.Println("========================")

	err = downloadAndSave(comic)
	if err != nil {
		fmt.Println("Error downloading comic: ", err)
	}

	fmt.Println("Comic downloaded successfully")
	return nil
}

func getGivenComic(c *cli.Context) error {
	number, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return fmt.Errorf("error parsing comic number: %v", err)
	}
	if number < 1 {
		fmt.Println("Comic number should be greater than 0")
	}
	return getComic(number)
}

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "Download any xkcd comic into current directory",
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
				return getGivenComic(c)
			}
		},
		Subcommands: []*cli.Command{
			{
				Name:  "latest",
				Usage: "Download the latest comic",
				Action: func(c *cli.Context) error {
					return getComic(0)
				},
			},
			{
				Name:  "random",
				Usage: "Download a random comic",
				Action: func(c *cli.Context) error {
					number, err := pickRandomComic()
					if err != nil {
						return err
					}
					return getComic(number)
				},
			},
			{
				Name:  "[number]",
				Usage: "Download a specific comic by number",
				Action: func(c *cli.Context) error {
					return getGivenComic(c)
				},
			},
		},
	}
}
