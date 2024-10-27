package tsd

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/logging"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type LL2Client struct {
	BaseClient
}

func NewLL2Client() *LL2Client {
	c := &LL2Client{
		BaseClient: BaseClient{
			queue:      make(chan *ClientRequest),
			ticker:     time.NewTicker(time.Millisecond * 100),
			shouldTick: true,
			client:     &http.Client{},
			logger:     logging.NewLogger("LL2 Client"),
		},
	}

	go func() {
		for req := range c.queue {
			<-c.ticker.C // Wait for the next tick

			if config.Config.LaunchLibrary.LaunchLibraryKey != "" {
				req.req.Header.Set("Authorization", "Token "+config.Config.LaunchLibrary.LaunchLibraryKey)
			}

			c.ProcessRequest(req) // Process the request
		}
	}()

	return c
}

func (c *LL2Client) GetLaunches(limit int, offset int) *[]LL2Launch {
	c.logger.Debug("Getting launches")
	res, err := c.Get(config.LL2FullBaseURL + "launches/upcoming?format=json&mode=detailed&limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))

	c.logger.Debug("Got response")

	if err != nil {
		c.logger.Fatal(err)
		return nil
	}

	c.logger.Debug("Reading response")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.logger.Fatal(err)
		}
	}(res.Body)

	c.logger.Debug("Checking status code")

	if res.StatusCode != 200 {
		c.logger.Fatal(fmt.Errorf("status code %d", res.StatusCode))
		return nil
	}

	c.logger.Debug("Decoding response")

	var response LL2LaunchesResponse

	decoder := json.NewDecoder(res.Body)
	c.logger.Debug("Processing response")
	err = decoder.Decode(&response)

	c.logger.Debug("Decoded response")

	if err != nil {
		c.logger.Fatal(err)
		return nil
	}

	c.logger.Debug("Got " + strconv.Itoa(len(response.Results)) + " launches")

	return &response.Results
}
