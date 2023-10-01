package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"rl/clients"
	"rl/models"
	"rl/scraper"
)

const ( // MEMO: 仮置き
	apiKey = "..."
	databaseId = "aaa"
	viewId = "bbb"
)

func main() {
	if err := checkArgs(); err != nil {
		log.Fatalf("Error checking args: %v", err)
	}

	url := os.Args[1]
	title, err := scraper.FetchTitle(url)
	if err != nil {
		log.Fatalf("Error fetching title: %v", err)
	}
	fmt.Println(title)

	c := clients.NewNotionClient(
		apiKey,
		databaseId,
		viewId,
	)

	article := models.Article{
		Title: title,
		Link: url,
	}

	c.PostArticle(article)
}

func checkArgs() error {
	args := os.Args
	if len(args) < 2 {
		return errors.New("引数がありません")
	}

	if !isValidUrl(args[1]) {
		return errors.New("URLが不正です")
	}
	return nil
}

func isValidUrl(url string) bool {
	re := regexp.MustCompile(`^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`)
	return re.MatchString(url)
}
// "https://www.notion.so/1f99b3341629486a961b5a105e8b41d1?v=444d419cb5da4b7e8b26a0d200252a81&pvs=4"