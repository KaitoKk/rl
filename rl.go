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

var (
	apiKey     = os.Getenv("NOTION_API_KEY")
	databaseId = os.Getenv("NOTION_DATABASE_ID")
)

func main() {
	if err := checkVar(); err != nil {
		log.Fatalf("Error checking env var: %v", err)
	}

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
	)

	article := models.Article{
		Title: title,
		Link:  url,
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

func checkVar() error {
	if apiKey == "" {
		return errors.New("環境変数NOTION_API_KEYがありません")
	}
	if databaseId == "" {
		return errors.New("環境変数NOTION_DATABASE_IDがありません")
	}
	return nil
}
