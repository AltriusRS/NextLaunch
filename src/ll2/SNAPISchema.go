package ll2

type SNAPIResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type SNAPINewsResponse struct {
	SNAPIResponse
	Results []NewsArticle `json:"results"`
}

type NewsArticle struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	ImageURL    string `json:"image_url"`
	NewsSite    string `json:"news_site"`
	Summary     string `json:"summary"`
	PublishedAt string `json:"published_at"`
	UpdatedAt   string `json:"updated_at"`
	Featured    bool   `json:"featured"`
	Launches    []NewsLaunch
	Events      []NewsEvent
}

type NewsLaunch struct {
	ID       string `json:"launch_id"`
	Provider string `json:"provider"`
}

type NewsEvent struct {
	ID       int    `json:"event_id"`
	Provider string `json:"provider"`
}

type SNAPIReportsResponse struct {
	SNAPIResponse
	Results []Report `json:"results"`
}

type Report struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	ImageURL    string `json:"image_url"`
	NewsSite    string `json:"news_site"`
	Summary     string `json:"summary"`
	PublishedAt string `json:"published_at"`
	UpdatedAt   string `json:"updated_at"`
}
