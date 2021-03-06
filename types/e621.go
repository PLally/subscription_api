package types

import (
	"errors"
	"fmt"
	"github.com/plally/e621"
	"github.com/plally/subscription_api/subscription"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

func RegisterE621() *E621Handler {
	handler := &E621Handler{
		Session: e621.NewSession("e621.net", "FoxBotSubscriptions/0.1"),
	}

	viper.SetDefault("e621.cache_update_rate", time.Second*30)
	viper.SetDefault("e621.cache_size", 1000)

	subscription.SetSubTypeHandler("e621", handler)
	return handler
}

type E621Handler struct {
	Session     *e621.Session
	LastUpdated int
	postCache   []*e621.Post
}

func (e6 *E621Handler) StartPostCacheUpdater() {
	go func() {
		for {
			e6.updatePostCache()
			time.Sleep(viper.GetDuration("e621.cache_update_rate"))
		}
	}()

}
func (r *E621Handler) updatePostCache() {

	posts, err := r.getRecentPosts(viper.GetInt("e621.cache_size"))
	if err != nil {
		log.Error(err)
	}
	log.Infof("E621 Post cache now contains %v posts", len(posts.Posts))
	r.postCache = posts.Posts
}

func (r *E621Handler) getRecentPosts(amount int) (posts e621.PostsResponse, err error) {
	var id int
	s := r.Session
	for amount > 0 {
		var resp e621.PostsResponse
		limit := e621.MAX_LIMIT
		if amount < limit {
			limit = amount
		}
		tags := ""
		if id != 0 {
			tags = "id:<" + strconv.Itoa(id)
		}
		resp, err = s.GetPosts(tags, limit)
		if err != nil {
			return
		}
		id = resp.Posts[len(resp.Posts)-1].ID
		posts.Posts = append(posts.Posts, resp.Posts...)
		amount -= limit
		time.Sleep(time.Millisecond * 500)
	}
	return
}

func (r *E621Handler) GetNewItems(tags string) []subscription.SubscriptionItem {
	var items []subscription.SubscriptionItem
	parsed, _ := e621.ParseTags(tags, false)
	for _, post := range r.postCache {
		if !parsed.Matches(post.Tags) {
			continue
		}

		sub_item := subscription.SubscriptionItem{
			Title:       fmt.Sprintf("E621 Post #%v", post.ID),
			Url:         r.Session.PostUrl(post),
			Description: post.Description[:minInt(len(post.Description), 500)],
			Author:      strings.Join(post.Tags.Artist, ", "),
			TimeID:      int64(post.ID),
			Image:       post.File.URL,
		}
		items = append(items, sub_item)
	}
	return items
}

func (r *E621Handler) Validate(tags string) (string, error) {
	parsed, err := e621.ParseTags(tags, false)
	if err != nil {
		return "", err
	}

	tags = parsed.Normalized()
	tagsSplit := strings.Split(tags, " ")
	tagAmount := len(tagsSplit)
	for i := 0; i < tagAmount; i++ {
		tag := tagsSplit[i]
		if tag == "" {
			continue
		}

		prefix := ""
		if tag[0] == '-' {
			prefix = "-"
			tag = tag[1:]
		}
		if tag[0] == '~' {
			prefix = "~"
			tag = tag[1:]
		}

		aliases, _ := r.Session.FindAliases(tag)

		if len(aliases) != 0 {
			tag = aliases[0].ConsequentName
			tagsSplit[i] = prefix + tag
			continue
		}

		tags, _ := r.Session.FindTag(tag)
		if len(tags) != 0 {
			tagsSplit[i] = prefix + tag
			continue
		}
		return "", errors.New(fmt.Sprintf("Invalid tag %v", tag))

	}

	return strings.Join(tagsSplit, " "), nil
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
