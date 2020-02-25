package sub_types

import (
	"github.com/mmcdole/gofeed"
	"github.com/plally/subscription_api/subscription"
	log "github.com/sirupsen/logrus"
)

func init() {
	subscription.SetSubTypeHandler("rss", &RSSHandler{})
}

var feedWhitelist = map[string]bool{
	"http://www.housepetscomic.com/feed": true,
}

type RSSHandler struct{}

func (r *RSSHandler) GetType() string { return "rss" }

func (r *RSSHandler) GetNewItems(tags string) []subscription.SubscriptionItem {
	var items []subscription.SubscriptionItem
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(tags)
	if err != nil {
		log.Error(err)
		log.Debug(tags)
	}


	for _, item := range feed.Items {
		sub_item := subscription.SubscriptionItem{
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			Author:      item.Author.Name,
			TimeID:      item.PublishedParsed.Unix(),
		}
		items = append(items, sub_item)
	}
	return items
}

func (r *RSSHandler) Validate(tags string) (string, error) {
	return tags, nil
}