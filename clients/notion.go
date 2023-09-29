package clients

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type NotionConfig struct {
	apiKey string
	databaseId string
	viewId string
}

type NotionClient struct {
	config NotionConfig
	client *http.Client
}

func NewNotionClient(apiKey string, databaseId string, viewId string) *NotionClient {
	return &NotionClient{
		config: NotionConfig{
			apiKey: apiKey,
			databaseId: databaseId,
			viewId: viewId,
		},
		client: &http.Client{},
	}
}

func (c *NotionClient) GetDatabase() {
	path := "https://api.notion.com/v1/databases/" + c.config.databaseId
	req, _ := http.NewRequest("GET", path, nil)
	req.Header.Set("Authorization", "Bearer " + c.config.apiKey)
	req.Header.Set("Notion-Version", "2022-06-28")

	res, err := c.client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	fmt.Println(string(body))
}
