package clients

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const API_BASE_URL = "https://api.notion.com/v1/"

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

func (c NotionClient) buildRequest(method string, path string) *http.Request {
	uri := API_BASE_URL + path
	req, _ := http.NewRequest(method , uri, nil)
	req.Header.Set("Authorization", "Bearer " + c.config.apiKey)
	req.Header.Set("Notion-Version", "2022-06-28")

	return req
}

func (c NotionClient) GetDatabase() {
	path := "databases/" + c.config.databaseId
	req := c.buildRequest("GET", path)

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
