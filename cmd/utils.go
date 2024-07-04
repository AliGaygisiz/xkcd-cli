package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
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

func fetchLatestComic() (XKCDComic, error) {
	var comic XKCDComic

	url := "https://xkcd.com/info.0.json"

	resp, err := http.Get(url)
	if err != nil {
		return comic, fmt.Errorf("error reaching url: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return comic, fmt.Errorf("error reading response: %v", err)
	}

	err = json.Unmarshal(body, &comic)
	if err != nil {
		return comic, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return comic, err
}

func fetchComicByNumber(number int) (XKCDComic, error) {
	tempComic, err := fetchLatestComic()
	if err != nil {
		return tempComic, err
	}

	if number > tempComic.Num {
		fmt.Printf("Comic number %d does not exist\n", number)
		fmt.Printf("Pick a number between 1-%d\n", tempComic.Num)
		os.Exit(1)
	}

	var comic XKCDComic

	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", number)

	resp, err := http.Get(url)
	if err != nil {
		return comic, fmt.Errorf("error reaching url: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return comic, fmt.Errorf("error reading response: %v", err)
	}

	err = json.Unmarshal(body, &comic)
	if err != nil {
		return comic, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return comic, err
}

func pickRandomComic() (int, error) {
	tempComic, err := fetchLatestComic()
	if err != nil {
		return 0, err
	}

	number := rand.Intn(tempComic.Num) + 1
	return number, nil
}
