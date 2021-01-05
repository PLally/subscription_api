package types

import (
	"github.com/dyatlov/go-opengraph/opengraph"
	"github.com/mmcdole/gofeed"
	"github.com/plally/subscription_api/subscription"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func RegisterRSS() {
	subscription.SetSubTypeHandler("rss", &RSSHandler{})
}

type RSSHandler struct{}

func (r *RSSHandler) GetNewItems(tags string) []subscription.SubscriptionItem {
	var items []subscription.SubscriptionItem
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(tags)
	if err != nil {
		log.Error(err)
		log.Debug(tags)
	}

	for _, item := range feed.Items {
		subItem := subscription.SubscriptionItem{
			Title:       item.Title,
			Url:         item.Link,
			Image:       getImage(item),
			Description: "-",
			Author:      item.Author.Name,
			TimeID:      item.PublishedParsed.Unix(),
		}
		items = append(items, subItem)
	}
	return items
}

func (r *RSSHandler) Validate(tags string) (string, error) {
	return tags, nil
}

func getImage(item *gofeed.Item) string {
	if item.Image != nil && item.Image.URL == "" {
		return item.Image.URL
	}

	resp, err := http.Get(item.Link)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	og := opengraph.NewOpenGraph()
	_ = og.ProcessHTML(resp.Body)
	if len(og.Images) < 1 {
		return ""
	}

	return og.Images[0].URL
}
