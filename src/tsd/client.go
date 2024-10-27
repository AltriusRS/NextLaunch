package tsd

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/logging"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// BaseClient is a struct that contains the client for the whenplane API
type BaseClient struct {
	queue      chan *ClientRequest
	ticker     *time.Ticker
	shouldTick bool
	client     *http.Client
	logger     *logging.Logger
}

type ClientRequest struct {
	req *http.Request
	res *http.Response
	err error
	c   chan bool
}

func (r *ClientRequest) Request() *http.Request {
	return r.req
}

func (r *ClientRequest) Response() *http.Response {
	return r.res
}

func (r *ClientRequest) Error() error {
	return r.err
}

func (r *ClientRequest) Done() bool {
	select {
	case <-r.c:
		return true
	default:
		return false
	}
}

func (r *ClientRequest) Wait() {
	<-r.c
}

func (r *ClientRequest) Callback() {
	r.c <- true
}

func NewRequest(req *http.Request) *ClientRequest {
	return &ClientRequest{
		req: req,
		c:   make(chan bool),
	}
}

// NewClient returns a new BaseClient
func NewClient() *BaseClient {
	c := &BaseClient{
		queue:      make(chan *ClientRequest),
		ticker:     time.NewTicker(time.Millisecond * 100),
		shouldTick: true,
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        10,
				IdleConnTimeout:     time.Second * 60,
				DisableCompression:  true,
				TLSHandshakeTimeout: time.Second * 5,
			},
			//CheckRedirect: func(req *http.Request, via []*http.Request) error {
			//	return http.ErrUseLastResponse
			//},
			Timeout: time.Second * 5,
		},
		logger: logging.NewLogger("HTTP Client"),
	}

	go func() {
		for req := range c.queue {
			<-c.ticker.C          // Wait for the next tick
			c.ProcessRequest(req) // Process the request
		}
	}()

	return c
}

// ProcessRequest adds a request to the queue
func (c *BaseClient) ProcessRequest(r *ClientRequest) {
	c.logger.Debug("Processing request")
	// Set the required headers for every request
	r.req.Header.Set("Accept", "application/json")

	// You may be warned this is "always true" it is not, as it is conditionally set by the compiler
	if //goland:noinspection GoBoolExpressions,GoBoolExpressions
	config.IsDev {
		r.req.Header.Set("User-Agent", "NextLaunch/"+config.Version+"-"+config.BuildDate+" DEVELOPMENT BUILD")
	} else {
		r.req.Header.Set("User-Agent", "NextLaunch/"+config.Version+"-"+config.BuildDate)
	}

	c.logger.Debug("Sending request to " + r.req.URL.String())
	c.logger.Debug("Headers")
	for k, v := range r.req.Header {
		c.logger.Debug("> " + k + ": " + v[0])
	}

	resp, err := c.client.Do(r.req)

	c.logger.Debug("Got response")
	r.res = resp
	r.err = err
	c.logger.Debug("Callback")
	r.Callback()
}

// Ticker returns the ticker for the BaseClient
func (c *BaseClient) Ticker() *time.Ticker {
	return c.ticker
}

// Queue returns the queue for the BaseClient
func (c *BaseClient) Queue() chan *ClientRequest {
	return c.queue
}

// ShouldTick returns the shouldTick for the BaseClient
func (c *BaseClient) ShouldTick() bool {
	return c.shouldTick
}

// Close shuts the client down immediately
func (c *BaseClient) Close() {
	c.ticker.Stop()
	close(c.queue)
	c.shouldTick = false
}

func (c *BaseClient) Get(url string) (*http.Response, error) {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req := NewRequest(r)
	c.queue <- req
	req.Wait()
	return req.Response(), req.Error()
}

// GetNewsArticles returns a list of NewsArticles from the given limit and offset
func (c *BaseClient) GetNewsArticles(limit int, offset int) []NewsArticle {
	res, err := c.Get(config.SNAPIBaseURL + config.SNAPIVersion + "/news?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))

	if err != nil {
		log.Println("Error getting news articles from SNAPI - Limit: ", limit, "Offset: ", offset)
		return []NewsArticle{}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing news articles from SNAPI - Limit: ", limit, "Offset: ", offset)
		}
	}(res.Body)
	if res.StatusCode != 200 {
		log.Println("Error getting news articles from SNAPI - Limit: ", limit, "Offset: ", offset)
		return []NewsArticle{}
	}

	var response SNAPINewsResponse

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)

	if err != nil {
		log.Println("Error decoding news articles from SNAPI - Limit: ", limit, "Offset: ", offset)
		return []NewsArticle{}
	}
	return response.Results
}
