package tsd

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/logging"
	"encoding/json"
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

// TestToken TODO: Implement api key check, not that this should be a problem,
// since this is moving towards a custom API that provides a proxy to LL2 using our own token
func (c *LL2Client) TestToken(token string) bool {
	return true
}

func (c *LL2Client) GetLaunches(limit int, offset int) *[]LL2Launch {
	c.logger.Debug("Getting launches")
	res, err := c.Get(config.LL2FullBaseURL + "launches/upcoming?format=json&mode=detailed&limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))

	c.logger.Debug("Got response")

	if err != nil {
		c.logger.Error(err)
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
		c.logger.Errorf("Status code %d", res.StatusCode)
		return nil
	}

	c.logger.Debug("Decoding response")

	var response LL2LaunchesResponse

	decoder := json.NewDecoder(res.Body)
	c.logger.Debug("Processing response")
	err = decoder.Decode(&response)

	c.logger.Debug("Decoded response")

	if err != nil {
		c.logger.Error(err)
		return &response.Results
	}

	c.logger.Debug("Got " + strconv.Itoa(len(response.Results)) + " launches")

	return &response.Results
}

// TODO: Implement update API support
//func (c *LL2Client) GetUpdates(limit int, offset int) *[]LL2Update {
//	c.logger.Debug("Getting updates")
//	res, err := c.Get(config.LL2FullBaseURL + "updates/upcoming?format=json&mode=detailed&limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))
//
//	c.logger.Debug("Got response")
//
//	if err != nil {
//		c.logger.Error(err)
//		return nil
//	}
//
//	c.logger.Debug("Reading response")
//
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//			c.logger.Fatal(err)
//		}
//	}(res.Body)
//
//	c.logger.Debug("Checking status code")
//
//	if res.StatusCode != 200 {
//		c.logger.Errorf("Status code %d", res.StatusCode)
//		return nil
//	}
//
//	c.logger.Debug("Decoding response")
//
//	var response LL2UpdatesResponse
//
//	decoder := json.NewDecoder(res.Body)
//	c.logger.Debug("Processing response")
//	err = decoder.Decode(&response)
//
//	c.logger.Debug("Decoded response")
//
//	if err != nil {
//		c.logger.Error(err)
//		return &response.Results
//	}
//
//	c.logger.Debug("Got " + strconv.Itoa(len(response.Results)) + " updates")
//
//	return &response.Results
//}
