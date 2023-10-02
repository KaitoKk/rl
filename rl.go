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
