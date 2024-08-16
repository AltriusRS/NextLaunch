package tsd

import (
	"Nextlaunch/src/config"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

type SnapiClient struct {
	BaseClient
}

func NewSnapiClient() *SnapiClient {
	c := &SnapiClient{
		BaseClient: BaseClient{
			queue:      make(chan *ClientRequest),
			ticker:     time.NewTicker(time.Millisecond * 100),
			shouldTick: true,
			client:     &http.Client{},
		},
	}

	go func() {
		for req := range c.queue {
			<-c.ticker.C          // Wait for the next tick
			c.ProcessRequest(req) // Process the request
		}
	}()

	return c
}

func (c *SnapiClient) GetNewsArticles(limit int, offset int) *[]NewsArticle {
	res, err := c.Get(config.SNAPIFullBaseURL + config.SNAPIVersion + "articles?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))

	if err != nil {
		c.logger.Fatal(err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.logger.Fatal(err)
		}
	}(res.Body)
	if res.StatusCode != 200 {
		c.logger.Fatal(err)
		return nil
	}

	var response SNAPINewsResponse

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)

	if err != nil {
		c.logger.Fatal(err)
		return nil
	}
	return &response.Results
}
