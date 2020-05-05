package types

import (
	"errors"
	"fmt"
	"github.com/plally/e621"
	"github.com/plally/subscription_api/subscription"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

func RegisterE621() {
	handler := &E621Handler{
		Session: e621.NewSession("e621.net", "FoxBotSubscriptions/0.1"),
	}
	go func() {
		for {
			handler.updatePostCache()
			time.Sleep(time.Minute * 15)
		}
	}()
	subscription.SetSubTypeHandler("e621", handler)
}

type E621Handler struct {
	Session     *e621.Session
	LastUpdated int
	postCache   []*e621.Post
}

func (r *E621Handler) updatePostCache() {
	posts, err := r.getRecentPosts(5000)
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

func (r *E621Handler) GetType() string { return "e621" }

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
			Description: fmt.Sprintf("Artists %v", strings.Join(post.Tags.Artist, ". ")),
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
		aliases, _ := r.Session.FindAliases(tag)

		if len(aliases) == 0 {
			return "", errors.New(fmt.Sprintf("Invalid tag %v", tag))
		}

		tag = aliases[0].ConsequentName
		tagsSplit[i] = tag
	}

	return strings.Join(tagsSplit, " "), nil
}
