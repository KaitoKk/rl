package main

import (
	"rl/clients"
	"rl/models"
)

const ( // MEMO: 仮置き
	apiKey = "..."
	databaseId = "aaa"
	viewId = "bbb"
)

func main() {
	c := clients.NewNotionClient(
		apiKey,
		databaseId,
		viewId,
	)

	article := models.Article{
		Title: "test",
		Link: "https://example.com",
	}

	c.PostArticle(article)
}
// "https://www.notion.so/1f99b3341629486a961b5a105e8b41d1?v=444d419cb5da4b7e8b26a0d200252a81&pvs=4"