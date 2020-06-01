package destinations

type embed struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Url string `json:"url"`
	Timestamp string `json:"timestamp"`
	Color int64 `json:"color"`
	Footer embedFooter `json:"footer"`
	Image embedImage `json:"image"`
}
type embedFooter struct {
	Text string `json:"text"`
	IconUrl string `json:"icon_url"`
}

type embedField struct {
	Name string `json:"name"`
	Value string `json:"Value"`
	Inline bool `json:"inline"`
}

type embedImage struct {
	Url string `json:"url"`
	Height int `json:"height"`
	Width int `json:"width"`
}

type discordMessage struct {
	Content string `json:"content"`
	Embed   embed  `json:"embed"`
}

