package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

type XKCDComic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

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

func getComic() error {
	url := "https://xkcd.com/info.0.json"

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error getting comic: %v", err)
	} else {

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading response: %v", err)
		}

		var comic XKCDComic
		err = json.Unmarshal(body, &comic)
		if err != nil {
			return fmt.Errorf("error unmarshalling response: %v", err)
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
	}
	return nil
}

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "Download the latest xkcd comic into current directory",
		Action: func(c *cli.Context) error {
			return getComic()
		},
	}
}
